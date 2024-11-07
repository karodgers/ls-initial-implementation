package list

// SortByTime sorts files by modification time (newest first)
// Uses bubble sort algorithm
func SortByTime(files []FileInfo) {
	n := len(files)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements
			if files[j].ModTime.Before(files[j+1].ModTime) {
				// Swap them if they are in the wrong order
				files[j], files[j+1] = files[j+1], files[j]
			}
		}
	}
}
