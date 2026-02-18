package xdg

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestCacheHome(t *testing.T) {
	cachedHomeDir = ""
	t.Cleanup(func() { cachedHomeDir = "" })
	t.Setenv("XDG_CACHE_HOME", "")

	got, err := CacheHome()
	if err != nil {
		t.Errorf("CacheHome() error = %v", err)
	}
	if got != defaultCacheHome() {
		t.Errorf("CacheHome() = %q, want %q", got, defaultCacheHome())
	}
}

func TestConfigHome(t *testing.T) {
	cachedHomeDir = ""
	t.Cleanup(func() { cachedHomeDir = "" })

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
	cachedHomeDir = ""
	t.Cleanup(func() { cachedHomeDir = "" })

	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "dir"), NewDirectoryPermissions); err != nil {
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
			if !slices.Equal(got, tt.want) {
				t.Errorf("ConfigDirs() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDataHome(t *testing.T) {
	cachedHomeDir = ""
	t.Cleanup(func() { cachedHomeDir = "" })
	t.Setenv("XDG_DATA_HOME", "")

	got, err := DataHome()
	if err != nil {
		t.Errorf("DataHome() error = %v", err)
	}
	if got != defaultDataHome() {
		t.Errorf("DataHome() = %q, want %q", got, defaultDataHome())
	}
}

func TestDataDirs(t *testing.T) {
	cachedHomeDir = ""
	t.Cleanup(func() { cachedHomeDir = "" })

	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "dir"), NewDirectoryPermissions); err != nil {
		t.Fatal(err)
	}

	dirs := strings.Join([]string{"/abs/path1", "/abs/path2", filepath.Join(dir, "dir")}, listSeparator)
	t.Setenv("XDG_DATA_DIRS", dirs)
	want := []string{filepath.Join(dir, "dir")}

	got, err := DataDirs()
	if err != nil {
		t.Errorf("DataDirs() error = %v", err)
	}
	if !slices.Equal(got, want) {
		t.Errorf("DataDirs() = %q, want %q", got, want)
	}
}

func TestStateHome(t *testing.T) {
	cachedHomeDir = ""
	t.Cleanup(func() { cachedHomeDir = "" })
	t.Setenv("XDG_STATE_HOME", "")

	got, err := StateHome()
	if err != nil {
		t.Errorf("StateHome() error = %v", err)
	}
	if got != defaultStateHome() {
		t.Errorf("StateHome() = %q, want %q", got, defaultStateHome())
	}
}

func TestAbsDirExists(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "dir1"), NewDirectoryPermissions); err != nil {
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
