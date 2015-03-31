// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package anvil

import "github.com/jteeuwen/mctools/anvil/item"

// Section defines a section of blocks.
//
// Each chunk is divided up into 16 equal sections.
// Only generated sections will be saved to the world file.
// This is done to save file space. Each section spans 16*16*16 blocks.
type Section struct {
	Blocks     []uint8 `nbt:"Blocks"`        // Primary block IDs -- 8 bits per block.
	Add        []uint8 `nbt:"Add,omitempty"` // Optional extra block ID information -- 4 bits per block.
	Data       []uint8 `nbt:"Data"`          // Block data -- 4 bits per block.
	BlockLight []uint8 `nbt:"BlockLight"`    // Amount of block-emitted light in each block -- 4 bits per block.
	SkyLight   []uint8 `nbt:"SkyLight"`      // Amount of sunlight or moonlight hitting each block -- 4 bits per block.
	Y          byte    `nbt:"Y"`             // Y index for this section.
}

// Init initializes the section to default, empty settings.
func (s *Section) Init(y byte) {
	if len(s.Blocks) > 0 {
		return // Already initializes.
	}

	s.Y = y
	s.Add = nil
	s.Blocks = make([]uint8, 4096)
	s.Data = make([]uint8, 2048)
	s.BlockLight = make([]uint8, 2048)
	s.SkyLight = make([]uint8, 2048)

	// Set the skylight to highest light level.
	for i := range s.SkyLight {
		s.SkyLight[i] = MaxLight
	}
}

// Write stores the given block struct for the specified coordinates.
//
// Returns false if the coordinates are out of range.
func (s *Section) Write(x, y, z int, b *Block) bool {
	y %= SectionsPerChunk
	index := y*16*16 + z*16 + x

	if index < 0 || index >= len(s.Blocks) {
		return false
	}

	s.Blocks[index] = uint8(b.Id)

	if add := uint8(b.Id >> 8); add > 0 {
		if len(s.Add) == 0 {
			s.Add = make([]uint8, 2048)
		}

		snibble(s.Add, index, add)
	}

	snibble(s.Data, index, uint8(b.Id>>16))
	snibble(s.BlockLight, index, b.BlockLight)
	snibble(s.SkyLight, index, b.SkyLight)
	return true
}

// Read fills the given block struct with data at the specified coordinates.
//
// Returns false if the block data could not be found. This can happen
// when the coordinates are out of range.
func (s *Section) Read(x, y, z int, b *Block) bool {
	y %= SectionsPerChunk
	index := y*16*16 + z*16 + x

	if index < 0 || index >= len(s.Blocks) {
		return false
	}

	b.Id = item.Id(s.Blocks[index])

	if len(s.Add) > 0 {
		b.Id |= item.Id(gnibble(s.Add, index)) << 8
	}

	b.Id |= item.Id(gnibble(s.Data, index)) << 16

	b.BlockLight = gnibble(s.BlockLight, index)
	b.SkyLight = gnibble(s.SkyLight, index)
	return true
}

// gnibble returns either upper or lower 4-bits for a given index.
func gnibble(arr []uint8, index int) uint8 {
	if index%2 == 0 {
		return arr[index/2] & 0xf
	}

	return (arr[index/2] >> 4) & 0x0f
}

// snibble sets either upper or lower 4-bits for a given index.
func snibble(arr []uint8, index int, v uint8) {
	if index%2 == 0 {
		arr[index/2] = (arr[index/2] & 0xf0) | (v & 0xf)
	} else {
		arr[index/2] = (arr[index/2] & 0xf) | (v&0xf)<<4
	}
}
