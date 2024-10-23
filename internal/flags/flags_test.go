package flags

import (
	"reflect"
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Options
		wantErr bool
	}{
		{
			name: "no flags",
			args: []string{},
			want: Options{TargetPaths: []string{"."}},
			wantErr: false,
		},
		{
			name: "all flags",
			args: []string{"-lRart"},
			want: Options{
				LongFormat:  true,
				Recursive:   true,
				ShowAll:     true,
				ReverseSort: true,
				SortByTime:  true,
				TargetPaths: []string{"."},
			},
			wantErr: false,
		},
		{
			name: "invalid flag",
			args: []string{"-z"},
			want: Options{},
			wantErr: true,
		},
		{
			name: "flags with path",
			args: []string{"-l", "testdir"},
			want: Options{
				LongFormat:  true,
				TargetPaths: []string{".", "testdir"},
			},
			wantErr: false,
		},
		{
			name: "multiple paths",
			args: []string{"dir1", "dir2"},
			want: Options{
				TargetPaths: []string{".", "dir1", "dir2"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFlags(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}