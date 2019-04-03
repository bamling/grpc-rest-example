#!/bin/bash
SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_PATH=$(realpath "$0")
PROJECT_HOME="$(dirname "${SCRIPT_PATH%$SCRIPT_NAME}")"

# ensure directory exists
mkdir -p ${PROJECT_HOME}/pkg/api

# generate files
docker run --rm \
    -v ${PROJECT_HOME}:/src \
    -u $(id -u ${USER}):$(id -g ${USER}) \
    znly/protoc \
    --proto_path=/src/api/proto \
    --go_out=plugins=grpc:/src/pkg/api \
    --grpc-gateway_out=logtostderr=true:/src/pkg/api \
    --swagger_out=logtostderr=true:/src/api/swagger \
    $(find ${PROJECT_HOME}/api -iname "*.proto" -printf "%f\n")
