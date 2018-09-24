#!/bin/sh

export CGO_ENABLED=0
my_dir="$(dirname "$0")"

ls -al /
which go
which find
which bash
which heroku
which docker

go build -o ${my_dir}/../bin/go-swap-server cmd/server/*.go
