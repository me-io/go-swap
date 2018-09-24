#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

my_dir="$(dirname "$0")"

docker pull meio/go-swap-server:latest
docker run --rm --name go-swap-server -u 0 -p 5000:5000 -it meio/go-swap-server:latest
