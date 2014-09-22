#!/bin/sh
#手动做成 by mint
#时间：2014/02/27 00:02
#作用：go语言环境变量加载

export GOPATH=$HOME/develop/work/globalWays
export GOBIN=$GOPATH/bin
export GOROOT=$HOME/develop/tools/go-1.3
export PATH=$PATH:$GOROOT/bin:$GOBIN
