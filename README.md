# aerospace utilities
A set of utilities to complement [aerospace](https://github.com/nikitabobko/AeroSpace).

`aww` is the cli tool to run in a terminal.

`awwapp` was a windowed MacOS version to run in a GUI. It is obsolete, since version 0.19.0 shows it in the main app toolbar. That is why it is removed from the repository.

Both programs show the workspaces and the windows in those workspaces. Also the second column shows the monitor on wich the program is shown.

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

## Sample output

Below you find an example of the output of the CLI version of the program.

```bash
> aww
2  (1)  ChatGPT : ChatGPT
3  (1)  1Password : Lock Screen — 1Password
3  (1)  Finder : _Apps
3  (1)  Finder : bin
Q  (2)  Brave Browser : Startpage - Private Search Engine
S  (1)  Signal : Signal
W  (2)  Ghostty : aww
W  (2)  Ghostty : nvim
(1) Built-in Retina Display
(2) BenQ PD2730S
> 
> aww --version
aww version 0.3.1
>
```

