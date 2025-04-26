package xdg

// MustHome is an optional wrapper for [CacheHome], [ConfigHome],
// [DataHome], and [StateHome] - to discard errors in successful
// calls, but treat them as panics if they do occur.
func MustHome(path string, err error) string {
	if err != nil {
		panic(err)
	}
	return path
}

// MustDirs is an optional wrapper for [ConfigDirs] and [DataDirs], to discard
// errors in successful calls, but treat them as panics if they do occur.
func MustDirs(paths []string, err error) []string {
	if err != nil {
		panic(err)
	}
	return paths
}
