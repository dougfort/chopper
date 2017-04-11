#! /bin/bash
# create the binary for the chopper server

set -e
set -x

pushd $GOPATH/src/github.com/dougfort/chopper/chopserv
go install -race
popd
