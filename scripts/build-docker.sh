#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

REPO_NAME="meio/go-swap-server"

DOCKER_TAG=`git describe --tags --always --dirty`

if [[ ! -z "${DOCKER_PASSWORD}" && ! -z "${DOCKER_USERNAME}" ]]
then
    echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
fi

TAG_EXIST=`curl -s "https://hub.docker.com/v2/repositories/${REPO_NAME}/tags/${DOCKER_TAG}/" | grep '"id":'`

if [[ ! -z ${TAG_EXIST}  ]]; then
    echo "${REPO_NAME}/${DOCKER_TAG} already exist"
    exit 0
fi

make container

if [[ $? != 0 ]]; then
    echo "${REPO_NAME}/${DOCKER_TAG} build failed"
    exit 1
fi


if [[ -z ${TAG_EXIST}  ]]; then
    docker push ${REPO_NAME}:${DOCKER_TAG}
    echo "${REPO_NAME}/${DOCKER_TAG} pushed successfully"
fi
