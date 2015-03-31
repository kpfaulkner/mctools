// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package nbt

import (
	"fmt"
	"reflect"
)

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

// UnmarshalError describes an invalid argument passed to Unmarshal/Decode.
type UnmarshalError struct {
	Type reflect.Type
}

func (e *UnmarshalError) Error() string {
	if e.Type == nil {
		return "nbt: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return fmt.Sprintf("nbt: Unmarshal non-pointer %s", e.Type)
	}

	return fmt.Sprintf("nbt: Unmarshal %s", e.Type)
}

// MarshalError describes an invalid argument passed to Marshal/Encode.
type MarshalError struct {
	Name string
	Type reflect.Type
}

func (e *MarshalError) Error() string {
	return fmt.Sprintf("nbt: unsupported type %s(%q)", e.Type, e.Name)
}
