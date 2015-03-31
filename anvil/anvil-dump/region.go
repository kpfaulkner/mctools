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

		fmt.Fprintf(w, "Chunk %d (offset: %d, sectors: %d):\n", i/4, offset, sectors)
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
	address := offset*sectorSize + 4
	_, err := r.Seek(address, 0)
	if err != nil {
		return err
	}

	var scheme [1]byte
	_, err = io.ReadFull(r, scheme[:])
	if err != nil {
		return err
	}

	var rr io.ReadCloser
	switch scheme[0] {
	case 1:
		rr, err = gzip.NewReader(r)
	case 2:
		rr, err = zlib.NewReader(r)
	default:
		return fmt.Errorf("chunk(%d); invalid compression scheme: %d", offset, scheme[0])
	}

	if err != nil {
		return err
	}

	err = dump(w, rr)
	rr.Close()
	return err
}
