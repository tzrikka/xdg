package xdg

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// FindCacheFile looks for a file in an app's [CacheHome] directory.
// If the file is found, this function returns its full path. If not,
// it returns an empty string but no error. An error is returned only
// if the input parameters are invalid, or in case of a runtime error.
func FindCacheFile(appName, filePath string) (string, error) {
	return findFile(CacheHome, nil, appName, filePath)
}

// FindConfigFile looks for a file in an app's [ConfigHome] and [ConfigDirs]
// directories. If the file is found, this function returns its full path.
// If not, it returns an empty string but no error. An error is returned
// only if the input parameters are invalid, or in case of a runtime error.
func FindConfigFile(appName, filePath string) (string, error) {
	return findFile(ConfigHome, ConfigDirs, appName, filePath)
}

// FindDataFile looks for a file in an app's [DataHome] and [DataDirs]
// directories. If the file is found, this function returns its full path.
// If not, it returns an empty string but no error. An error is returned
// only if the input parameters are invalid, or in case of a runtime error.
func FindDataFile(appName, filePath string) (string, error) {
	return findFile(DataHome, DataDirs, appName, filePath)
}

// FindStateFile looks for a file in an app's [StateHome] directory.
// If the file is found, this function returns its full path. If not,
// it returns an empty string but no error. An error is returned only
// if the input parameters are invalid, or in case of a runtime error.
func FindStateFile(appName, filePath string) (string, error) {
	return findFile(StateHome, nil, appName, filePath)
}

func findFile(home func() (string, error), dirs func() ([]string, error), appName, filePath string) (string, error) {
	appName = filepath.Clean(appName)
	if appName == "." {
		return "", errors.New("app name is empty")
	}
	if strings.Contains(appName, pathSep) {
		return "", errors.New("app name must not contain separator")
	}

	filePath = filepath.Clean(filePath)
	if filePath == "." {
		return "", errors.New("file path is empty")
	}

	firstPath, err := home()
	if err != nil {
		return "", err
	}

	var morePaths []string
	if dirs != nil {
		morePaths, err = dirs()
		if err != nil {
			return "", err
		}
	}

	paths := make([]string, 1, 1+len(morePaths))
	paths[0] = firstPath

	for _, path := range append(paths, morePaths...) {
		fp, err := fullPath(path, appName, filePath)
		if err != nil {
			return "", err
		}
		if fp != "" {
			return fp, nil
		}
	}

	return "", nil
}

func fullPath(path, appName, filePath string) (string, error) {
	appPath := filepath.Join(path, appName)
	info, err := os.Stat(appPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}
	if !info.IsDir() {
		return "", nil // Found app file instead of app directory.
	}

	root, err := os.OpenRoot(appPath)
	if err != nil || root == nil {
		return "", err // Error may or may not be nil.
	}
	defer root.Close()

	info, err = root.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}
	if info.IsDir() {
		return "", nil // Found subdirectory instead of file.
	}

	return filepath.Join(appPath, filePath), nil
}
