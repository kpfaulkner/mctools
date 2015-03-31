// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jteeuwen/mctools/anvil"
)

func main() {
	var err error
	file := parseArgs()

	// Region files need some extra processing,
	if strings.EqualFold(filepath.Ext(file), anvil.RegionFileExtension) {
		err = dumpRegion(os.Stdout, file)
	} else {
		err = dumpFile(os.Stdout, file)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// parseArgs parses and validates command line arguments.
func parseArgs() string {
	flag.Usage = func() {
		fmt.Println("usage:", os.Args[0], "[options] <file>")
		flag.PrintDefaults()
	}

	version := flag.Bool("version", false, "Display version information.")
	flag.Parse()

	if *version {
		fmt.Println(Version())
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	return flag.Arg(0)
}
