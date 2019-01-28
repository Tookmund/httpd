#!/bin/sh
mkdir -p build

for o in windows linux darwin freebsd
do
	for a in 386 amd64 arm arm64
	do
		echo "$o $a"
		GOOS="$o" GOARCH="$a" go build -o "build/httpd-$o-$a"
	done
done
