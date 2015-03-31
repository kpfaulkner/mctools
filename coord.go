// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mctools

import (
	"math"

	"github.com/jteeuwen/mctools/anvil"
)

// RegionCoords returns region coordinates from the given,
// absolute block position.
func RegionCoords(x, z int) (int, int) {
	rx := math.Floor(float64(x) / anvil.BlocksPerRegion)
	rz := math.Floor(float64(z) / anvil.BlocksPerRegion)
	return int(rx), int(rz)
}

// ChunkCoords returns the chunk coordinates for the given,
// absolute block position.
func ChunkCoords(x, z int) (int, int) {
	cx := math.Floor(float64(x) / anvil.BlocksPerChunk)
	cz := math.Floor(float64(z) / anvil.BlocksPerChunk)
	return int(cx), int(cz)
}

// BlockCoords returns block offsets in a chunk for the given,
// absolute block position.
func BlockCoords(x, y, z int) (int, int, int) {
	x = x % anvil.BlocksPerChunk
	z = z % anvil.BlocksPerChunk

	if x < 0 {
		x = anvil.BlocksPerChunk + x
	}

	if z < 0 {
		z = anvil.BlocksPerChunk + z
	}

	if y < 0 {
		y = 0
	}

	if y >= anvil.MaxChunkHeight {
		y = anvil.MaxChunkHeight - 1
	}

	return x, y, z
}
