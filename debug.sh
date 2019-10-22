#!/bin/bash

rm -rf ./nanoblog && GOPROXY=https://goproxy.cn,direct go build -o nanoblog -i -v && ./nanoblog
