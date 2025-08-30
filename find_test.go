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
			name:     "bad_file_name",
			home:     dirForTest(homeDir, nil),
			dirs:     dirsForTest([]string{dirsDir}, nil),
			fileName: "file/name",
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

func TestFileExists(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test_file")

	if err := os.WriteFile(path, []byte("test"), 0o600); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "file_exists",
			path: path,
			want: true,
		},
		{
			name: "file_does_not_exist",
			path: path + "bad",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExists(tt.path); got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
