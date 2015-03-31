// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package nbt

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// Unmarshal parses the uncompressed, NBT-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(r io.Reader, v interface{}) error {
	dec := NewDecoder(r)
	return dec.Decode(v)
}

// Decoder defines a NBT decoder, used to unmarshal uncompressed,
// NBT formatted data into a Go type.
type Decoder struct {
	r       io.Reader // Input stream.
	scratch [8]byte   // Temporary read buffer.
}

// NewDecoder creates a new decoder for the given input stream.
func NewDecoder(r io.Reader) *Decoder { return &Decoder{r: r} }

// Decode recursively reads tags and unmarshals them into  the given value.
func (d *Decoder) Decode(v interface{}) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &UnmarshalError{reflect.TypeOf(v)}
	}

	id, name, err := d.readHeader(tagUnknown)
	if err != nil {
		return fmt.Errorf("nbt: %v", err)
	}

	err = d.decode(id, name, rv)
	if err != nil {
		return fmt.Errorf("nbt: %v", err)
	}

	return nil
}

func (d *Decoder) decode(id tagId, name string, rv reflect.Value) error {
	// Initialize a new instance of a pointer type if needed.
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}

		// Pointer types should be dereferenced before we go on.
		rv = rv.Elem()
	}

	//fmt.Printf("%s(%q) => %v\n", id, name, rv)

	var err error
	switch id {
	case tagList:
		err = d.decodeList(name, rv)

	case tagCompound:
		err = d.decodeCompound(name, rv)

	default:
		err = d.decodeValue(id, name, rv)
	}

	return err
}

