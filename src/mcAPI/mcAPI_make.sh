#!/bin/sh

export GOPATH=/usr/local/gopath
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOBIN

# make
cd /usr/local/gopath/src/mcAPI;go install mcAPI;cd -

# gen docs
cd /usr/local/gopath/src/mcAPI;bee generate docs;cd -

# copy
cd /usr/local/gopath/bin;cp -rf mcAPI /usr/local/globalways/memberCard_API/;cd -
