#!/bin/bash

set -eu

. /opt/golang/preferred/bin/go_env.sh

export GOPATH="$(pwd)/go"
export PATH="$GOPATH/bin:$PATH"

cd "./$CLONE_PATH"

go get ./...

make test


