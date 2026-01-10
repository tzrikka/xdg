package xdg

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestFindFile(t *testing.T) {
	homeDir := t.TempDir()
	homeFile, err := CreateFile(dirForTest(homeDir, nil), "appName", "homeFile")
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	dirsDir := t.TempDir()
	dirsFile, err := CreateFile(dirForTest(dirsDir, nil), "appName", "dirsFile")
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	tests := []struct {
		name     string
		home     func() (string, error)
		dirs     func() ([]string, error)
		fileName string
		want     string
		wantErr  bool
	}{
		{
			name:     "home_file_found_without_dirs",
			home:     dirForTest(homeDir, nil),
			fileName: "homeFile",
			want:     homeFile,
		},
		{
			name:     "home_file_found_with_dirs",
			home:     dirForTest(homeDir, nil),
			dirs:     dirsForTest([]string{dirsDir}, nil),
			fileName: "homeFile",
			want:     homeFile,
		},
		{
			name:     "dirs_file_found",
			home:     dirForTest(homeDir, nil),
			dirs:     dirsForTest([]string{dirsDir}, nil),
			fileName: "dirsFile",
			want:     dirsFile,
		},
		{
			name:     "file_not_found",
			home:     dirForTest(homeDir, nil),
			dirs:     dirsForTest([]string{dirsDir}, nil),
			fileName: "nonexistentFileName",
		},
		{
			name:     "missing_file_name",
			home:     dirForTest(homeDir, nil),
			dirs:     dirsForTest([]string{dirsDir}, nil),
			fileName: "",
			wantErr:  true,
		},
		{
			name:     "home_dir_error",
			home:     dirForTest(homeDir, errors.New("home dir error")),
			dirs:     dirsForTest([]string{dirsDir}, nil),
			fileName: "homeFile",
			wantErr:  true,
		},
		{
			name:     "dirs_error",
			home:     dirForTest(homeDir, nil),
			dirs:     dirsForTest([]string{dirsDir}, errors.New("dirs error")),
			fileName: "dirsFile",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findFile(tt.home, tt.dirs, "appName", tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("findFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func dirForTest(dir string, err error) func() (string, error) {
	return func() (string, error) {
		if err != nil {
			return "", err
		}
		return dir, nil
	}
}

func dirsForTest(dirs []string, err error) func() ([]string, error) {
	return func() ([]string, error) {
		if err != nil {
			return nil, err
		}
		return dirs, nil
	}
}

func TestFullPath(t *testing.T) {
	appName := "my_app"
	baseFile := "test_file"

	tests := []struct {
		name     string
		subdir   string
		dontFind bool
	}{
		{
			name: "shallow_file_exists",
		},
		{
			name:     "shallow_file_does_not_exist",
			dontFind: true,
		},
		{
			name:   "deep_file_exists",
			subdir: "subdir1/subdir2",
		},
		{
			name:     "deep_file_does_not_exist",
			subdir:   "subdir",
			dontFind: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()

			path := filepath.Join(tempDir, appName, tt.subdir)
			if err := os.MkdirAll(path, NewDirectoryPermissions); err != nil {
				t.Fatalf("failed to create test subdir: %v", err)
			}

			path = filepath.Join(path, baseFile)
			if err := os.WriteFile(path, []byte("test"), NewFilePermissions); err != nil {
				t.Fatalf("failed to write test file: %v", err)
			}

			filename := baseFile
			if tt.dontFind {
				filename += "_not_found"
			}
			got, err := fullPath(tempDir, appName, filepath.Join(tt.subdir, filename))
			if err != nil {
				t.Errorf("fullPath() error = %v", err)
				return
			}

			if tt.dontFind {
				path = ""
			}
			if got != path {
				t.Errorf("fullPath() = %q, want %q", got, path)
			}
		})
	}
}
