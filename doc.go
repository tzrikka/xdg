// Package xdg implements the [XDG Base Directory Specification],
// with adjustments for [macOS] and [Windows].
//
// Supported operating systems: all Unix flavors, macOS, and Windows.
//
// It was created with these design goals/principles:
//   - Code and usage simplicity
//   - Seamless handling of "XDG_*" environment variable changes during runtime
//   - Minimal dependencies on other third-party Go packages
//
// [XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest/
// [macOS]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html#//apple_ref/doc/uid/TP40010672-CH10-SW1
// [Windows]: http://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid
package xdg
