// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mcra

import (
	"fmt"
	"math"

	"github.com/jteeuwen/mctools/anvil"
	"github.com/jteeuwen/mctools/anvil/item"
)

// Location defines a specific block location in a world.
type Location struct {
	RX int8  // Region X.
	RZ int8  // Region Z.
	CX int8  // Chunk X in region.
	CZ int8  // Chunk Z in region.
	BX uint8 // Block X in chunk.
	BY uint8 // Block Y in chunk.
	BZ uint8 // Block Z in chunk.
}

func (l Location) String() string {
	x := int(l.RX)*anvil.BlocksPerRegion + int(l.CX)*anvil.BlocksPerChunk + int(l.BX)
	z := int(l.RZ)*anvil.BlocksPerRegion + int(l.CZ)*anvil.BlocksPerChunk + int(l.BZ)
	y := int(l.BY)
	return fmt.Sprintf("%d, %d, %d", x, y, z)
}

// DistanceTo returns the distance, in blocks, between the two given locations.
func (l Location) DistanceTo(b Location) uint {
	ax := int(l.RX)*anvil.BlocksPerRegion + int(l.CX)*anvil.BlocksPerChunk + int(l.BX)
	az := int(l.RZ)*anvil.BlocksPerRegion + int(l.CZ)*anvil.BlocksPerChunk + int(l.BZ)
	ay := int(l.BY)

	bx := int(b.RX)*anvil.BlocksPerRegion + int(b.CX)*anvil.BlocksPerChunk + int(b.BX)
	bz := int(b.RZ)*anvil.BlocksPerRegion + int(b.CZ)*anvil.BlocksPerChunk + int(b.BZ)
	by := int(b.BY)

	dx := float64(ax - bx)
	dy := float64(ay - by)
	dz := float64(az - bz)
	dist := math.Abs(math.Sqrt(dx*dx + dy*dy + dz*dz))
	return uint(dist)
}

// Block defines a block location in region, chunk and block coordinates.
type Block struct {
	Id item.Id // Type of block at this location.
	Location
}

func (b Block) String() string {
	return fmt.Sprintf("%s, %s", b.Id, b.Location)
}

// BlockList defines a set of blocks.
type BlockList []Block

func (s BlockList) Len() int { return len(s) }
