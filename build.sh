#!/bin/bash

GOPROXY=https://goproxy.cn,direct go build -i -v
cd console && yarn install && yarn run build
cd ../theme && yarn install && yarn run build

echo 'build pipe done'
