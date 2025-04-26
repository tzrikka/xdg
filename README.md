# Cross-Platform XDG Base Directories

[![Go Reference](https://pkg.go.dev/badge/github.com/tzrikka/xdg.svg)](https://pkg.go.dev/github.com/tzrikka/xdg)
[![Go Report Card](https://goreportcard.com/badge/github.com/tzrikka/xdg)](https://goreportcard.com/report/github.com/tzrikka/xdg)

This package implements the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/latest/), with adjustments for [macOS](https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html#//apple_ref/doc/uid/TP40010672-CH10-SW1) and [Windows](http://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid).

Supported operating systems: all Unix flavors, macOS, Windows.

It was created with these design goals/principles:

1. Code and usage simplicity
2. Seamless handling of `XDG_*` environment variable changes during runtime
3. Minimal dependencies on other third-party Go packages
