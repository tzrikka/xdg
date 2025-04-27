# Cross-Platform XDG Base Directories

[![Go Reference](https://pkg.go.dev/badge/github.com/tzrikka/xdg.svg)](https://pkg.go.dev/github.com/tzrikka/xdg)
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

### All Unix Flavors

`XDG_CACHE_HOME`

| Operating System | Path                   |
| :--------------- | :--------------------- |
| Unix             | `$HOME/.cache`         |
| macOS            | `$HOME/Library/Caches` |

`XDG_CONFIG_HOME`

| Operating System | Path                                |
| :--------------- | :---------------------------------- |
| Unix             | `$HOME/.config`                     |
| macOS            | `$HOME/Library/Application Support` |

`XDG_CONFIG_DIRS`

| Operating System | Path(s)                        |
| :--------------- | :----------------------------- |
| Unix             | `/etc/xdg`                     |
| macOS            | `/Library/Application Support` |
|                  | `$HOME/.config`                |
|                  | `/etc/xdg`                     |

`XDG_DATA_HOME`

| Operating System | Path                                |
| :--------------- | :---------------------------------- |
| Unix             | `$HOME/.local/share`                |
| macOS            | `$HOME/Library/Application Support` |

`XDG_DATA_DIRS`

| Operating System | Path(s)                        |
| :--------------- | :----------------------------- |
| Unix             | `/usr/local/share`             |
|                  | `/usr/share`                   |
| macOS            | `/Library/Application Support` |
|                  | `$HOME/.local/share`           |
|                  | `/usr/local/share`             |
|                  | `/usr/share`                   |

`XDG_STATE_HOME`

| Operating System | Path                                |
| :--------------- | :---------------------------------- |
| Unix             | `$HOME/.local/state`                |
| macOS            | `$HOME/Library/Application Support` |

### Microsoft Windows

[Known folder IDs](https://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid)

[Recognized environment variables](https://learn.microsoft.com/en-us/windows/deployment/usmt/usmt-recognized-environment-variables)

`XDG_CACHE_HOME`

| Type         | Value                   |
| :----------- | :---------------------- |
| Known Folder | `FOLDERID_LocalAppData` |
| Env Var      | `%LOCALAPPDATA%`        |

`XDG_CONFIG_HOME`

| Type         | Value                     |
| :----------- | :------------------------ |
| Known Folder | `FOLDERID_RoamingAppData` |
| Env Var      | `%APPDATA%`               |

`XDG_CONFIG_DIRS`

| Type         | Value                                  |
| :----------- | :------------------------------------- |
| Known Folder | `FOLDERID_ProgramData`                 |
| Env Var      | `%ALLUSERSPROFILE%` or `%ProgramData%` |

`XDG_DATA_HOME`

| Type         | Value                   |
| :----------- | :---------------------- |
| Known Folder | `FOLDERID_LocalAppData` |
| Env Var      | `%LOCALAPPDATA%`        |

`XDG_DATA_DIRS`

| Type         | Value                                  |
| :----------- | :------------------------------------- |
| Known Folder | `FOLDERID_ProgramData`                 |
| Env Var      | `%ALLUSERSPROFILE%` or `%ProgramData%` |

`XDG_STATE_HOME`

| Type         | Value                   |
| :----------- | :---------------------- |
| Known Folder | `FOLDERID_LocalAppData` |
| Env Var      | `%LOCALAPPDATA%`        |
