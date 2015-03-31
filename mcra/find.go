// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mcra

import (
	"github.com/jteeuwen/mctools/anvil"
	"github.com/jteeuwen/mctools/anvil/item"
)

// FindStrongholds finds all strongholds by locating End Portal blocks.
func FindStrongholds(r *anvil.Region) BlockList {
	return FindInRegion(r, NewInclusionQuery(
		item.EndPortal,
		item.EndPortalFrame,
	))
}

// FindDungeons finds all dungeons by locating and returning all mob spawners.
func FindDungeons(r *anvil.Region) []*anvil.TileEntity {
	var out []*anvil.TileEntity
	var chunk anvil.Chunk

	for _, xz := range r.Chunks() {
		if !r.ReadChunk(xz[0], xz[1], &chunk) {
			continue
		}

		for i := range chunk.TileEntities {
			if chunk.TileEntities[i].Id == "MobSpawner" {
				out = append(out, &chunk.TileEntities[i])
			}
		}
	}

	return out
}

// FindInRegion locates all blocks in the specified region,
// matching the given query.
func FindInRegion(r *anvil.Region, q Query) BlockList {
	var out BlockList
	var chunk anvil.Chunk
	var loc Block

	loc.RX = int8(r.X)
	loc.RZ = int8(r.Z)

	for _, xz := range r.Chunks() {
		if !r.ReadChunk(xz[0], xz[1], &chunk) {
			continue
		}

		findInChunk(&chunk, q, loc, &out)
	}

	return out
}

// FindInChunk locates all blocks in the specified chunk,
// matching the given query.
func FindInChunk(c *anvil.Chunk, q Query) BlockList {
	var out BlockList
	var loc Block

	loc.CX = int8(c.X)
	loc.CZ = int8(c.Z)

	findInChunk(c, q, loc, &out)
	return out
}

func findInChunk(c *anvil.Chunk, q Query, loc Block, out *BlockList) {
	for i := range c.Sections {
		loc.BY = c.Sections[i].Y * 16
		findInSection(&c.Sections[i], q, loc, out)
	}
}

func findInSection(s *anvil.Section, q Query, loc Block, out *BlockList) {
	var x, y, z int
	var block anvil.Block

	sy := loc.BY
	slice := *out

	for y = 0; y < 16; y++ {
		for x = 0; x < anvil.BlocksPerChunk; x++ {
			for z = 0; z < anvil.BlocksPerChunk; z++ {
				loc.BX = uint8(x)
				loc.BY = sy + uint8(y)
				loc.BZ = uint8(z)
				loc.Id = block.Id

				if s.Read(x, y, z, &block) && q.IsTarget(loc) {
					slice = append(slice, loc)
				}
			}
		}
	}

	*out = slice
}
