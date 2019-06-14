#!/bin/bash
set -e -u -x

dir=$(cd $(dirname "${0}") && pwd)
cd "${dir}"

goversion="1.12.6"
rootpath=$(git rev-parse --show-toplevel)
package=$(realpath --relative-to=$(go env GOPATH)/src $(git rev-parse --show-toplevel))

docker run \
    -i \
    -t \
    --rm \
    -v "${rootpath}:/go/src/${package}" \
    -w "/go/src/${package}" \
    "golang:${goversion}" \
    bash -e -u -x -c "./build.sh && chown -R $(id -u):$(id -g) build"
