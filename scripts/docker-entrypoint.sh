#!/bin/sh

echo Your container args are: "$@"
echo Your BINSRC_ENV are: "${BINSRC_ENV}"

exec ${BINSRC_ENV} "$@"

