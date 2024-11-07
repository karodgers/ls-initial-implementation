package list

import (
	"fmt"
	"my-ls-1/internal/flags"
	"os"
	"strings"
)

// getBaseName returns the last element of a path
func getBaseName(path string) string {
	// Handle empty path
	if path == "" {
		return "."
	}

	// Remove trailing separators
	for len(path) > 0 && (path[len(path)-1] == '/' || path[len(path)-1] == '\\') {
		path = path[:len(path)-1]
	}

	// Find last separator
	lastSep := strings.LastIndexAny(path, "/\\")
	if lastSep == -1 {
		return path
	}
	return path[lastSep+1:]
}

// joinPath joins path elements with the system separator
func joinPath(elem ...string) string {
	result := strings.Join(elem, string(os.PathSeparator))
	// Clean up any double separators
	for strings.Contains(result, string(os.PathSeparator)+string(os.PathSeparator)) {
		result = strings.ReplaceAll(result, string(os.PathSeparator)+string(os.PathSeparator), string(os.PathSeparator))
	}
	return result
}

// ProcessPath handles the recursive listing of directories
func ProcessPath(path string, opts flags.Options) error {
	// Skip if the path itself is a hidden file/directory
	if !opts.ShowHidden && isHidden(getBaseName(path)) {
		return nil
	}

	// Process the current directory
	if err := listDirectory(path, opts); err != nil {
		return err
	}

	// If recursive flag is set, process subdirectories
	if opts.Recursive {
		return walkDirectory(path, opts)
	}
	return nil
}

// walkDirectory implements directory traversal without filepath.Walk
func walkDirectory(root string, opts flags.Options) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		path := joinPath(root, entry.Name())

		// Skip hidden files/directories unless ShowHidden is true
		if !opts.ShowHidden && isHidden(entry.Name()) {
			continue
		}

		if entry.IsDir() {
			fmt.Printf("\n%s:\n", path)
			if err := listDirectory(path, opts); err != nil {
				return err
			}
			// Recursively process subdirectory
			if err := walkDirectory(path, opts); err != nil {
				return err
			}
		}
	}
	return nil
}

// isHidden checks if a file or directory is hidden (starts with a dot)
func isHidden(name string) bool {
	return len(name) > 0 && name[0] == '.'
}
