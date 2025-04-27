package xdg

// MustCacheHome is like [CacheHome]. It discards the error
// and returns only the path, but panics if there is an error.
func MustCacheHome() string {
	return mustHome(CacheHome())
}

// MustConfigHome is like [ConfigHome]. It discards the error
// and returns only the path, but panics if there is an error.
func MustConfigHome() string {
	return mustHome(ConfigHome())
}

// MustConfigDirs is like [ConfigDirs]. It discards the error
// and returns only the path, but panics if there is an error.
func MustConfigDirs() []string {
	return mustDirs(ConfigDirs())
}

// MustDataHome is like [DataHome]. It discards the error
// and returns only the path, but panics if there is an error.
func MustDataHome() string {
	return mustHome(DataHome())
}

// MustDataDirs is like [DataDirs]. It discards the error
// and returns only the path, but panics if there is an error.
func MustDataDirs() []string {
	return mustDirs(DataDirs())
}

// MustStateHome is like [StateHome]. It discards the error
// and returns only the path, but panics if there is an error.
func MustStateHome() string {
	return mustHome(StateHome())
}

func mustHome(path string, err error) string {
	if err != nil {
		panic(err)
	}
	return path
}

func mustDirs(paths []string, err error) []string {
	if err != nil {
		panic(err)
	}
	return paths
}
