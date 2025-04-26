package xdg

import (
	"path/filepath"
	"strings"
)

func defaultCacheHome() string {
	return filepath.Join(HomeDir(), "Library/Caches")
}

func defaultConfigHome() string {
	return filepath.Join(HomeDir(), "Library/Application Support")
}

func defaultConfigDirs() string {
	return strings.Join([]string{
		filepath.Join(HomeDir(), "Library/Preferences"),
		"/Library/Application Support",
		"/Library/Preferences",
		filepath.Join(HomeDir(), ".config"),
		"/etc/xdg",
	}, listSeparator)
}

func defaultDataHome() string {
	return filepath.Join(HomeDir(), "Library/Application Support")
}

func defaultDataDirs() string {
	return strings.Join([]string{
		"/Library/Application Support",
		filepath.Join(HomeDir(), ".local/share"),
		"/usr/local/share",
		"/usr/share",
	}, listSeparator)
}

func defaultStateHome() string {
	return filepath.Join(HomeDir(), "Library/Application Support")
}
