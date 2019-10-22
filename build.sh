#!/bin/bash

SRC_DIR=$(pwd)
cd console && yarn install && yarn run build && cd "${SRC_DIR}" && \
cd theme && yarn install && yarn run build && cd "${SRC_DIR}" && \
GOPROXY=https://goproxy.cn,direct go build -o nanoblog -i -v . && \
echo 'build pipe done'
