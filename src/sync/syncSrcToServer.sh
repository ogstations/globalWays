#!/bin/sh


#modify rootPath if possible
sourcePath1="/Users/mint/develop/work/globalWays/src/"
targetPath1="/usr/local/gopath/src"

excludeFiles=".DS_Store"

rm -rf /Users/mint/develop/work/globalWays/src/mcAPI/mcAPI

#default server address
server="106.185.38.123" #staging server

dest_user="root"
dest_passwd="Zhaoming123"

currPath=`dirname $0`

if [ $# -gt 1 ];then
  server=$3
  if [ "$server" = "release" ];then
  	server=""
  	dest_passwd=""
  fi
fi

echo "currPath:$currPath"
echo "server: "$server

echo "$currPath/rsyncd.exp ${server} ${destPath} ${dest_user} ${dest_passwd} ${sourcePath} 100000"
$currPath/rsyncd_no_delete.exp "${server}" "${targetPath1}" "${dest_user}" "${dest_passwd}" "${sourcePath1}" "100000" "${excludeFiles}"
