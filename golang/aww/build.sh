#!/usr/bin/env bash

function ShowInfo {
	echo "--------------"
	echo "--" $1
	echo "--------------"
}

ShowInfo "Check directory"
if [[ "$(basename "$PWD")" == "aww" ]]; then
	echo "Running in the 'aww' directory."
else
	echo "Not running in the 'aww' directory, but in:"
	echo "$PWD"
	exit
fi

ShowInfo "Go fmt"
gofmt -l -w . ../asu/

if [[ -z $GITHUB_ACTIONS ]]; then
	version=$(git describe)
	# version=1.2.3
fi

ShowInfo "Version $version"

ShowInfo "Remove old binary"
rm -f ./aww

ShowInfo "Build"
CGO_ENABLED=0 go build -ldflags "-X 'golang/asu.Version=$version'"
# CGO_ENABLED=0 go build

ShowInfo "Build version"
./aww --version
# ./aww

# ShowInfo "Test"
# go test ./mgc_test
# # go test -v ./mgc_test
