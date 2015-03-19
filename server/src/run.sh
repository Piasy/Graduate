#!/bin/sh
CWD=`pwd`
cd .. && export GOPATH=`pwd`:$GOPATH && cd $CWD

bee run
