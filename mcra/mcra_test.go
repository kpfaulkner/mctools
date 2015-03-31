// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mcra

import (
	"fmt"
	"os"
	"testing"

	"github.com/jteeuwen/mctools"
	"github.com/jteeuwen/mctools/anvil"
	"github.com/jteeuwen/mctools/anvil/item"
)

var (
	world  *mctools.World
	region *anvil.Region
)

const worldPath = "../testdata/newworld/"

func TestMain(t *testing.M) {
	var err error

	world, err = mctools.Open(worldPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	region, err = world.LoadRegion(mctools.DimensionOverworld, 0, -1)
	if err != nil {
		fmt.Println(err)
		return
	}

	code := t.Run()

	os.Exit(code)
}

func TestDungeons(t *testing.T) {
	set := FindDungeons(region)

	if len(set) != 5 {
		t.Fatalf("expected 5 results; have %d", len(set))
	}
}

func TestStrongholds(t *testing.T) {
	set := FindStrongholds(region)

	if len(set) != 0 {
		t.Fatalf("expected 0 results; have %d", len(set))
	}
}

func TestInclusionQuery(t *testing.T) {
	result := FindInRegion(region, NewInclusionQuery(
		item.Sand,
		item.Sandstone,
	))

	if result.Len() != 159713 {
		t.Fatalf("expected 159713 results; have %d", result.Len())
	}
}

func TestExclusionQuery(t *testing.T) {
	result := FindInRegion(region, NewExclusionQuery(
		item.Air,
		item.Bedrock,
	))

	if result.Len() != 6355205 {
		t.Fatalf("expected 6355205 results; have %d", result.Len())
	}
}

func TestRadiusQuery(t *testing.T) {
	result := FindInRegion(region, NewRadiusQuery(
		Location{
			RX: int8(region.X),
			RZ: int8(region.Z),
		},
		100, // radius
		item.Bedrock,
		item.Air,
	))

	if result.Len() != 964347 {
		t.Fatalf("expected 964347 results; have %d", result.Len())
	}
}

func TestTally(t *testing.T) {
	tally := TallyInRegion(
		region,
		item.RedstoneOre,
		item.DiamondOre,
	)

	var want uint64

loopy:
	for k, v := range tally {
		switch k {
		case item.RedstoneOre:
			want = 11568
			if v == want {
				continue loopy
			}

		case item.DiamondOre:
			want = 1440
			if v == want {
				continue loopy
			}
		}

		t.Fatalf("%s: expected %d results; have %d", k, want, v)
	}
}
