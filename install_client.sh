#! /bin/bash
# create the binary for the chopper client

set -e
set -x

pushd $GOPATH/src/github.com/dougfort/chopper/chopclient
go install -race
popd
