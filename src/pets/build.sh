#!/bin/bash

VERSION=`cat version`
go build -ldflags "-X pets.Version=$VERSION" -o bin/pets ./cmd/pets
