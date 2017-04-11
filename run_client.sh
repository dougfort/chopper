#! /bin/bash
# run the binary for the chopper client

set -e
set -x

$GOPATH/bin/chopclient 2> $HOME/chopclient.log
