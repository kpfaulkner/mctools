// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package item

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	idMask = 0xffff
	idBits = 16
)

// Id represents a single Minecraft resource id.
// It holds the primary id value in the lower 16 bits, along with an
// optional sub-id in the upper 16 bits.
//
// The sub id represents a damage value for a given item. While this is
// normally intended to indicate remaining durability, some blocks use it
// to denote the same item in different colours or styles. Wool is an
// example of this.
//
// Refer to item.go for a complete listing of known item ids.
type Id uint32

func (i Id) String() string {
	if str, ok := _Id_map[i]; ok {
		return str
	}

	// Some ids may contain extra data in the higher bits which are not
	// directly related to the item type. e.g.: an item with only a primary
	// id and no sub-id. Those will not be caught by the looking above us.
	// For these, we mask off the upper 16 bit and go only for the primary
	// id.
	if str, ok := _Id_map[i&0xffff]; ok {
		return str
	}

	// If even that fails, just yield the number as-is.
	return fmt.Sprintf("Id(%d)", i)
}

// ParseId returns a block id from the given string value.
// It is expected to be in the form: x:y
// Returns false if the string is not valid.
func ParseId(v string) (Id, bool) {
	elem := strings.Split(v, ":")
	if len(elem) != 2 {
		return 0, false
	}

	a, err := strconv.ParseInt(elem[0], 10, 32)
	if err != nil {
		return 0, false
	}

	b, err := strconv.ParseInt(elem[1], 10, 32)
	if err != nil {
		return 0, false
	}

	return NewId(int(a), int(b)), true
}

// NewId creates a new Id from the given values.
func NewId(primary, sub int) Id {
	return (Id(primary) & idMask) | ((Id(sub) & idMask) << idBits)
}

// Compare compares the given id with the current one.
//
//     Returns 0 if they are equal.
//     Returns <0 if the current id is smaller.
//     Returns >0 if the current id is greater.
//
func (a Id) Compare(b Id) int {
	if a == b {
		return 0
	}

	va, vb := a.Primary(), b.Primary()
	if va < vb {
		return -1
	}

	if va > vb {
		return +1
	}

	va, vb = a.Sub(), b.Sub()
	if va < vb {
		return -1
	}

	return +1
}

// Primary returns the primary id component.
func (id Id) Primary() int { return int(id & idMask) }

// Sub returns the sub id component.
func (id Id) Sub() int { return int(id>>idBits) & idMask }
