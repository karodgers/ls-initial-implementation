package main

import (
	"fmt"
	"os"
	
	"my-ls-1/internal/flags"
	"my-ls-1/internal/list"
)

func main() {
	// Parse command line arguments (skip program name)
	opts, err := flags.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Process each target path
	for _, path := range opts.TargetPaths {
		if err := list.ProcessPath(path, opts); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", path, err)
		}
	}
}