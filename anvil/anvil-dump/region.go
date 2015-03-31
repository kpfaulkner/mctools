// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

// sectorSize defines the byte size of a single sector.
const sectorSize = 4096

// dumpRegion reads a region file and extracts its NBT encoded chunk descriptors.
func dumpRegion(w io.Writer, file string) error {
	var locations [sectorSize]byte

	fd, err := os.Open(file)
	if err != nil {
		return err
	}

	defer fd.Close()

	_, err = io.ReadFull(fd, locations[:])
	if err != nil {
		return err
	}

	for i := 0; i < len(locations); i += 4 {
		d := locations[i:]
		offset := int64(d[0])<<16 | int64(d[1])<<8 | int64(d[2])
		sectors := d[3]

		if offset == 0 && sectors == 0 {
			continue
		}

		err := dumpChunk(w, fd, offset)
		if err != nil {
			return err
		}
	}

	return nil
}

// dumpChunk extracts a compressed chunk from the given reader and
// dumps its NBT tag contents.
func dumpChunk(w io.Writer, r io.ReadSeeker, offset int64) error {
	address := offset * sectorSize
	_, err := r.Seek(address, 0)
	if err != nil {
		return err
	}

	var header [5]byte
	_, err = io.ReadFull(r, header[:])
	if err != nil {
		return err
	}

	var rr io.ReadCloser
	switch header[4] {
	case 1:
		rr, _ = gzip.NewReader(r)
	case 2:
		rr, _ = zlib.NewReader(r)
	default:
		return fmt.Errorf("chunk(%d); invalid compression scheme: %d", offset, header[4])
	}

	defer rr.Close()

	fmt.Fprintf(w, "Chunk %d:\n", offset)
	return dump(w, rr)
}
