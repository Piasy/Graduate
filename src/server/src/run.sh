#!/bin/sh
CWD=`pwd`
cd .. && export GOPATH=`pwd`:$GOPATH && cd $CWD
#mongod --dbpath ~/data/mongodb > /dev/null 

#bee run
#go test -bench . -benchmem -parallel 10 tests/
bash
