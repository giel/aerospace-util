#!/usr/bin/env bash

# todo: set to versioned build tools
go get fyne.io/fyne/v2
go install fyne.io/tools/cmd/fyne@latest

function ShowInfo {
	echo "--------------"
	echo -e "--" $1
	echo "--------------"
}

ShowInfo "Go get fyne library"
go get fyne.io/fyne/v2

ShowInfo "Check directory"
if [[ "$(basename "$PWD")" == "awwapp" ]]; then
	echo "Running in the 'awwapp' directory."
else
	echo "Not running in the 'awwapp' directory, but in:"
	echo "$PWD"
	exit
fi

ShowInfo "Go fmt"
go fmt

ShowInfo "Remove old build"
rm -f ./awwapp
rm -rf ./awwapp.app

if [[ -z $GITHUB_ACTIONS ]]; then
	version=$(git describe)
	short_version="${version%%-*}"
	# version=1.2.3
fi

ShowInfo "Version"
echo $version
echo $short_version

ShowInfo "Build"
go build -ldflags "-X 'golang/asu.Version=$version'"
# go build

ShowInfo "Wrap in MacOS package"
fyne package --executable awwapp --os darwin --icon awwapp-image.png --app-version $short_version
