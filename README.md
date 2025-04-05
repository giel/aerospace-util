# aerospace utilities
A set of utilities to complement aerospace.

`aww` is the cli tool to run in a terminal.

`awwapp` is a windowed MacOS version to run in a GUI.


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

