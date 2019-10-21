#!/bin/bash

GOPROXY=https://goproxy.cn,direct go build -i -v && ./pipe
