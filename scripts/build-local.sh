#!/bin/sh

export CGO_ENABLED=0
my_dir="$(dirname "$0")"

ls -al /
ls -al /bin/
ls -al /usr/local/
ls -al /usr/local/heroku/
ls -al /usr/local/heroku/bin
go build -o ${my_dir}/../bin/go-swap-server cmd/server/*.go
