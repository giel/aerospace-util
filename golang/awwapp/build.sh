#!/usr/bin/env bash

go get fyne.io/fyne/v2
go install fyne.io/fyne/v2/cmd/fyne@latest

function ShowInfo {
	echo "--------------"
	echo "--" $1
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
	# version=1.2.3
fi

ShowInfo "Version $version"

ShowInfo "Build"
go build -ldflags "-X 'golang/asu.Version=$version'"
# go build

ShowInfo "Wrap in MacOS package"
fyne package -executable awwapp -os darwin -icon awwapp-image.png -appVersion=$version

# ShowInfo "Build version"
# ./aww --version
# ./awwapp

# ShowInfo "Test"
# go test ./mgc_test
# # go test -v ./mgc_test
