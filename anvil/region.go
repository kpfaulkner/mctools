// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package anvil

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Reference: http://minecraft.gamepedia.com/Region_file_format

// These values are here to aid in converting various coordinate
// sets between one another: block->chunk, block->region, etc.
const (
	ChunksPerRegion = 32
	BlocksPerChunk  = 16
	BlocksPerRegion = BlocksPerChunk * ChunksPerRegion

	// RegionFileExtension defines the file extension for region files.
	RegionFileExtension = ".mca"

	// Defines the byte size of a single sector.
	sectorSize = 4096
)

// RegionCoords returns the x and z coordinates associated with the
// given region file.
//
// The name should be in the form 'r.<X>.<Z>.mca', where X and Z are
// the signed integer coordinates.
//
// Returns false if the coordinates could not be determined.
func RegionCoords(name string) (int, int, bool) {
	_, name = filepath.Split(name)

	elem := strings.Split(name, ".")
	if len(elem) != 4 {
		return 0, 0, false
	}

	x, ex := strconv.ParseInt(elem[1], 10, 32)
	z, ez := strconv.ParseInt(elem[2], 10, 32)
	return int(x), int(z), ex == nil && ez == nil
}

// A region describes chunks with block data in a Minecraft world.
type Region struct {
	file   string                // Input file for this region.
	chunks [1024]ChunkDescriptor // Chunk definitions in this region.
	X      int                   // Region's X coordinate.
	Z      int                   // Region's Z coordinate.
}

// CreateRegion creates an empty region file at the given location.
// This contains only an empty header, without any chunks.
// Returns an error if the file already exists.
func CreateRegion(file string) (*Region, error) {
	_, _, ok := RegionCoords(file)
	if !ok {
		return nil, fmt.Errorf("anvil: create region: invalid file name")
	}

	fd, err := os.Create(file)
	if err != nil {
		return nil, fmt.Errorf("anvil: create region: %v", err)
	}

	var buf [sectorSize * 2]byte
	_, err = fd.Write(buf[:])
	if err != nil {
		fd.Close()
		return nil, fmt.Errorf("anvil: create region: %v", err)
	}

	fd.Close()

	return LoadRegion(file)
}

// LoadRegion opens a region from the given file.
func LoadRegion(file string) (*Region, error) {
	var err error

	rx, rz, ok := RegionCoords(file)
	if !ok {
		return nil, fmt.Errorf("anvil: open region: invalid file %q", file)
	}

	fd, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("anvil: r(%d %d): %v", rx, rz, err)
	}

	defer fd.Close()

	// Read header data.
	locations, timestamps, err := readHeader(fd)
	if err != nil {
		return nil, fmt.Errorf("anvil: r(%d %d): read header: %v", rx, rz, err)
	}

	r := &Region{
		file: file,
		X:    rx,
		Z:    rz,
	}

	// Load up all valid chunk descriptors.
	for x := 0; x < ChunksPerRegion; x++ {
		for z := 0; z < ChunksPerRegion; z++ {
			offset, sectors := readOffset(locations, x, z)

			// If either are zero, this chunk has not been generated yet.
			if offset == 0 && sectors == 0 {
				continue
			}

			n := chunkIndex(x, z)
			r.chunks[n], err = readChunk(fd, x, z, offset, timestamps)
			if err != nil {
				return nil, fmt.Errorf("anvil: r(%d %d) c(%d %d): read chunk: %v",
					rx, rz, x, z, err)
			}
		}
	}

	return r, nil
}

// Save writes all region data to the underlying file.
func (r *Region) Save() error {
	fd, err := os.Create(r.file)
	if err != nil {
		return fmt.Errorf("anvil: r(%d %d): %v", r.X, r.Z, err)
	}

	defer fd.Close()

	err = writeHeader(fd, r.chunks[:])
	if err != nil {
		return fmt.Errorf("anvil: r(%d %d): %v", r.X, r.Z, err)
	}

	offset := 2 // Skip first two offsets for header data.

	for _, cd := range r.chunks {
		err = writeChunk(fd, cd, offset)
		if err != nil {
			return fmt.Errorf("anvil: r(%d %d) c(%d %d): %v", r.X, r.Z, cd.X, cd.Z, err)
		}

		offset += cd.SectorCount()
	}

	return nil
}

// Clear removes all blocks and all chunks from the region.
// Note that Region.Save() must be called to persist these changes.
func (r *Region) Clear() {
	for i := range r.chunks {
		r.chunks[i].X = 0
		r.chunks[i].Z = 0
		r.chunks[i].scheme = 0
		r.chunks[i].data = nil
	}
}

// ChunkLen returns the number of /valid/ chunks in this region.
// Meaning chunks which have actually been generated and have data.
func (r *Region) ChunkLen() int {
	var count int

	for _, cd := range r.chunks {
		if !cd.Empty() {
			count++
		}
	}

	return count
}

// Chunks yields a list of chunk X/Z cooridnates for all valid chunks in
// this region. There are a maximum of 1024 chunks per region.
func (r *Region) Chunks() [][2]int {
	out := make([][2]int, 0, len(r.chunks))

	for _, cd := range r.chunks {
		if !cd.Empty() {
			out = append(out, [2]int{cd.X, cd.Z})
		}
	}

	return out
}

// HasChunk returns true if the given chunk exists in this region.
func (r *Region) HasChunk(x, z int) bool {
	n := chunkIndex(x, z)
	return !r.chunks[n].Empty()
}

// ReadChunk reads chunk data for the given coordinates into the specified
// structure.
//
// Returns false if there is no valid chunk available, or the chunk data can
// not be decompressed.
func (r *Region) ReadChunk(x, z int, c *Chunk) bool {
	n := chunkIndex(x, z)

	if !r.chunks[n].Empty() {
		return r.chunks[n].Read(c)
	}

	return false
}

