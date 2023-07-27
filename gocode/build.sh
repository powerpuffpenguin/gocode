#!/bin/bash
set -e

cd `dirname "$BASH_SOURCE"`

export GOOS=linux
export GOARCH=amd64

go build

tar -zcvf gocode.tar.gz gocode