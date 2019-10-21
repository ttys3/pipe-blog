#!/bin/bash

cd console && yarn install && yarn run build
cd ../theme && yarn install && yarn run build

GOPROXY=https://goproxy.cn,direct go build -i -v

echo 'build pipe done'