// WriteChunk writes compresses the given chunk data, so it may later be
// persisted using Region.Save().
func (r *Region) WriteChunk(x, z int, c *Chunk) bool {
	n := chunkIndex(x, z)
	return r.chunks[n].Write(c)
}

// writeHeader writes header data into the given writer.
func writeHeader(w io.WriteSeeker, set []ChunkDescriptor) error {
	var locations, timestamps [sectorSize]byte

	offset := 2 // Skip first two sectors for the header.

	for _, cd := range set {
		if cd.Empty() {
			continue
		}

		sectors := cd.SectorCount()
		writeOffset(locations[:], cd.X, cd.Z, offset, sectors)
		offset += sectors

		writeTimestamp(timestamps[:], cd.X, cd.Z, cd.LastModified)
	}

	_, err := w.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = w.Write(locations[:])
	if err != nil {
		return err
	}

	_, err = w.Write(timestamps[:])
	return err
}

// readHeader reads header data from the given stream.
func readHeader(r io.ReadSeeker) ([]byte, []byte, error) {
	var locations, timestamps [sectorSize]byte

	_, err := r.Seek(0, 0)
	if err != nil {
		return nil, nil, err
	}

	// Read chunk location data.
	_, err = io.ReadFull(r, locations[:])
	if err != nil {
		return nil, nil, err
	}

	// Read chunk access timestamps.
	_, err = io.ReadFull(r, timestamps[:])
	if err != nil {
		return nil, nil, err
	}

	return locations[:], timestamps[:], nil
}

// writeChunk writes a chunk to the given stream.
func writeChunk(w io.WriteSeeker, cd ChunkDescriptor, offset int) error {
	if cd.Empty() {
		return nil
	}

	// Jump to chunk sector.
	_, err := w.Seek(int64(offset)*sectorSize, 0)
	if err != nil {
		return err
	}

	// Write compressed data size.
	err = writeU32(w, uint32(len(cd.data)))
	if err != nil {
		return err
	}

	// Write compression scheme.
	err = writeU8(w, ZLib)
	if err != nil {
		return err
	}

	// Write compressed data.
	_, err = w.Write(cd.data)
	return err
}

// readChunk reads a chunk from the given stream.
func readChunk(r io.ReadSeeker, x, z, offset int, timestamps []byte) (ChunkDescriptor, error) {
	var cd ChunkDescriptor
	cd.X = x
	cd.Z = z
	cd.LastModified = readTimestamp(timestamps, x, z)

	// Jump to chunk sector.
	_, err := r.Seek(int64(offset)*sectorSize, 0)
	if err != nil {
		return cd, err
	}

	// Read compressed data size.
	size, err := readU32(r)
	if err != nil {
		return cd, err
	}

	// Read compression scheme.
	cd.scheme, err = readU8(r)
	if err != nil {
		return cd, err
	}

	// Read compressed data.
	cd.data = make([]byte, size)
	_, err = io.ReadFull(r, cd.data)
	return cd, err
}

// readTimestamp returns the time at which the given chunk was last modified.
// Returns the zero time value of the data could not be found.
func readTimestamp(data []byte, x, z int) time.Time {
	n := chunkIndex(x, z)
	d := data[n*4:]
	stamp := int64(d[0])<<24 | int64(d[1])<<16 | int64(d[2])<<8 | int64(d[3])
	return time.Unix(stamp, 0)
}

// writeTimestamp sets the time at which the given chunk was last modified.
func writeTimestamp(data []byte, x, z int, t time.Time) {
	stamp := uint32(t.Unix())
	n := chunkIndex(x, z)
	d := data[n*4:]

	d[0] = byte(stamp >> 24)
	d[1] = byte(stamp >> 16)
	d[2] = byte(stamp >> 8)
	d[3] = byte(stamp)
}

// readOffset returns the address offset and sector count for the given
// chunk's data.
//
// The first value is the offset (in 4KiB sectors) from the start of the file.
// The second is the length of the chunk (in 4KiB sectors).
//
// Both values will be 0 if the given chunk is not present in the region.
// This is the case when it has not been generated yet.
func readOffset(data []byte, x, z int) (offset, sectors int) {
	n := chunkIndex(x, z)
	d := data[n*4:]
	return int(d[0])<<16 | int(d[1])<<8 | int(d[2]), int(d[3])
}

// writeOffset sets the address offset for the given chunk's data.
//
// The first value is the offset (in 4KiB sectors) from the start of the file.
// The second is the length of the chunk (in 4KiB sectors).
func writeOffset(data []byte, x, z, offset, sectors int) {
	n := chunkIndex(x, z)
	d := data[n*4:]

	d[0] = byte(offset >> 16)
	d[1] = byte(offset >> 8)
	d[2] = byte(offset)
	d[3] = byte(sectors)
}

func readU32(r io.Reader) (uint32, error) {
	var v uint32
	err := binary.Read(r, binary.BigEndian, &v)
	return v, err
}

func writeU32(w io.Writer, v uint32) error {
	return binary.Write(w, binary.BigEndian, v)
}

func readU8(r io.Reader) (uint8, error) {
	var v uint8
	err := binary.Read(r, binary.BigEndian, &v)
	return v, err
}

func writeU8(w io.Writer, v uint8) error {
	return binary.Write(w, binary.BigEndian, v)
}

// chunkIndex computes the offset in the Locations and Timestamos
// arrays for a chunk at coordinates x,z.
func chunkIndex(x, z int) int {
	x %= 32
	z %= 32

	if x < 0 {
		x += 32
	}

	if z < 0 {
		z += 32
	}

	return (x + z*32)
}
