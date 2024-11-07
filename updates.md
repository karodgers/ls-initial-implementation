remove sort package in list
import (
    "fmt"
    "my-ls-1/internal/flags"
    "strings"
)

func listDirectory(path string, opts flags.Options) error {
    // Get directory contents
    files, err := GetDirectoryContents(path, opts.ShowAll)
    if err != nil {
        return err
    }

    // Bubble sort implementation for files
    n := len(files)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-1-i; j++ {
            // Compare file names case-insensitively
            if strings.ToLower(files[j].Name) > strings.ToLower(files[j+1].Name) {
                // Swap files
                files[j], files[j+1] = files[j+1], files[j]
            }
        }
    }

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