# aerospace utilities
A set of utilities to complement [aerospace](https://github.com/nikitabobko/AeroSpace).

`aww` is the cli tool to run in a terminal.

`awwapp` is a windowed MacOS version to run in a GUI.

At the moment both programs just show the workspaces and the windows in those workspaces.

## Installation one time
```bash
go mod init aws
go mod tidy
```

## Build

    ./build.sh

Tag the version

    git tag --list
    git tag -a 0.1.0 -m "First release"
    git push origin 0.1.0

