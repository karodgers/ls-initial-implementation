package list

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"
)

// FileInfo holds the detailed information about a file
type FileInfo struct {
	Name    string
	Mode    string
	Links   int
	User    string
	Group   string
	Size    int64
	ModTime time.Time
	IsDir   bool
	Path    string
}

// FormatLongList formats file information in the long listing format (-l flag)
func FormatLongList(info FileInfo) string {
	// Format similar to: -rw-r--r-- 1 user group 123 Jan 1 12:34 filename
	timeStr := formatTime(info.ModTime)
	return fmt.Sprintf("%s %3d %-8s %-8s %8d %s %s",
		info.Mode,
		info.Links,
		info.User,
		info.Group,
		info.Size,
		timeStr,
		info.Name)
}

// GetFileInfo retrieves detailed file information
func GetFileInfo(path string, name string) (FileInfo, error) {
	fullPath := path + "/" + name
	stat, err := os.Lstat(fullPath)
	if err != nil {
		return FileInfo{}, err
	}

	sys := stat.Sys().(*syscall.Stat_t)

	// Get user and group info
	usr, err := user.LookupId(strconv.Itoa(int(sys.Uid)))
	if err != nil {
		usr = &user.User{Username: strconv.Itoa(int(sys.Uid))}
	}

	group, err := user.LookupGroupId(strconv.Itoa(int(sys.Gid)))
	if err != nil {
		group = &user.Group{Name: strconv.Itoa(int(sys.Gid))}
	}

	return FileInfo{
		Name:    name,
		Mode:    formatMode(stat.Mode()),
		Links:   int(sys.Nlink),
		User:    usr.Username,
		Group:   group.Name,
		Size:    stat.Size(),
		ModTime: stat.ModTime(),
		IsDir:   stat.IsDir(),
		Path:    fullPath,
	}, nil
}

// formatMode converts file mode bits to string representation
func formatMode(mode os.FileMode) string {
	modeStr := mode.String()
	if len(modeStr) > 10 {
		modeStr = modeStr[len(modeStr)-10:]
	}
	return modeStr
}

// formatTime formats the modification time similar to ls
func formatTime(t time.Time) string {
	now := time.Now()
	sixMonthsAgo := now.AddDate(0, -6, 0)

	if t.Before(sixMonthsAgo) || t.After(now) {
		return t.Format("Jan _2  2006")
	}
	return t.Format("Jan _2 15:04")
}
