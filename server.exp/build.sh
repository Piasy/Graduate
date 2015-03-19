#!/bin/sh
export GOPATH=`pwd`:$GOPATH
cd bin
go build types && go install types && \
go build dbhelper && go install dbhelper && \
go build ../src/server.go && echo `pwd` && bash
