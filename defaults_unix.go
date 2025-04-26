//go:build unix && !(darwin || ios)

package xdg

import (
	"path/filepath"
)

func defaultCacheHome() string {
	return filepath.Join(HomeDir(), ".cache")
}

func defaultConfigHome() string {
	return filepath.Join(HomeDir(), ".config")
}

func defaultConfigDirs() string {
	return "/etc/xdg"
}

func defaultDataHome() string {
	return filepath.Join(HomeDir(), ".local/share")
}

func defaultDataDirs() string {
	return "/usr/local/share:/usr/share"
}

func defaultStateHome() string {
	return filepath.Join(HomeDir(), ".local/state")
}
