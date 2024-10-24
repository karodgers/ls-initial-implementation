package flags

import (
	"errors"
	"strings"
)

type Options struct {
	LongFormat  bool // -l flag
	Recursive   bool // -R flag
	ShowAll     bool // -a flag
	ReverseSort bool // -r flag
	SortByTime  bool // -t flag
	ShowHidden  bool
	TargetPaths []string
}

// ParseFlags parses command line arguments and returns Options
func ParseFlags(args []string) (Options, error) {
	opts := Options{
		TargetPaths: []string{"."}, // Default to current directory
	}

	if len(args) == 0 {
		return opts, nil
	}

	// Process all arguments
	for i := 0; i < len(args); i++ {
		arg := args[i]

		// Handle flags
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			// Remove the "-" prefix
			flags := strings.TrimPrefix(arg, "-")

			// Process each character in the flag string
			for _, flag := range flags {
				switch flag {
				case 'l':
					opts.LongFormat = true
				case 'R':
					opts.Recursive = true
				case 'a':
					opts.ShowAll = true
				case 'r':
					opts.ReverseSort = true
				case 't':
					opts.SortByTime = true
				default:
					return opts, errors.New("invalid flag: " + string(flag))
				}
			}
		} else {
			// If not a flag, treat as a target path
			opts.TargetPaths = append(opts.TargetPaths, arg)
		}
	}

	return opts, nil
}
