#!/usr/bin/env bash

go get fyne.io/fyne/v2

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

if [[ -z $GITHUB_ACTIONS ]]; then
	version=$(git describe)
	version=1.2.3
fi

ShowInfo "Version $version"

ShowInfo "Build"
# CGO_ENABLED=0 go build -ldflags "-X 'awwapp/awwapp.Version=$version'"
go build

ShowInfo "Build version"
# ./aww --version
# ./awwapp

# ShowInfo "Test"
# go test ./mgc_test
# # go test -v ./mgc_test
