#!/bin/bash

rm -rf ./pipe && GOPROXY=https://goproxy.cn,direct go build -i -v && ./pipe
