package xdg

import (
	"os"
)

var cachedHomeDir string

// HomeDir returns the current user's home directory.
// It is assumed that this value does not change during runtime.
func HomeDir() string {
	if cachedHomeDir != "" {
		return cachedHomeDir
	}

	path, err := os.UserHomeDir()
	if err == nil {
		cachedHomeDir = path
	}

	return path
}
