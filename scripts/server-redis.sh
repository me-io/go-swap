#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

my_dir="$(dirname "$0")"

GO_FILES=`find ${my_dir}/../cmd/server/. -type f \( -iname "*.go" ! -iname "*_test.go" \)`
cmd="go run ${GO_FILES} -CACHE=redis -REDIS_URL=redis://localhost:6379"
${cmd}
