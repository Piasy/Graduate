#!/bin/sh
CWD=`pwd`
cd .. && export GOPATH=`pwd`:$GOPATH && cd $CWD
#bee run
#go test -bench . -benchmem -parallel 10 tests/
bash
