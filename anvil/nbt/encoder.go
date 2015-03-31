// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package nbt

import (
	"io"
	"math"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

// Marshal translates data into uncompressed, NBT-encoded data and writes
// it to the given stream.
func Marshal(w io.Writer, v interface{}) error {
	enc := NewEncoder(w)
	return enc.Encode(v)
}

// Encoder translates a Go type into a stream of NBT encoded data.
type Encoder struct {
	w io.Writer
}

// NewEncoder creates a new encoder for the given value.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode translates v into uncompressed, NBT-encoded data and writes
// it to the underlying stream.
func (e *Encoder) Encode(v interface{}) error {
	rv := reflect.ValueOf(v)

	if !rv.IsValid() {
		return &MarshalError{Type: reflect.TypeOf(v)}
	}

	return e.encode(rv, "", false)
}

// Encode translates v into uncompressed, NBT-encoded data and writes
// it to the underlying stream.
func (e *Encoder) encode(rv reflect.Value, name string, inlist bool) error {
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return nil
	}

	rv = reflect.Indirect(rv)

	switch rv.Kind() {
	case reflect.Struct:
		return e.encodeStruct(rv, name, inlist)

	case reflect.Array, reflect.Slice:
		return e.encodeSlice(rv, name, inlist)

	case reflect.String:
		return e.encodeString(rv, name, inlist)

	case reflect.Bool:
		return e.encodeByte(rv, name, inlist)

	case reflect.Float32:
		return e.encodeFloat(rv, name, inlist)

	case reflect.Float64:
		return e.encodeDouble(rv, name, inlist)

	case reflect.Int8, reflect.Uint8:
		return e.encodeByte(rv, name, inlist)

	case reflect.Int16, reflect.Uint16:
		return e.encodeShort(rv, name, inlist)

	case reflect.Int32, reflect.Uint32:
		return e.encodeInt(rv, name, inlist)

	case reflect.Int64, reflect.Uint64:
		return e.encodeLong(rv, name, inlist)
	}

	return &MarshalError{Name: name, Type: rv.Type()}
}

func (e *Encoder) encodeStruct(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagCompound, name, inlist)
	if err != nil {
		return err
	}

	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)

		if hasField(ft.Tag.Get("nbt"), "omitempty") && isEmpty(fv) {
			continue
		}

		fname := tagField(ft.Tag.Get("nbt"), 0)
		if len(fname) == 0 {
			fname = ft.Name
		}

		// Special-case time.Time
		if e.isTime(ft.Type) {
			t := fv.Interface().(time.Time).Unix()
			fv = reflect.ValueOf(t)
		}

		err = e.encode(fv, fname, false)
		if err != nil {
			return err
		}
	}

	return e.writeU8(uint8(tagEnd))
}

// isTime returns true if rv is a valid type for time.Time.
func (e *Encoder) isTime(rt reflect.Type) bool {
	var t time.Time
	tt := reflect.TypeOf(t)
	return reflect.DeepEqual(tt, rt)
}

func (e *Encoder) encodeSlice(rv reflect.Value, name string, inlist bool) error {
	rt := rv.Type()
	et := rt.Elem()

	switch et.Kind() {
	case reflect.Int8, reflect.Uint8:
		return e.encodeByteArray(rv, name, inlist)

	case reflect.Int32, reflect.Uint32:
		return e.encodeIntArray(rv, name, inlist)

	default:
		return e.encodeList(rv, name, inlist)
	}

	return nil
}

func (e *Encoder) encodeByteArray(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagByteArray, name, inlist)
	if err != nil {
		return err
	}

	err = e.writeU32(uint32(rv.Len()))
	if err != nil {
		return err
	}

	if rv.Len() == 0 {
		return nil
	}

	data := rv.Interface()
	out, ok := data.([]uint8)
	if !ok {
		i8 := data.([]int8)
		out = (*(*[1<<31 - 1]uint8)(unsafe.Pointer(&i8[0])))[:len(i8)]
	}

	_, err = e.w.Write(out)
	return err
}

func (e *Encoder) encodeIntArray(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagIntArray, name, inlist)
	if err != nil {
		return err
	}

	size := uint32(rv.Len())
	err = e.writeU32(size)
	if err != nil {
		return err
	}

	if size == 0 {
		return nil
	}

	data := rv.Interface()
	out := make([]byte, 0, size*4)

	u32, ok := data.([]uint32)
	if !ok {
		i32 := data.([]int32)
		u32 = (*(*[1<<31 - 1]uint32)(unsafe.Pointer(&i32[0])))[:rv.Len()]
	}

	for _, v := range u32 {
		out = append(out, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
	}

	_, err = e.w.Write(out)
	return err
}

