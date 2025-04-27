package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	listSeparator = string(os.PathListSeparator)
)

// CacheHome returns the absolute path of the base directory in which
// user-specific non-essential (cached) data should be written.
//
// Users should create their own application-specific
// subdirectory within this one and use that.
//
// Subtly different from [os.UserCacheDir] in the following ways:
//   - XDG_CACHE_HOME honored in all operating systems,
//   - Nested environment variables are auto-expanded,
//   - Environment variables may end with a trailing slash.
func CacheHome() (string, error) {
	return dir("XDG_CACHE_HOME", defaultCacheHome)
}

// ConfigHome returns the absolute path of the base directory
// in which user-specific configuration files should be written.
//
// Users should create their own application-specific
// subdirectory within this one and use that.
//
// Subtly different from [os.UserConfigDir] in the following ways:
//   - XDG_CONFIG_HOME honored in all operating systems,
//   - Nested environment variables are auto-expanded,
//   - Environment variables may end with a trailing slash.
func ConfigHome() (string, error) {
	return dir("XDG_CONFIG_HOME", defaultConfigHome)
}

// ConfigDirs returns a set of preference-ordered base directories relative
// to which configuration files should be searched, after [ConfigHome].
func ConfigDirs() ([]string, error) {
	return dirs("XDG_CONFIG_DIRS", defaultConfigDirs)
}

// DataHome returns the absolute path of the base directory
// in which user-specific data files should be written.
//
// Users should create their own application-specific
// subdirectory within this one and use that.
func DataHome() (string, error) {
	return dir("XDG_DATA_HOME", defaultDataHome)
}

// DataDirs returns a set of preference-ordered base directories
// relative to which data files should be searched, after [DataHome].
func DataDirs() ([]string, error) {
	return dirs("XDG_DATA_DIRS", defaultDataDirs)
}

// StateHome returns the absolute path of the base directory
// in which user-specific state data should be written.
//
// Users should create their own application-specific
// subdirectory within this one and use that.
//
// State data should persist between (application) restarts,
// but that is not important or portable enough to the user
// that it should be stored in [DataHome]. It may contain:
//   - Actions history (logs, history, recently used files, ...),
//   - Current state of the application that can be reused on a restart
//     (view, layout, open files, undo history, ...).
func StateHome() (string, error) {
	return dir("XDG_STATE_HOME", defaultStateHome)
}

func dir(envVarName string, defaultFunc func() string) (string, error) {
	path := os.ExpandEnv(os.Getenv(envVarName))
	if path == "" {
		path = defaultFunc()
	}

	if filepath.IsAbs(path) {
		return path, nil
	}

	return "", fmt.Errorf("path in %s is relative: %q", envVarName, path)
}

func dirs(envVarName string, defaultFunc func() string) ([]string, error) {
	path := os.ExpandEnv(os.Getenv(envVarName))
	if path == "" {
		path = defaultFunc()
	}

	var paths []string
	for p := range strings.SplitSeq(path, listSeparator) {
		if absDirExists(p) {
			paths = append(paths, p)
		}
	}

	return paths, nil
}

// absDirExists checks whether the given path is an existing absolute directory.
func absDirExists(path string) bool {
	if !filepath.IsAbs(path) {
		return false
	}

	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}
