package xdg

import (
	"testing"
)

func TestCreateDir(t *testing.T) {
	tests := []struct {
		name    string
		appName string
		want    string
		wantErr bool
	}{
		{
			name:    "empty_app_name",
			wantErr: true,
		},
		{
			name:    "empty_app_name",
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
				t.Errorf("TouchDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TouchDir() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCreateFile(t *testing.T) {
	tests := []struct {
		name     string
		appName  string
		fileName string
		want     string
		wantErr  bool
	}{
		{
			name:     "empty_app_name",
			fileName: "my_file",
			wantErr:  true,
		},
		{
			name:     "empty_app_name",
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
		{
			name:     "empty_file_name",
			appName:  "my_app",
			fileName: "",
			wantErr:  true,
		},
		{
			name:     "empty_file_name",
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
			name:     "insecure_file_name",
			appName:  "my_app",
			fileName: "../../other_app/other_file",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFile(tempDir(t), tt.appName, tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("TouchFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TouchFile() = %q, want %q", got, tt.want)
			}
		})
	}
}

func tempDir(t *testing.T) func() (string, error) {
	return func() (string, error) {
		return t.TempDir(), nil
	}
}
