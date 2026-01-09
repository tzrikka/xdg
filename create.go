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

// CreateDir returns the path to the given app's directory under the given
// XDG base directory. It creates any directories that don't exist yet.
//
// Note: this function normalizes the app name and ensures it does not
// contain path elements, but the caller is responsible for input vetting.
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

// CreateSubdir returns the path to the given subdirectory under the given app's directory,
// under the given XDG base directory. It creates any directories that don't exist yet.
//
// Note 1: this function normalizes the app name and ensures it doesn't
// contain path elements, but the caller is responsible for input vetting.
//
// Note 2: the subpath parameter may contain 0 or more path elements.
// This function ensures that it does not escape the app's directory.
func CreateSubdir(dirType func() (string, error), appName, subpath string) (string, error) {
	path, err := CreateDir(dirType, appName)
	if err != nil {
		return "", err
	}

	subpath = filepath.Clean(subpath)
	if subpath == "." {
		return path, nil
	}

	root, err := os.OpenRoot(path)
	if err != nil {
		return "", err
	}
	defer root.Close()

	if err := root.MkdirAll(subpath, NewDirectoryPermissions); err != nil {
		return "", err
	}

	return filepath.Join(path, subpath), nil
}

// CreateFile returns the path to the given app's file under the given XDG base directory.
// It creates the file and any parent directories if they don't exist yet.
//
// Note: this function normalizes the app and file names and ensures they
// don't contain path elements, but the caller is responsible for input vetting.
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

	return path, f.Close()
}

// CreateFilePath returns the path to the given app's file (which may be in a subdirectory)
// under the given XDG base directory. It creates the file and any parent directories
// if they don't exist yet.
//
// Note 1: this function normalizes the app name and ensures it doesn't
// contain path elements, but the caller is responsible for input vetting.
//
// Note 2: the filePath parameter must contain at least a filename, and may contain a prefix of
// 0 or more path elements. This function ensures that it does not escape the app's directory.
func CreateFilePath(dirType func() (string, error), appName, filePath string) (string, error) {
	if _, file := filepath.Split(filePath); file == "" {
		return "", fmt.Errorf("file path must end with a file name")
	}

	filePath = filepath.Clean(filePath)
	if filePath == "." {
		return "", fmt.Errorf("file path is empty")
	}

	subpath, file := filepath.Split(filePath)
	path, err := CreateSubdir(dirType, appName, subpath)
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, file)
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, NewFilePermissions) //gosec:disable G304
	if err != nil {
		return "", err
	}

	return path, f.Close()
}
