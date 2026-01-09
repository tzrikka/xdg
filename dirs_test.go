package xdg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAbsDirExists(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "dir1"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "file1"), nil, NewFilePermissions); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "existent_directory",
			path: "dir1",
			want: true,
		},
		{
			name: "existent_file",
			path: "file1",
			want: false,
		},
		{
			name: "non_existent_directory",
			path: "dir2",
			want: false,
		},
		{
			name: "non_existent_file",
			path: "file2",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := filepath.Join(root, tt.path)
			if got := absDirExists(p); got != tt.want {
				t.Errorf("absDirExists(%s) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