func (e *Encoder) encodeList(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagList, name, inlist)
	if err != nil {
		return err
	}

	rt := rv.Type()
	et := rt.Elem()

	var id tagId

	switch et.Kind() {
	case reflect.Ptr:
		id = tagCompound

	case reflect.Struct:
		if e.isTime(et) { // Special-case time.Time
			id = tagLong
		} else {
			id = tagCompound
		}

	case reflect.Uint16, reflect.Int16:
		id = tagShort

	case reflect.Uint64, reflect.Int64:
		id = tagLong

	case reflect.Float32:
		id = tagFloat

	case reflect.Float64:
		id = tagDouble

	default:
		return &MarshalError{Name: name, Type: rt}
	}

	err = e.writeU8(uint8(id))
	if err != nil {
		return err
	}

	err = e.writeU32(uint32(rv.Len()))
	if err != nil {
		return err
	}

	if rv.Len() == 0 {
		return nil
	}

	for i := 0; i < rv.Len(); i++ {
		iv := rv.Index(i)

		if e.isTime(et) {
			t := iv.Interface().(time.Time).Unix()
			iv = reflect.ValueOf(t)
		}

		err = e.encode(iv, "", true)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeString(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagString, name, inlist)
	if err != nil {
		return err
	}

	return e.writeString(rv.String())
}

func (e *Encoder) encodeByte(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagByte, name, inlist)
	if err != nil {
		return err
	}

	switch rv.Kind() {
	case reflect.Int8:
		return e.writeI8(int8(rv.Int()))

	case reflect.Uint8:
		return e.writeU8(uint8(rv.Uint()))

	case reflect.Bool:
		if rv.Bool() {
			return e.writeU8(1)
		}
		return e.writeU8(0)
	}

	return nil
}

func (e *Encoder) encodeShort(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagShort, name, inlist)
	if err != nil {
		return err
	}

	switch rv.Kind() {
	case reflect.Int16:
		return e.writeI16(int16(rv.Int()))

	case reflect.Uint16:
		return e.writeU16(uint16(rv.Uint()))
	}

	return nil
}

func (e *Encoder) encodeInt(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagInt, name, inlist)
	if err != nil {
		return err
	}

	switch rv.Kind() {
	case reflect.Int32:
		return e.writeI32(int32(rv.Int()))

	case reflect.Uint32:
		return e.writeU32(uint32(rv.Uint()))
	}

	return nil
}

func (e *Encoder) encodeLong(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagLong, name, inlist)
	if err != nil {
		return err
	}

	switch rv.Kind() {
	case reflect.Int64:
		return e.writeI64(rv.Int())

	case reflect.Uint64:
		return e.writeU64(rv.Uint())
	}

	return nil
}

func (e *Encoder) encodeFloat(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagFloat, name, inlist)
	if err != nil {
		return err
	}
	return e.writeF32(float32(rv.Float()))
}

func (e *Encoder) encodeDouble(rv reflect.Value, name string, inlist bool) error {
	err := e.emit(tagDouble, name, inlist)
	if err != nil {
		return err
	}
	return e.writeF64(rv.Float())
}

func (e *Encoder) emit(id tagId, name string, inlist bool) error {
	if inlist {
		return nil
	}

	err := e.writeU8(uint8(id))
	if err != nil {
		return err
	}

	return e.writeString(name)
}

func (e *Encoder) writeF32(v float32) error {
	bits := math.Float32bits(v)
	return e.writeU32(bits)
}

func (e *Encoder) writeF64(v float64) error {
	bits := math.Float64bits(v)
	return e.writeU64(bits)
}

func (e *Encoder) writeString(v string) error {
	err := e.writeI16(int16(len(v)))
	if err != nil {
		return err
	}

	_, err = e.w.Write([]byte(v))
	return err
}

func (e *Encoder) writeI8(v int8) error {
	return e.writeU8(byte(v))
}

func (e *Encoder) writeU8(v uint8) error {
	_, err := e.w.Write([]byte{v})
	return err
}

func (e *Encoder) writeI16(v int16) error {
	_, err := e.w.Write([]byte{byte(v >> 8), byte(v)})
	return err
}

func (e *Encoder) writeU16(v uint16) error {
	_, err := e.w.Write([]byte{byte(v >> 8), byte(v)})
	return err
}

func (e *Encoder) writeI32(v int32) error {
	_, err := e.w.Write([]byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	})
	return err
}

func (e *Encoder) writeU32(v uint32) error {
	_, err := e.w.Write([]byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	})
	return err
}

func (e *Encoder) writeI64(v int64) error {
	_, err := e.w.Write([]byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	})
	return err
}

func (e *Encoder) writeU64(v uint64) error {
	_, err := e.w.Write([]byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	})
	return err
}

// isEmpty returns true if the given value defines a zero value for whatever
// type it represents.
func isEmpty(rv reflect.Value) bool {
	if !rv.IsValid() {
		return true
	}

	switch rv.Kind() {
	case reflect.Ptr:
		return rv.IsNil()
	case reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len() == 0
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	}

	return false
}

// tagField returns the nth value in a comma-separated list of tag properties.
// E.g.: `nbt:"foo,omitempty"`. field 0 would be "foo", field 1 would be
// "omitempty".
func tagField(tag string, n int) string {
	elem := strings.Split(tag, ",")
	if len(elem) < n {
		return ""
	}
	return elem[n]
}

// hasField returns true if the given tag field exists.
func hasField(tag, value string) bool {
	elem := strings.Split(tag, ",")

	for _, v := range elem {
		if len(v) > 0 && strings.EqualFold(v, value) {
			return true
		}
	}

	return false
}
