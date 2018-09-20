#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

my_dir="$(dirname "$0")"

GO_FILES=`find ${my_dir}/../cmd/server/. -type f \( -iname "*.go" ! -iname "*_test.go" \)`
go run ${GO_FILES}
