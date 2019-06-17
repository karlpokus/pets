#!/bin/bash

# build the pets binary for the specified os

# USAGE
# ./build.sh <os>

# set cwd to wherever this script is located
cd "$(dirname "$0")"

OS=$1
VERSION=`cat version`

if test -z $OS; then
  echo "missing arg"
  exit 1
fi

if test $OS != "darwin" && test $OS != "linux"; then
  echo "unknown os"
  exit 1
fi

echo "building pets $VERSION for $OS"
GOOS=$OS GOARCH=amd64 go build -ldflags "-X pets.Version=$VERSION" -o bin/$OS/pets ./cmd/pets
echo done
