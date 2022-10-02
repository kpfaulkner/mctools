// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mctools

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/kpfaulkner/mctools/anvil"
)

// Known dimension names.
const (
	DimensionOverworld = "region"
	DimensionNether    = "DIM-1/region"
	DimensionEnd       = "DIM1/region"
)

// World defines a single Minecraft world.
type World struct {
	*anvil.Level                     // level.dat contents.
	root         string              // Directory with world data.
	regions      map[string][][2]int // List of known regions in this world - grouped by dimension.
}

// Open opens a new world in the given root directory.
func Open(root string) (*World, error) {
	var err error

	w := &World{
		root:    root,
		regions: make(map[string][][2]int),
	}

	// Load level.dat
	w.Level, err = anvil.LoadLevel(filepath.Join(root, "level.dat"))
	if err != nil {
		return nil, fmt.Errorf("mctools: load level.dat: %v", err)
	}

	// Find all regions in all dimensions.
	w.regions[DimensionOverworld] = listFiles(root, DimensionOverworld)
	w.regions[DimensionNether] = listFiles(root, DimensionNether)
	w.regions[DimensionEnd] = listFiles(root, DimensionEnd)
	return w, nil
}

// Save saves the level.dat information for this world.
func (w *World) Save() error {
	return w.Level.Save(filepath.Join(w.root, "level.dat"))
}

// Regions returns the coordinates for all regions in the world.
// This yields a map which groups region X/Z pairs for each dimension.
func (w *World) Regions() map[string][][2]int { return w.regions }

// DeleteRegion deletes the given region.
//
// If you have an open handle to this region, close it before calling this,
// as accessing its data afterwards will have undefined behaviour.
//
// Note that this permanently deletes the region file from disk.
// This operation can not be undone.
func (w *World) DeleteRegion(dim string, x, z int) error {
	file := w.regionFile(dim, x, z)
	err := os.Remove(file)

	if err != nil {
		return fmt.Errorf("mctools: delete region: %v", err)
	}

	// Update region list.
	w.regions[dim] = listFiles(w.root, dim)
	return nil
}

// CreateRegion creates a new region in this world and specified dimension,
// using the given coordinates.
// Returns an error if the region already exists.
func (w *World) CreateRegion(dim string, x, z int) (*anvil.Region, error) {
	file := w.regionFile(dim, x, z)
	region, err := anvil.CreateRegion(file)

	if err != nil {
		return nil, fmt.Errorf("mctools: create region: %v", err)
	}

	// Update region list.
	w.regions[dim] = listFiles(w.root, dim)
	return region, nil
}

// LoadRegion loads the given region in the specified dimension.
// Returns nil if the region could not be loaded.
func (w *World) LoadRegion(dim string, x, z int) (*anvil.Region, error) {
	file := w.regionFile(dim, x, z)
	region, err := anvil.LoadRegion(file)

	if err != nil {
		return nil, fmt.Errorf("mctools: load region %d.%d: %v", x, z, err)
	}

	return region, nil
}

// regionFile returns the full region file path for the given dimension and
// coordinates.
func (w *World) regionFile(dim string, x, z int) string {
	file := filepath.Join(w.root, dim)
	return filepath.Join(file, fmt.Sprintf("r.%d.%d.mca", x, z))
}

func listFiles(root, dimension string) [][2]int {
	fd, err := os.Open(filepath.Join(root, dimension))
	if err != nil {
		return nil
	}

	files, err := fd.Readdirnames(-1)
	fd.Close()

	if err != nil {
		return nil
	}

	out := make([][2]int, 0, len(files))

	for _, f := range files {
		if path.Ext(f) == anvil.RegionFileExtension {
			rx, rz, ok := anvil.RegionCoords(f)
			if ok {
				out = append(out, [2]int{rx, rz})
			}
		}
	}

	return out
}
