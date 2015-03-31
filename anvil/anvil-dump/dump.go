// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
)

func dumpFile(w io.Writer, file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}

	defer fd.Close()

	gz, err := gzip.NewReader(fd)
	if err != nil {
		return err
	}

	defer gz.Close()
	return dump(w, gz)
}

func dump(w io.Writer, r io.Reader) (err error) {
	d := &dumper{w: w, r: r}

	defer func() {
		x := recover()
		if x == nil {
			return
		}

		switch v := x.(type) {
		case runtime.Error:
			panic(x) // Re-throw runtime errors.
		case error:
			err = v
		default:
			err = fmt.Errorf("%v", v)
		}
	}()

	d.dump("", tagUnknown)
	return
}

// assert throws a panic if the given error is not nil.
//
// It prefixes the error with the given formatted message.
func assert(err error, msg string, argv ...interface{}) {
	if err != nil {
		msg = fmt.Sprintf(msg, argv...)
		panic(fmt.Errorf("nbt: %s: %v", msg, err))
	}
}

// errorf aborts the parsing by panicing with the given, formatted error.
func errorf(msg string, argv ...interface{}) {
	panic(fmt.Sprintf("nbt: %s", fmt.Sprintf(msg, argv...)))
}

type dumper struct {
	w       io.Writer // Output stream.
	r       io.Reader // Input stream.
	scratch [8]byte   // Temporary read buffer.
}

func (d *dumper) dump(indent string, id tagId) bool {
	var name string

	if id == tagUnknown {
		id = tagId(d.readByte())
		if id != tagEnd {
			name = d.readString()
		}
	}

	switch id {
	case tagEnd:
		return false

	case tagList:
		d.dumpList(indent, name)

	case tagCompound:
		d.dumpCompound(indent, name)

	case tagByteArray:
		d.dumpByteArray(indent, name)

	case tagIntArray:
		d.dumpIntArray(indent, name)

	case tagString:
		d.dumpString(indent, name)

	default:
		d.dumpScalar(indent, id, name)
	}

	return true
}

func (d *dumper) dumpList(indent, name string) {
	id := tagId(d.readByte())
	size := int(d.readInt())

	fmt.Fprintf(d.w, "%s%s(%q) [%d] {\n", indent, tagList, name, size)

	for i := 0; i < size; i++ {
		d.dump(indent+"  ", id)
	}

	fmt.Fprintf(d.w, "%s}\n", indent)
}

func (d *dumper) dumpCompound(indent, name string) {
	fmt.Fprintf(d.w, "%s%s(%q) {\n", indent, tagCompound, name)

	for d.dump(indent+"  ", tagUnknown) {
	}

	fmt.Fprintf(d.w, "%s}\n", indent)
}

func (d *dumper) dumpByteArray(indent, name string) {
	value := d.readByteArray()

	fmt.Fprintf(d.w, "%s%s(%q) [%d] {\n", indent, tagByteArray, name, len(value))

	if len(value) < 20 {
		fmt.Fprintf(d.w, "%s%v\n", indent+"  ", value)
	} else {
		fmt.Fprintf(d.w, "%s%v...\n", indent+"  ", value[:20])
	}

	fmt.Fprintf(d.w, "%s}\n", indent)
}

func (d *dumper) dumpIntArray(indent, name string) {
	value := d.readIntArray()

	fmt.Fprintf(d.w, "%s%s(%q) [%d] {\n", indent, tagIntArray, name, len(value))

	if len(value) < 10 {
		fmt.Fprintf(d.w, "%s%v\n", indent+"  ", value)
	} else {
		fmt.Fprintf(d.w, "%s%v...\n", indent+"  ", value[:10])
	}

	fmt.Fprintf(d.w, "%s}\n", indent)
}

func (d *dumper) dumpString(indent, name string) {
	value := d.readString()

	if len(value) < 30 {
		fmt.Fprintf(d.w, "%s%s(%q): %q\n", indent, tagString, name, value)
	} else {
		fmt.Fprintf(d.w, "%s%s(%q): %.30q...\n", indent, tagString, name, value)
	}
}

func (d *dumper) dumpScalar(indent string, id tagId, name string) {
	var value interface{}

	switch id {
	case tagByte:
		value = d.readByte()
	case tagShort:
		value = d.readShort()
	case tagInt:
		value = d.readInt()
	case tagLong:
		value = d.readLong()
	case tagFloat:
		value = d.readFloat()
	case tagDouble:
		value = d.readDouble()
	}

	fmt.Fprintf(d.w, "%s%s(%q): %v\n", indent, id, name, value)
}

func (d *dumper) readByte() int8 {
	_, err := io.ReadFull(d.r, d.scratch[:1])
	assert(err, "read byte")
	return int8(d.scratch[0])
}

func (d *dumper) readShort() int16 {
	_, err := io.ReadFull(d.r, d.scratch[:2])
	assert(err, "read short")
	return int16(d.scratch[0])<<8 | int16(d.scratch[1])
}

func (d *dumper) readInt() int32 {
	_, err := io.ReadFull(d.r, d.scratch[:4])
	assert(err, "read int")
	return int32(d.scratch[0])<<24 | int32(d.scratch[1])<<16 |
		int32(d.scratch[2])<<8 | int32(d.scratch[3])
}

func (d *dumper) readLong() int64 {
	_, err := io.ReadFull(d.r, d.scratch[:8])
	assert(err, "read long")
	return int64(d.scratch[0])<<56 | int64(d.scratch[1])<<48 |
		int64(d.scratch[2])<<40 | int64(d.scratch[3])<<32 |
		int64(d.scratch[4])<<24 | int64(d.scratch[5])<<16 |
		int64(d.scratch[6])<<8 | int64(d.scratch[7])
}

func (d *dumper) readFloat() float32 {
	_, err := io.ReadFull(d.r, d.scratch[:4])
	assert(err, "read float")
	v := uint32(d.scratch[0])<<24 | uint32(d.scratch[1])<<16 |
		uint32(d.scratch[2])<<8 | uint32(d.scratch[3])
	return math.Float32frombits(v)
}

func (d *dumper) readDouble() float64 {
	_, err := io.ReadFull(d.r, d.scratch[:8])
	assert(err, "read double")
	v := uint64(d.scratch[0])<<56 | uint64(d.scratch[1])<<48 |
		uint64(d.scratch[2])<<40 | uint64(d.scratch[3])<<32 |
		uint64(d.scratch[4])<<24 | uint64(d.scratch[5])<<16 |
		uint64(d.scratch[6])<<8 | uint64(d.scratch[7])
	return math.Float64frombits(v)
}

func (d *dumper) readByteArray() []byte {
	size := d.readInt()
	if size < 0 {
		errorf("TAG_Byte_Array with size < 0")
	}

	if size == 0 {
		return nil
	}

	v := make([]byte, size)
	_, err := io.ReadFull(d.r, v)
	assert(err, "read byte array of size %d", size)
	return v
}

func (d *dumper) readString() string {
	size := d.readShort()
	if size < 0 {
		errorf("TAG_String with size < 0")
	}

	if size == 0 {
		return ""
	}

	v := make([]byte, size)
	_, err := io.ReadFull(d.r, v)
	assert(err, "read string of size %d", size)
	return string(v)
}

func (d *dumper) readIntArray() []int32 {
	size := d.readInt()
	if size < 0 {
		errorf("TAG_Int_Array with size < 0")
	}

	if size == 0 {
		return nil
	}

	v := make([]int32, size)
	for i := 0; i < int(size); i++ {
		v[i] = d.readInt()
	}

	return v
}
