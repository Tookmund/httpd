#!/bin/sh

for o in windows linux darwin freebsd
do
	for a in 386 amd64 arm arm64
	do
		echo "$o $a"
		GOOS="$o" GOARCH="$a" go build -o "httpd-$o-$a"
	done
done
