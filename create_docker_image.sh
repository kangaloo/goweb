#!/usr/bin/env bash

currentDir=$(cd "$(dirname "${BASH_SOURCE-$0}")"; pwd)
cd ${currentDir}

CGO_ENABLE=0 GOOS=linux go build

docker build -t goweb:v1.3 .
