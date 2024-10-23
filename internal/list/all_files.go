package list

import (
	"os"
	"strings"
)

// GetDirectoryContents returns a list of files in the directory
func GetDirectoryContents(path string, showHidden bool) ([]FileInfo, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Read all entries
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	var files []FileInfo
	for _, name := range names {
		// Skip hidden files unless showHidden is true
		if !showHidden && strings.HasPrefix(name, ".") {
			continue
		}

		info, err := GetFileInfo(path, name)
		if err != nil {
			continue
		}
		files = append(files, info)
	}

	return files, nil
}
