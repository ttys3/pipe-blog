#!/bin/bash

SRC_DIR=$(pwd)
cd console && yarn install && yarn run build && cd "${SRC_DIR}" && \
cd theme && yarn install && yarn run build && cd "${SRC_DIR}" && \
echo 'build asset done'
