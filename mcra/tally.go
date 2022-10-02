// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mcra

import (
	"github.com/kpfaulkner/mctools/anvil"
	"github.com/kpfaulkner/mctools/anvil/item"
)

// TallyResult defines a table of unique item ids and the number of
// times they occur in a specific region or chunk.
type TallyResult map[item.Id]uint64

// TallyInRegion counts the number of times each of the given items occurs
// in the specified region.
//
// If the given item set is empty, all blocks will be counted.
func TallyInRegion(r *anvil.Region, items ...item.Id) TallyResult {
	var chunk anvil.Chunk
	out := make(TallyResult)

	for _, xz := range r.Chunks() {
		if !r.ReadChunk(xz[0], xz[1], &chunk) {
			continue
		}

		tallyInChunk(&chunk, items, out)
	}

	return out
}

// TallyInChunk counts the number of times each of the given items occurs
// in the specified chunk.
//
// If the given item set is empty, all blocks will be counted.
func TallyInChunk(c *anvil.Chunk, items ...item.Id) TallyResult {
	out := make(TallyResult)
	tallyInChunk(c, items, out)
	return out
}

func tallyInChunk(c *anvil.Chunk, items []item.Id, out TallyResult) {
	for i := range c.Sections {
		tallyInSection(&c.Sections[i], items, out)
	}
}

func tallyInSection(s *anvil.Section, items []item.Id, out TallyResult) {
	var x, y, z int
	var block anvil.Block

	for y = 0; y < 16; y++ {
		for x = 0; x < anvil.BlocksPerChunk; x++ {
			for z = 0; z < anvil.BlocksPerChunk; z++ {
				if s.Read(x, y, z, &block) && hasItem(items, block.Id) {
					out[block.Id]++
				}
			}
		}
	}
}

// hasItem returns true if set contains v.
func hasItem(set []item.Id, v item.Id) bool {
	if len(set) == 0 {
		return true
	}

	for _, id := range set {
		if id == v {
			return true
		}
	}

	return false
}
