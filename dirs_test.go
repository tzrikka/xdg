package xdg

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestConfigHome(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		want    string
		wantErr bool
	}{
		{
			name: "empty_env_var",
			want: defaultConfigHome(),
		},
		{
			name: "valid_env_var",
			env:  "/absolute/path",
			want: "/absolute/path",
		},
		{
			name:    "relative_env_var",
			env:     "relative/path",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("XDG_CONFIG_HOME", tt.env)

			got, err := ConfigHome()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigHome() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ConfigHome() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConfigDirs(t *testing.T) {
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "dir"), 0o700); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		env     string
		want    []string
		wantErr bool
	}{
		{
			name: "valid_env_var",
			env: strings.Join([]string{
				"/absolute/path1",
				"/absolute/path2",
				filepath.Join(dir, "dir"),
			}, listSeparator),
			want: []string{filepath.Join(dir, "dir")},
		},
		{
			name: "relative_env_var",
			env: strings.Join([]string{
				"relative/path1",
				"relative/path2",
			}, listSeparator),
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("XDG_CONFIG_DIRS", tt.env)

			got, err := ConfigDirs()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigDirs() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(got) != len(tt.want) {
				t.Errorf("ConfigDirs() = %v, want %v", got, tt.want)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("ConfigDirs() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAbsDirExists(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "dir1"), 0o700); err != nil {
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
			name: "relative_path",
			path: "relative/path",
			want: false,
		},
		{
			name: "existing_directory",
			path: filepath.Join(root, "dir1"),
			want: true,
		},
		{
			name: "existing_file",
			path: filepath.Join(root, "file1"),
			want: false,
		},
		{
			name: "non_existing_directory",
			path: filepath.Join(root, "dir2"),
			want: false,
		},
		{
			name: "non_existing_file",
			path: filepath.Join(root, "file2"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := absDirExists(tt.path); got != tt.want {
				t.Errorf("absDirExists(%s) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
