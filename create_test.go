package xdg

import (
	"testing"
)

func TestCreateDir(t *testing.T) {
	tests := []struct {
		name    string
		appName string
		wantErr bool
	}{
		{
			name:    "happy_path",
			appName: "my_app",
			wantErr: false,
		},
		{
			name:    "empty_app_name_1",
			wantErr: true,
		},
		{
			name:    "empty_app_name_2",
			appName: ".",
			wantErr: true,
		},
		{
			name:    "root_app_name",
			appName: "/",
			wantErr: true,
		},
		{
			name:    "insecure_app_name",
			appName: "../other_app",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateDir(tempDir(t), tt.appName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && got != "" {
				t.Errorf("CreateDir() = %q, want %q", got, "")
			}
		})
	}
}

func TestCreateSubdir(t *testing.T) {
	tests := []struct {
		name    string
		appName string
		subdir  string
		wantErr bool
	}{
		// Happy-path test cases.
		{
			name:    "happy_path_0_subdirs",
			appName: "my_app",
			wantErr: false,
		},
		{
			name:    "happy_path_1_subdir",
			appName: "my_app",
			subdir:  "my_subdir",
			wantErr: false,
		},
		{
			name:    "happy_path_3_subdirs",
			appName: "my_app",
			subdir:  "subdir1/subdir2/subdir3",
			wantErr: false,
		},
		{
			name:    "happy_path_3_subdirs_with_trailing_slash",
			appName: "my_app",
			subdir:  "subdir1/subdir2/subdir3/",
			wantErr: false,
		},
		// CreateDir test cases.
		{
			name:    "empty_app_and_subdir_name",
			wantErr: true,
		},
		{
			name:    "empty_app_name_1",
			subdir:  "my_subdir",
			wantErr: true,
		},
		{
			name:    "empty_app_name_2",
			appName: ".",
			subdir:  "my_subdir",
			wantErr: true,
		},
		{
			name:    "root_app_name",
			appName: "/",
			subdir:  "my_subdir",
			wantErr: true,
		},
		{
			name:    "insecure_app_name",
			appName: "../other_app",
			subdir:  "my_subdir",
			wantErr: true,
		},
		// Subdir-focused test cases.
		{
			name:    "root_subdir",
			appName: "my_app",
			subdir:  "/",
			wantErr: true,
		},
		{
			name:    "insecure_subdir_1",
			appName: "my_app",
			subdir:  "../other_app",
			wantErr: true,
		},
		{
			name:    "insecure_subdir_2",
			appName: "my_app",
			subdir:  "../other_app/other_subdir",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSubdir(tempDir(t), tt.appName, tt.subdir)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSubdir() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && got != "" {
				t.Errorf("CreateSubdir() = %q, want %q", got, "")
			}
		})
	}
}

func TestCreateFile(t *testing.T) {
	tests := []struct {
		name     string
		appName  string
		fileName string
		wantErr  bool
	}{
		// Happy-path test cases.
		{
			name:     "happy_path",
			appName:  "my_app",
			fileName: "my_file",
			wantErr:  false,
		},
		// CreateDir test cases.
		{
			name:    "empty_app_and_subdir_name",
			wantErr: true,
		},
		{
			name:     "empty_app_name_1",
			fileName: "my_file",
			wantErr:  true,
		},
		{
			name:     "empty_app_name_2",
			appName:  ".",
			fileName: "my_file",
			wantErr:  true,
		},
		{
			name:     "root_app_name",
			appName:  "/",
			fileName: "my_file",
			wantErr:  true,
		},
		{
			name:     "insecure_app_name",
			appName:  "../other_app",
			fileName: "my_file",
			wantErr:  true,
		},
		// Filename-focused test cases.
		{
			name:     "empty_file_name_1",
			appName:  "my_app",
			fileName: "",
			wantErr:  true,
		},
		{
			name:     "empty_file_name_2",
			appName:  "my_app",
			fileName: ".",
			wantErr:  true,
		},
		{
			name:     "root_file_name",
			appName:  "my_app",
			fileName: "/",
			wantErr:  true,
		},
		{
			name:     "insecure_file_name_1",
			appName:  "my_app",
			fileName: "../other_app",
			wantErr:  true,
		},
		{
			name:     "insecure_file_name_2",
			appName:  "my_app",
			fileName: "../other_app/other_file",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFile(tempDir(t), tt.appName, tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && got != "" {
				t.Errorf("CreateFile() = %q, want %q", got, "")
			}
		})
	}
}

func TestCreateFilePath(t *testing.T) {
	tests := []struct {
		name     string
		appName  string
		filePath string
		wantErr  bool
	}{
		// Happy-path test cases.
		{
			name:     "happy_path_with_0_subdirs",
			appName:  "my_app",
			filePath: "my_file",
			wantErr:  false,
		},
		{
			name:     "happy_path_with_1_subdir",
			appName:  "my_app",
			filePath: "my_subdir/my_file",
			wantErr:  false,
		},
		{
			name:     "happy_path_with_3_subdirs",
			appName:  "my_app",
			filePath: "subdir1/subdir2/subdir3/my_file",
			wantErr:  false,
		},
		// CreateDir test cases.
		{
			name:    "empty_app_and_subdir_name",
			wantErr: true,
		},
		{
			name:     "empty_app_name_1",
			filePath: "my_file",
			wantErr:  true,
		},
		{
			name:     "empty_app_name_2",
			appName:  ".",
			filePath: "my_file",
			wantErr:  true,
		},
		{
			name:     "root_app_name",
			appName:  "/",
			filePath: "my_file",
			wantErr:  true,
		},
		{
			name:     "insecure_app_name",
			appName:  "../other_app",
			filePath: "my_file",
			wantErr:  true,
		},
		// CreateFile test cases.
		{
			name:     "empty_file_name_1",
			appName:  "my_app",
			filePath: "",
			wantErr:  true,
		},
		{
			name:     "empty_file_name_2",
			appName:  "my_app",
			filePath: ".",
			wantErr:  true,
		},
		{
			name:     "root_file_name",
			appName:  "my_app",
			filePath: "/",
			wantErr:  true,
		},
		{
			name:     "insecure_file_name_1",
			appName:  "my_app",
			filePath: "../other_app",
			wantErr:  true,
		},
		{
			name:     "insecure_file_name_2",
			appName:  "my_app",
			filePath: "../other_app/other_file",
			wantErr:  true,
		},
		// File-path test cases.
		{
			name:     "root_file_path",
			appName:  "my_app",
			filePath: "/my_file",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFilePath(tempDir(t), tt.appName, tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFilePath() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && got != "" {
				t.Errorf("CreateFilePath() = %q, want %q", got, "")
			}
		})
	}
}

func tempDir(t *testing.T) func() (string, error) {
	t.Helper()

	return func() (string, error) {
		return t.TempDir(), nil
	}
}
