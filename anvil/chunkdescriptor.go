// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package anvil

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"math"
	"time"

	"github.com/jteeuwen/mctools/anvil/nbt"
)

// ChunkDescriptor describes a single chunk as stored in a region.
type ChunkDescriptor struct {
	data         []byte    // Compressed chunk data.
	LastModified time.Time // Last time thischunk was modified.
	X, Z         int       // Chunk coordinates in region.
	scheme       byte      // Compression scheme.
}

// Empty returns true if this chunk has not been generated yet.
func (cd *ChunkDescriptor) Empty() bool { return len(cd.data) == 0 }

// SectorCount returns the number of sectors this chunk occupies.
func (cd *ChunkDescriptor) SectorCount() int {
	return int(math.Ceil(float64(len(cd.data)) / sectorSize))
}

// Read decompresses chunk data into the given structure.
// Returns false if ther eis no data or the decompression failed.
func (cd *ChunkDescriptor) Read(c *Chunk) bool {
	var r io.ReadCloser
	var err error

	buf := bytes.NewBuffer(cd.data)

	switch cd.scheme {
	case GZip:
		r, err = gzip.NewReader(buf)
	case ZLib:
		r, err = zlib.NewReader(buf)
	default:
		return false
	}

	if err != nil {
		return false
	}

	// Clear out existing data; the nbt decoder will append to the existing slices.
	c.Sections = nil
	c.Biomes = nil
	c.HeightMap = nil
	c.Entities = nil
	c.TileEntities = nil
	c.TileTicks = nil

	var v struct {
		Level *Chunk
	}
	v.Level = c

	err = nbt.Unmarshal(r, &v)
	r.Close()
	return err == nil
}

// Write compresses the given chunk and writes the data into the current
// chunk descriptor.
func (cd *ChunkDescriptor) Write(c *Chunk) bool {
	cd.LastModified = time.Now()
	c.LastUpdate = cd.LastModified.Unix()

	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)

	var v struct {
		Level *Chunk
	}
	v.Level = c

	err := nbt.Marshal(w, v)
	w.Close()

	cd.data = buf.Bytes()
	return err == nil
}
