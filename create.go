package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// NewDirectoryPermissions represents secure directory permissions: rwx --- ---.
	NewDirectoryPermissions = 0o700

	// NewFilePermissions represents secure file permissions: rw- --- ---.
	NewFilePermissions = 0o600

	pathSep = string(filepath.Separator)
)

// CreateDir returns the path to the given app's directory, under the
// given XDG base directory, and creates it if it doesn't already exist.
// Attention: this function normalizes the app name and ensures it does
// not contain path elements, but the caller is responsible for input vetting.
func CreateDir(dirType func() (string, error), appName string) (string, error) {
	appName = filepath.Clean(appName)
	if appName == "." {
		return "", fmt.Errorf("app name is empty")
	}
	if strings.Contains(appName, pathSep) {
		return "", fmt.Errorf("app name must not contain separator")
	}

	path, err := dirType()
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, appName)
	if err := os.MkdirAll(path, NewDirectoryPermissions); err != nil {
		return "", err
	}

	return path, nil
}

// CreateFile returns the path to the given app's file, under the
// given XDG base directory, and creates it if it doesn't already exist.
// Attention: this function normalizes the app and file names and ensures they
// do not contain path elements, but the caller is responsible for input vetting.
func CreateFile(dirType func() (string, error), appName, fileName string) (string, error) {
	fileName = filepath.Clean(fileName)
	if fileName == "." {
		return "", fmt.Errorf("file name is empty")
	}
	if strings.Contains(fileName, pathSep) {
		return "", fmt.Errorf("file name must not contain separator")
	}

	path, err := CreateDir(dirType, appName)
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, fileName)
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, NewFilePermissions) //gosec:disable G304
	if err != nil {
		return "", err
	}

	_ = f.Close()
	return path, nil
}
