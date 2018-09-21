#!/usr/bin/env bash

my_dir="$(dirname "$0")"
filename=${1-multi_1}
file=${my_dir}/${filename}.json

echo "file: ${file}"

curl -X POST -H "Content-Type: application/json" -d @${file} http://localhost:5000/convert
