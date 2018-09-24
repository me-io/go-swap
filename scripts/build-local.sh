#!/bin/sh

export CGO_ENABLED=0
my_dir="$(dirname "$0")"

go build -o ${my_dir}/../bin/server cmd/server/*.go
