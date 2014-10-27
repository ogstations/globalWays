#!/bin/sh

export GOPATH=/usr/local/gopath
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOBIN

# stop
count=`ps -ef|grep -w "mcAPI"|grep -v grep|grep -v "\.log"| wc -l`
if [ $count -ne 0 ]; then
	echo "mcAPI process is exists, kill it.."
	killall -9 mcAPI
    sleep 1
fi

# run
cd /usr/local/globalways/memberCard_API/;nohup ./mcAPI &
cd -