func (d *Decoder) decodeCompound(name string, rv reflect.Value) error {
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("%s(%q): value %v must be a struct", tagCompound, name, rv)
	}

	// Decode until we have a matching TagEnd.
	for {
		id, name, err := d.readHeader(tagUnknown)
		if err != nil {
			return err
		}

		if id == tagEnd {
			break
		}

		fv := readField(rv, name)

		if fv.Kind() == reflect.Invalid {
			err = d.skip(id)
		} else {
			err = d.decode(id, name, fv)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) decodeList(name string, rv reflect.Value) error {
	if rv.Kind() != reflect.Slice {
		return fmt.Errorf("%s(%q): value %v must be slice", tagCompound, name, rv)
	}

	n, err := d.readByte()
	if err != nil {
		return err
	}

	size, err := d.readInt()
	if err != nil {
		return err
	}

	if size == 0 {
		return nil
	}

	id := tagId(n)
	rt := rv.Type()
	et := rt.Elem()

	new := reflect.MakeSlice(rt, 0, int(size))

	for i := 0; i < int(size); i++ {
		var elem reflect.Value

		if et.Kind() == reflect.Ptr {
			elem = reflect.New(rt.Elem().Elem())
			err = d.decode(id, "", elem)
		} else {
			elem = reflect.New(rt.Elem())
			elem = reflect.Indirect(elem)
			err = d.decode(id, "", reflect.Indirect(elem))
		}

		if err != nil {
			return err
		}

		new = reflect.Append(new, elem)
	}

	rv.Set(new)
	return nil
}

func (d *Decoder) decodeValue(id tagId, name string, rv reflect.Value) error {
	var value interface{}
	var err error

	switch id {
	case tagByte:
		value, err = d.readByte()
	case tagShort:
		value, err = d.readShort()
	case tagInt:
		value, err = d.readInt()
	case tagLong:
		value, err = d.readLong()
	case tagFloat:
		value, err = d.readFloat()
	case tagDouble:
		value, err = d.readDouble()
	case tagString:
		value, err = d.readString()
	case tagByteArray:
		value, err = d.readByteArray()
	case tagIntArray:
		value, err = d.readIntArray()
	default:
		err = fmt.Errorf("unsupported value %s for field %q", id, name)
	}

	if err != nil {
		return err
	}

	return d.set(id, name, rv, value)
}

// set assigns src to dst if possible.
// It performs implicit type conversions where applicable.
func (d *Decoder) set(id tagId, name string, dst reflect.Value, src interface{}) error {
	if !dst.IsValid() {
		return fmt.Errorf("%s(%q): invalid value for %T", id, name, src)
	}

	if !dst.CanSet() {
		return fmt.Errorf("%s(%q): can not assign %T to %v", id, name, src, dst.Type())
	}

	var err error
	dt := dst.Type()
	sv := reflect.ValueOf(src)
	st := sv.Type()

	if !st.AssignableTo(dt) {
		sv, err = convert(sv, dt, st)
		if err != nil {
			return fmt.Errorf("%s(%q): %v", id, name, err)
		}
	}

	dst.Set(sv)
	return nil
}

func (d *Decoder) skip(id tagId) error {
	var err error

	switch id {
	case tagList:
		err = d.skipList()
	case tagCompound:
		err = d.skipCompound()
	default:
		err = d.skipValue(id)
	}

	return err
}

func (d *Decoder) skipCompound() error {
	for {
		id, _, err := d.readHeader(tagUnknown)
		if err != nil {
			return err
		}

		if id == tagEnd {
			break
		}

		err = d.skip(id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) skipList() error {
	n, err := d.readByte()
	if err != nil {
		return err
	}

	size, err := d.readInt()
	if err != nil {
		return err
	}

	if size == 0 {
		return nil
	}

	for i := 0; i < int(size); i++ {
		err = d.skip(tagId(n))
		if err != nil {
			return err
		}
	}

	return nil
}
func (d *Decoder) skipValue(id tagId) error {
	var err error

	switch id {
	case tagByte:
		_, err = d.readByte()
	case tagShort:
		_, err = d.readShort()
	case tagInt:
		_, err = d.readInt()
	case tagLong:
		_, err = d.readLong()
	case tagFloat:
		_, err = d.readFloat()
	case tagDouble:
		_, err = d.readDouble()
	case tagString:
		_, err = d.readString()
	case tagByteArray:
		_, err = d.readByteArray()
	case tagIntArray:
		_, err = d.readIntArray()
	default:
		err = fmt.Errorf("unsupported value %s", id)
	}

	return err
}

// readHeader reads the next tag header.
func (d *Decoder) readHeader(id tagId) (tagId, string, error) {
	if id != tagUnknown {
		// We're in a list -- type/name already known.
		return id, "", nil
	}

	n, err := d.readByte()
	if err != nil {
		return tagEnd, "", err
	}

	id = tagId(n)

	var name string
	if id != tagEnd {
		name, err = d.readString()
		if err != nil {
			return tagEnd, "", err
		}
	}

	return id, name, nil
}

func (d *Decoder) readByte() (int8, error) {
	_, err := io.ReadFull(d.r, d.scratch[:1])
	return int8(d.scratch[0]), err
}

func (d *Decoder) readShort() (int16, error) {
	_, err := io.ReadFull(d.r, d.scratch[:2])
	return int16(d.scratch[0])<<8 | int16(d.scratch[1]), err
}

func (d *Decoder) readInt() (int32, error) {
	_, err := io.ReadFull(d.r, d.scratch[:4])
	return int32(d.scratch[0])<<24 | int32(d.scratch[1])<<16 |
		int32(d.scratch[2])<<8 | int32(d.scratch[3]), err
}

func (d *Decoder) readLong() (int64, error) {
	_, err := io.ReadFull(d.r, d.scratch[:8])
	return int64(d.scratch[0])<<56 | int64(d.scratch[1])<<48 |
		int64(d.scratch[2])<<40 | int64(d.scratch[3])<<32 |
		int64(d.scratch[4])<<24 | int64(d.scratch[5])<<16 |
		int64(d.scratch[6])<<8 | int64(d.scratch[7]), err
}

func (d *Decoder) readFloat() (float32, error) {
	_, err := io.ReadFull(d.r, d.scratch[:4])
	v := uint32(d.scratch[0])<<24 | uint32(d.scratch[1])<<16 |
		uint32(d.scratch[2])<<8 | uint32(d.scratch[3])
	return math.Float32frombits(v), err
}

func (d *Decoder) readDouble() (float64, error) {
	_, err := io.ReadFull(d.r, d.scratch[:8])
	v := uint64(d.scratch[0])<<56 | uint64(d.scratch[1])<<48 |
		uint64(d.scratch[2])<<40 | uint64(d.scratch[3])<<32 |
		uint64(d.scratch[4])<<24 | uint64(d.scratch[5])<<16 |
		uint64(d.scratch[6])<<8 | uint64(d.scratch[7])
	return math.Float64frombits(v), err
}

func (d *Decoder) readByteArray() ([]byte, error) {
	size, err := d.readInt()
	if err != nil {
		return nil, err
	}

	if size < 0 {
		return nil, fmt.Errorf("%s with size < 0", tagByteArray)
	}

	if size == 0 {
		return nil, nil
	}

	out := make([]byte, size)
	_, err = io.ReadFull(d.r, out)
	return out, err
}

func (d *Decoder) readString() (string, error) {
	size, err := d.readShort()
	if err != nil {
		return "", err
	}

	if size < 0 {
		return "", fmt.Errorf("%s with size < 0", tagString)
	}

	if size == 0 {
		return "", nil
	}

	out := make([]byte, size)
	_, err = io.ReadFull(d.r, out)
	return string(out), err
}

func (d *Decoder) readIntArray() ([]int32, error) {
	size, err := d.readInt()
	if err != nil {
		return nil, err
	}

	if size < 0 {
		return nil, fmt.Errorf("%s with size < 0", tagIntArray)
	}

	if size == 0 {
		return nil, nil
	}

	out := make([]int32, size)

	for i := 0; i < int(size); i++ {
		out[i], err = d.readInt()
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

// readField finds a field in the given struct with the specified name
// and returns its value.
//
// If no match can be found in the given struct, this recursively finds and
// searches anonymous, embedded structs and checks them for matching fields.
// If multiple anonymous structs export a field with the same name, the
// first one we encounter will be used.
//
// If no match can be found, reflect.Invalid is returned.
func readField(rv reflect.Value, name string) reflect.Value {
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return rv
	}

	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		ft := rt.Field(i)

		if hasFieldName(ft, name) {
			return rv.Field(i)
		}

		if !ft.Anonymous {
			continue
		}

		ret := readField(rv.Field(i), name)
		if ret.Kind() != reflect.Invalid {
			return ret
		}
	}

	return reflect.Value{}
}

// hasFieldName returns true if the given struct field has the specified name.
// This first checks for the presence of a matching "nbt" field tag. Otherwise
// the field name itself is considered.
//
// This is a case-sensitive check.
func hasFieldName(ft reflect.StructField, name string) bool {
	tag := ft.Tag.Get("nbt")
	elem := strings.Split(tag, ",")

	for _, v := range elem {
		if len(v) > 0 && v == name {
			return true
		}
	}

	return ft.Name == name
}

// convert tries to convert rv from the source type to destination type.
// Returns an error if this fails. This applies the numerous type conversions
// we need to support when going between NBT and Go data.
//
// Refer to the "Type compatibility" section in the `nbt` package README.
func convert(rv reflect.Value, dst, src reflect.Type) (reflect.Value, error) {
	if src.ConvertibleTo(dst) {
		return rv.Convert(dst), nil
	}

	switch src.Kind() {
	case reflect.Slice:
		if dst.Kind() != reflect.Slice {
			break
		}

		switch src.Elem().Kind() {
		case reflect.Int8:
			v := rv.Interface().([]int8)

			switch dst.Elem().Kind() {
			case reflect.Uint8:
				if len(v) == 0 {
					return reflect.ValueOf(([]uint8)(nil)), nil
				}

				ptr := (*(*[1<<31 - 1]uint8)(unsafe.Pointer(&v[0])))[:len(v)]
				return reflect.ValueOf(ptr), nil
			}

		case reflect.Uint8:
			v := rv.Interface().([]uint8)

			switch dst.Elem().Kind() {
			case reflect.Int8:
				if len(v) == 0 {
					return reflect.ValueOf(([]int8)(nil)), nil
				}

				ptr := (*(*[1<<31 - 1]int8)(unsafe.Pointer(&v[0])))[:len(v)]
				return reflect.ValueOf(ptr), nil
			}

		case reflect.Int32:
			v := rv.Interface().([]int32)

			switch dst.Elem().Kind() {
			case reflect.Uint32:
				if len(v) == 0 {
					return reflect.ValueOf(([]uint32)(nil)), nil
				}

				ptr := (*(*[1<<31 - 1]uint32)(unsafe.Pointer(&v[0])))[:len(v)]
				return reflect.ValueOf(ptr), nil
			}

		case reflect.Uint32:
			v := rv.Interface().([]uint32)

			switch dst.Elem().Kind() {
			case reflect.Int32:
				if len(v) == 0 {
					return reflect.ValueOf(([]int32)(nil)), nil
				}

				ptr := (*(*[1<<31 - 1]int32)(unsafe.Pointer(&v[0])))[:len(v)]
				return reflect.ValueOf(ptr), nil
			}
		}

	case reflect.String:
		v := rv.String()

		switch dst.Kind() {
		case reflect.Bool:
			b, err := strconv.ParseBool(v)
			if err == nil {
				return reflect.ValueOf(b), nil
			}
		}

	case reflect.Int8:
		v := rv.Interface().(int8)

		switch dst.Kind() {
		case reflect.Bool:
			return reflect.ValueOf(v != 0), nil
		case reflect.Uint8:
			return reflect.ValueOf(uint8(v)), nil
		case reflect.Int16:
			return reflect.ValueOf(int16(v)), nil
		case reflect.Uint16:
			return reflect.ValueOf(uint16(v)), nil
		case reflect.Int32:
			return reflect.ValueOf(int32(v)), nil
		case reflect.Uint32:
			return reflect.ValueOf(uint32(v)), nil
		case reflect.Int64:
			return reflect.ValueOf(int64(v)), nil
		case reflect.Uint64:
			return reflect.ValueOf(uint64(v)), nil
		}

	case reflect.Int16:
		v := rv.Interface().(int16)

		switch dst.Kind() {
		case reflect.Uint16:
			return reflect.ValueOf(uint16(v)), nil
		case reflect.Int32:
			return reflect.ValueOf(int32(v)), nil
		case reflect.Uint32:
			return reflect.ValueOf(uint16(v)), nil
		case reflect.Int64:
			return reflect.ValueOf(int64(v)), nil
		case reflect.Uint64:
			return reflect.ValueOf(uint64(v)), nil
		}

	case reflect.Int32:
		v := rv.Interface().(int32)

		switch dst.Kind() {
		case reflect.Uint32:
			return reflect.ValueOf(uint32(v)), nil
		case reflect.Int64:
			return reflect.ValueOf(int64(v)), nil
		case reflect.Uint64:
			return reflect.ValueOf(uint64(v)), nil
		}

	case reflect.Int64:
		v := rv.Interface().(int64)

		switch dst.Kind() {
		case reflect.Struct:
			if dst.Name() == "Time" {
				return reflect.ValueOf(time.Unix(v, 0)), nil
			}
		case reflect.Uint64:
			return reflect.ValueOf(uint64(v)), nil
		}

	case reflect.Float32:
		v := rv.Interface().(float32)

		switch dst.Kind() {
		case reflect.Float64:
			return reflect.ValueOf(float64(v)), nil
		}
	}

	return rv, fmt.Errorf("can not convert %v(%v) to %v", src, rv.Interface(), dst)
}
