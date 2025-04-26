package xdg

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows"
)

func defaultCacheHome() string {
	return filepath.Join(localAppData(), "Cache")
}

func defaultConfigHome() string {
	return localAppData()
}

func defaultConfigDirs() string {
	return strings.Join([]string{
		programData(),
		roamingAppData(),
	}, listSeparator)
}

func defaultDataHome() string {
	return localAppData()
}

func defaultDataDirs() string {
	return strings.Join([]string{
		roamingAppData(),
		programData(),
	}, listSeparator)
}

func defaultStateHome() string {
	return localAppData()
}

// folderPath returns the first non-empty path of a specific Known Folder.
// It returns an empty string if all attempts have failed, in which case
// the caller should construct a default speculative path.
func folderPath(id *windows.KNOWNFOLDERID, envVarName string) string {
	// First attempt (most authoritative).
	for _, flag := range []uint32{windows.KF_FLAG_DEFAULT, windows.KF_FLAG_DEFAULT_PATH} {
		if path, err := windows.KnownFolderPath(id, flag); err == nil && path != "" {
			return path
		}
	}

	// Second attempt (less reliable because values can be modified manually).
	if path := os.ExpandEnv(os.Getenv(envVarName)); path != "" {
		return path
	}

	return ""
}

func localAppData() string {
	if path := folderPath(windows.FOLDERID_LocalAppData, "LOCALAPPDATA"); path != "" {
		return path
	}
	return filepath.Join(HomeDir(), "AppData", "Local")
}

func roamingAppData() string {
	if path := folderPath(windows.FOLDERID_RoamingAppData, "APPDATA"); path != "" {
		return path
	}
	return filepath.Join(HomeDir(), "AppData", "Roaming")
}

func programData() string {
	if path := folderPath(windows.FOLDERID_ProgramData, "ALLUSERSPROFILE"); path != "" {
		return path
	}

	if path := os.ExpandEnv(os.Getenv("ProgramData")); path != "" {
		return path
	}

	return filepath.Join(systemDrive(), "ProgramData")
}

func systemDrive() string {
	if path := os.Getenv("SystemDrive"); path != "" {
		return path
	}
	return "C:"
}
