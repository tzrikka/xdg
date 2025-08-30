package xdg

// MustCacheHome is like [CacheHome]. It discards the error
// and returns only the path, but panics if there is an error.
func MustCacheHome() string {
	return must(CacheHome())
}

// MustConfigHome is like [ConfigHome]. It discards the error
// and returns only the path, but panics if there is an error.
func MustConfigHome() string {
	return must(ConfigHome())
}

// MustConfigDirs is like [ConfigDirs]. It discards the error
// and returns only the path, but panics if there is an error.
func MustConfigDirs() []string {
	return must(ConfigDirs())
}

// MustDataHome is like [DataHome]. It discards the error
// and returns only the path, but panics if there is an error.
func MustDataHome() string {
	return must(DataHome())
}

// MustDataDirs is like [DataDirs]. It discards the error
// and returns only the path, but panics if there is an error.
func MustDataDirs() []string {
	return must(DataDirs())
}

// MustStateHome is like [StateHome]. It discards the error
// and returns only the path, but panics if there is an error.
func MustStateHome() string {
	return must(StateHome())
}

func must[T any](path T, err error) T {
	if err != nil {
		panic(err)
	}
	return path
}
