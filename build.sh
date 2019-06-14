#!/bin/bash
set -e -u -x

dir=$(cd $(dirname "${0}") && pwd)

appname=$(basename $(pwd))
output="${dir}/build"

upx_version="3.95"
if [ ! -f "${output}/tool/upx-${upx_version}-amd64_linux/upx" ]; then
    mkdir -p "${output}/tool"
    curl -fsSL "https://github.com/upx/upx/releases/download/v${upx_version}/upx-${upx_version}-amd64_linux.tar.xz" | tar xJ -C "${output}/tool"
fi

go build -ldflags "-w -s" -o "${output}/${appname}"
"${output}/tool/upx-${upx_version}-amd64_linux/upx" "${output}/${appname}"
