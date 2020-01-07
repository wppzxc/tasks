#!/bin/sh

set -e

version=$1
if [ -z $version ]; then
    version=`date "+%Y%m%d%H%M%S"`
    echo  "不指定版本号, 启用默认版本号: $version"
    echo "eg: build/build.sh $version"
    echo
fi
echo "version is : $version"
docker run --rm  -e GO111MODULE=off -v $PWD:/go/src/github.com/wppzxc/tasks golang:1.13.4 go build --ldflags '-w -X github.com/wppzxc/tasks/src/pkg/version.version='$version'' -o /go/src/github.com/wppzxc/tasks/build/docker/tasks /go/src/github.com/wppzxc/tasks/src/main.go
docker build -t wppzxc/tasks:$version ./build/docker