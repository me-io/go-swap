#!/bin/sh

my_dir="$(dirname "$0")"

GO_FILES=`find ${my_dir}/../cmd/server/. -type f \( -iname "*.go" ! -iname "*_test.go" \)`
REDIS_URL=redis://localhost:6379 cache=redis go run ${GO_FILES}
