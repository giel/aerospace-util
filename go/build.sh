#!/usr/bin/env bash

function ShowInfo {
	echo "--------------"
	echo "--" $1
	echo "--------------"
}

ShowInfo "Check directory"
if [[ "$(basename "$PWD")" == "go" ]]; then
	echo "Running in the 'go' directory."
else
	echo "Not running in the 'go' directory."
	cd ./go
fi

ShowInfo "Go fmt"
go fmt

if [[ -z $GITHUB_ACTIONS ]]; then
	version=$(git describe)
	# version=1.2.3
fi

ShowInfo "Version $version"

ShowInfo "Build"
# CGO_ENABLED=0 go build -ldflags "-X 'aww/aww.Version=$version'"
CGO_ENABLED=0 go build

ShowInfo "Build version"
# ./aww --version
./aww

# ShowInfo "Test"
# go test ./mgc_test
# # go test -v ./mgc_test
