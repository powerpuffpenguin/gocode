#!/bin/bash
set -e

cd `dirname "$BASH_SOURCE"`

export GOOS=linux
export GOARCH=amd64

go build

tar -zcvf linux_amd64.tar.gz gocode