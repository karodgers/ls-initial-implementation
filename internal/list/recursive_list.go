package list

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"my-ls-1/internal/flags"
)

// ProcessPath handles the recursive listing of directories
func ProcessPath(path string, opts flags.Options) error {
	// Skip if the path itself is a hidden file/directory
	if !opts.ShowHidden && isHidden(filepath.Base(path)) {
		return nil
	}

	// Process the current directory
	if err := listDirectory(path, opts); err != nil {
		return err
	}

	// If recursive flag is set, process subdirectories
	if opts.Recursive {
		return filepath.Walk(path, func(p string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip the root directory as it's already processed
			if p == path {
				return nil
			}

			// Skip hidden files/directories unless ShowHidden is true
			if !opts.ShowHidden && isHidden(info.Name()) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			if info.IsDir() {
				fmt.Printf("\n%s:\n", p)
				return listDirectory(p, opts)
			}
			return nil
		})
	}

	return nil
}

// isHidden checks if a file or directory is hidden (starts with a dot)
func isHidden(name string) bool {
	return len(name) > 0 && name[0] == '.'
}
