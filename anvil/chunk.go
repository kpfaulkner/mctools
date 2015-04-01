// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package anvil

import "time"

// Known chunk compression schemes.
const (
	GZip = 1
	ZLib = 2
)

const (
	// MaxLight defines the highest possible light level.
	MaxLight = 15

	// SectionsPerChunk defines the maximum number of sections per chunk
	SectionsPerChunk = 16

	// MaxChunkHeight defines the maximum number of blocks a chunk can have
	// in its Y direction.
	MaxChunkHeight = SectionsPerChunk * BlocksPerChunk
)

// Tile Ticks represent block updates that need to happen because they could
// not happen before the chunk was saved. Examples reasons for tile ticks
// include redstone circuits needing to continue updating, water and lava
// that should continue flowing, recently placed sand or gravel that should
// fall, etc. Tile ticks are not used for purposes such as leaf decay, where
// the decay information is stored in the leaf block data values and handled
// by Minecraft when the chunk loads.
//
// For map makers, tile ticks can be used to update blocks after a period of
// time has passed with the chunk loaded into memory.
type TileTick struct {
	Id string `nbt:"i"`
	T  int32  `nbt:"t"`
	P  int32  `nbt:"p"`
	X  int32  `nbt:"x"`
	Y  int32  `nbt:"Y"`
	Z  int32  `nbt:"Z"`
}

// Chunk represents a single chunk in a region file.
//
// Reference: http://minecraft.gamepedia.com/Chunk_format
type Chunk struct {
	Entities         []Entity     `nbt:"Entities"`
	TileEntities     []TileEntity `nbt:"TileEntities"`
	TileTicks        []TileTick   `nbt:"TileTicks"`
	Sections         []Section    `nbt:"Sections"`
	Biomes           []int8       `nbt:"Biomes"`
	HeightMap        []int32      `nbt:"HeightMap"`
	LastUpdate       int64        `nbt:"LastUpdate"`
	InhabitedTime    int64        `nbt:"InhabitedTime"`
	X                int32        `nbt:"xPos"`
	Z                int32        `nbt:"zPos"`
	V                int8         `nbt:"V"`
	LightPopulated   bool         `nbt:"LightPopulated"`
	TerrainPopulated bool         `nbt:"TerrainPopulated"`
}

// Init initializes the chunk to a default, empty state.
// Writing in block data will create sections as necessary.
func (c *Chunk) Init(x, z int) {
	c.LastUpdate = time.Now().Unix()
	c.X = int32(x)
	c.Z = int32(z)
	c.InhabitedTime = 0
	c.LightPopulated = true
	c.TerrainPopulated = true
	c.Biomes = make([]int8, 256)
	c.HeightMap = make([]int32, 256)
	c.Sections = make([]Section, 0, 16)
	c.Entities = nil
	c.TileEntities = nil
	c.TileTicks = nil
	c.V = 1

	// Reset biomes.
	for i := range c.Biomes {
		c.Biomes[i] = 1
	}
}

// Section returns the section for the given Y coordinate.
// The coordinate is expected to be in the range 0-255 (MaxChunkHeight-1).
//
// If create is true, a new section is added if it doesn't yet exist,
// provided the y value is in an acceptable range.
//
// Returns nil if the coordinate is invalid.
// Returns nil if the section has not yet been generated and create is false.
func (c *Chunk) Section(y int, create bool) *Section {
	index := byte(y / SectionsPerChunk)

	for i := range c.Sections {
		if c.Sections[i].Y == index {
			return &c.Sections[i]
		}
	}

	if !create {
		return nil
	}

	sz := len(c.Sections)
	c.Sections = append(c.Sections, Section{})
	s := &c.Sections[sz]
	s.Init(index)
	return s
}

// UpdateHeightmap refills the heightmap with current block data.
// Each value in the heightmap records the lowest level in each column where
// the light from the sky is at full strength.
//
// This speeds computing of the SkyLight.
func (c *Chunk) UpdateHeightmap() {
	if len(c.HeightMap) == 0 {
		return
	}

	var x, y, z int
	var block Block

	for x = 0; x < BlocksPerChunk; x++ {
		for z = 0; z < BlocksPerChunk; z++ {
			for y = MaxChunkHeight - 1; y > 0; y-- {
				s := c.Section(y, false)
				if s == nil {
					continue
				}

				if s.Read(x, y%SectionsPerChunk, z, &block) && block.SkyLight < MaxLight {
					break
				}
			}

			c.HeightMap[z*BlocksPerChunk+x] = int32(y)
		}
	}
}
