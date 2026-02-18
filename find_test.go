package xdg

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

type findFunc func(appName, fileName string) (string, error)

func TestFindFileAppNameErrors(t *testing.T) {
	tests := []struct {
		name    string
		fn      findFunc
		appName string
	}{
		{
			name:    "FindCacheFile",
			fn:      FindCacheFile,
			appName: "",
		},
		{
			name:    "FindConfigFile",
			fn:      FindConfigFile,
			appName: "",
		},
		{
			name:    "FindDataFile",
			fn:      FindDataFile,
			appName: filepath.Join("subdir", "file"),
		},
		{
			name:    "FindStateFile",
			fn:      FindStateFile,
			appName: filepath.Join("subdir", "file"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn(tt.appName, "filePath")
			if err == nil {
				t.Errorf("%s() error = nil, wantErr true", tt.name)
			}
			if got != "" {
				t.Errorf("%s() = %q, want %q", tt.name, got, "")
			}
		})
	}
}

func TestFullPathAppNameErrors(t *testing.T) {
	path := t.TempDir()
	if err := os.WriteFile(filepath.Join(path, "notDir"), []byte{}, NewFilePermissions); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(path, "appName", "subdir"), NewDirectoryPermissions); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		appName  string
		filePath string
		want     string
	}{
		{
			name:     "app_name_not_found",
			appName:  "nonexistentApp",
			filePath: "filePath",
		},
		{
			name:     "app_name_not_a_directory",
			appName:  "notDir",
			filePath: "filePath",
		},
		{
			name:     "file_name_not_a_file",
			appName:  "appName",
			filePath: "subdir",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fullPath(path, tt.appName, tt.filePath)
			if err != nil {
				t.Errorf("fullPath() error = %v", err)
			}
			if got != "" {
				t.Errorf("fullPath() = %q, want %q", got, "")
			}
		})
	}
}

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
			name:    "empty_file_name",
			home:    dirForTest(homeDir, nil),
			dirs:    dirsForTest([]string{dirsDir}, nil),
			wantErr: true,
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
