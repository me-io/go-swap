#!/usr/bin/env bash

my_dir="$(dirname "$0")"
filename=${1-multi_1}
file=${my_dir}/${filename}.json
host=${2-localhost}
port=${3-5000}
echo "file: ${file}"

curl -X POST -H "Content-Type: application/json" -d @${file} http://${host}:${port}/convert
