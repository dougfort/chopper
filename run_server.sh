#! /bin/bash
# run the binary for the chopper server

set -e
set -x

$GOPATH/bin/chopserv 2> $HOME/chopserv.log &
