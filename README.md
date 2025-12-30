# Cross-Platform XDG Base Directories

[![Go Reference](https://pkg.go.dev/badge/github.com/tzrikka/xdg.svg)](https://pkg.go.dev/github.com/tzrikka/xdg)
[![Code Wiki](https://img.shields.io/badge/Code_Wiki-gold?logo=googlegemini)](https://codewiki.google/github.com/tzrikka/xdg)
[![Go Report Card](https://goreportcard.com/badge/github.com/tzrikka/xdg)](https://goreportcard.com/report/github.com/tzrikka/xdg)

This package implements the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/latest/), with adjustments for [macOS](https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html#//apple_ref/doc/uid/TP40010672-CH10-SW1) and [Windows](http://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid).

Supported operating systems: all Unix flavors, macOS, and Windows.

It was created with these design goals/principles:

- Code and usage simplicity
- Seamless handling of `XDG_*` environment variable changes during runtime
- Minimal dependencies on other third-party Go packages

## Installation

```shell
go get github.com/tzrikka/xdg
```

## Default Paths

### Unix & macOS

| Env Var           | Unix                 | macOS                               |
| :---------------- | :------------------- | :---------------------------------- |
| `XDG_CACHE_HOME`  | `$HOME/.cache`       | `$HOME/Library/Caches`              |
| `XDG_CONFIG_HOME` | `$HOME/.config`      | `$HOME/.config`                     |
| `XDG_CONFIG_DIRS` | `/etc/xdg`           | `$HOME/Library/Application Support` |
|                   |                      | `/Library/Application Support`      |
|                   |                      | `/etc/xdg`                          |
| `XDG_DATA_HOME`   | `$HOME/.local/share` | `$HOME/Library/Application Support` |
| `XDG_DATA_DIRS`   | `/usr/local/share`   | `/Library/Application Support`      |
|                   | `/usr/share`         | `$HOME/.local/share`                |
|                   |                      | `/usr/local/share`                  |
|                   |                      | `/usr/share`                        |
| `XDG_STATE_HOME`  | `$HOME/.local/state` | `$HOME/Library/Application Support` |

### Microsoft Windows

| XDG Env Var       | Known Folder              | Windows Env Var     |
| :---------------- | :------------------------ | :------------------ |
| `XDG_CACHE_HOME`  | `FOLDERID_LocalAppData`   | `%LOCALAPPDATA%`    |
| `XDG_CONFIG_HOME` | `FOLDERID_RoamingAppData` | `%APPDATA%`         |
| `XDG_CONFIG_DIRS` | `FOLDERID_ProgramData`    | `%ALLUSERSPROFILE%` |
|                   |                           | or `%ProgramData%`  |
| `XDG_DATA_HOME`   | `FOLDERID_LocalAppData`   | `%LOCALAPPDATA%`    |
| `XDG_DATA_DIRS`   | `FOLDERID_ProgramData`    | `%ALLUSERSPROFILE%` |
|                   |                           | or `%ProgramData%`  |
| `XDG_STATE_HOME`  | `FOLDERID_LocalAppData`   | `%LOCALAPPDATA%`    |

- [Known folder IDs](https://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid)
- [Recognized environment variables](https://learn.microsoft.com/en-us/windows/deployment/usmt/usmt-recognized-environment-variables)
