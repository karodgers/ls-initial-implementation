package list

import (
	"sort"
)

// SortByTime sorts files by modification time (newest first)
func SortByTime(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime.After(files[j].ModTime)
	})
}