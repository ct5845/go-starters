#!/bin/bash

. ./scripts/build_mode.sh

# go build
CGO_ENABLED=0 go build -o ./${BUILD_DIR}/ ./src/main.go

echo -e "\033[34mBUILD_GO:\033[32m Built to \033[33m${BUILD_DIR}\033[0m"