package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindCacheFile looks for a filename in an app's [CacheHome] directory.
// If the file is found, this function returns its full path. If not,
// it returns an empty string but no error. An error is returned only
// if the input parameters are invalid, or in case of a runtime error.
func FindCacheFile(appName, fileName string) (string, error) {
	return findFile(CacheHome, nil, appName, fileName)
}

// FindConfigFile looks for a filename in an app's [ConfigHome] and [ConfigDirs]
// directories. If the file is found, this function returns its full path.
// If not, it returns an empty string but no error. An error is returned
// only if the input parameters are invalid, or in case of a runtime error.
func FindConfigFile(appName, fileName string) (string, error) {
	return findFile(ConfigHome, ConfigDirs, appName, fileName)
}

// FindDataFile looks for a filename in an app's [DataHome] and [DataDirs]
// directories. If the file is found, this function returns its full path.
// If not, it returns an empty string but no error. An error is returned
// only if the input parameters are invalid, or in case of a runtime error.
func FindDataFile(appName, fileName string) (string, error) {
	return findFile(DataHome, DataDirs, appName, fileName)
}

// FindStateFile looks for a filename in an app's [StateHome] directory.
// If the file is found, this function returns its full path. If not,
// it returns an empty string but no error. An error is returned only
// if the input parameters are invalid, or in case of a runtime error.
func FindStateFile(appName, fileName string) (string, error) {
	return findFile(StateHome, nil, appName, fileName)
}

func findFile(home func() (string, error), dirs func() ([]string, error), appName, fileName string) (string, error) {
	appName = filepath.Clean(appName)
	if appName == "." {
		return "", fmt.Errorf("app name is empty")
	}
	if strings.Contains(appName, pathSep) {
		return "", fmt.Errorf("app name must not contain separator")
	}

	fileName = filepath.Clean(fileName)
	if fileName == "." {
		return "", fmt.Errorf("file name is empty")
	}
	if strings.Contains(fileName, pathSep) {
		return "", fmt.Errorf("file name must not contain separator")
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
	paths = append(paths, morePaths...)

	for _, path := range paths {
		if fullPath := filepath.Join(path, appName, fileName); fileExists(fullPath) {
			return fullPath, nil
		}
	}

	return "", nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
