// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package anvil

import (
	"reflect"
	"testing"
)

func TestLevelRoundtrip(t *testing.T) {
	const File1 = "../testdata/newworld/level.dat"
	const File2 = "../testdata/newworld/level.dat.2"

	la, err := LoadLevel(File1)
	if err != nil {
		t.Fatalf("Load 1: %v", err)
	}

	err = la.Save(File2)
	if err != nil {
		t.Fatalf("Save 2: %v", err)
	}

	lb, err := LoadLevel(File2)
	if err != nil {
		t.Fatalf("Load 2: %v", err)
	}

	if !reflect.DeepEqual(la, lb) {
		t.Fatalf("roundtrip mismatch:\nHave: %+v\nWant: %+v", lb, la)
	}
}
