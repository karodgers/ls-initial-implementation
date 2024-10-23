package list

import (
	"fmt"
	"my-ls-1/internal/flags"
	"sort"
	"strings"
)

func listDirectory(path string, opts flags.Options) error {
	// Get directory contents
	files, err := GetDirectoryContents(path, opts.ShowAll)
	if err != nil {
		return err
	}

	// Sort files
	sort.Slice(files, func(i, j int) bool {
		return strings.ToLower(files[i].Name) < strings.ToLower(files[j].Name)
	})

	// Apply time sort if requested
	if opts.SortByTime {
		SortByTime(files)
	}

	// Apply reverse sort if requested
	if opts.ReverseSort {
		ReverseSort(files)
	}

	// Display files
	for _, file := range files {
		if opts.LongFormat {
			fmt.Println(FormatLongList(file))
		} else {
			fmt.Println(file.Name)
		}
	}

	return nil
}
