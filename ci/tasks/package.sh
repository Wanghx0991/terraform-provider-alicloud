#!/usr/bin/env bash

set -e -o pipefail
cd terraform-provider-alicloud
go get golang.org/x/tools/cmd/goimports
#go env -w GOPROXY=https://goproxy.cn,direct
#GOOS=linux GOARCH=amd64 go build -o bin/terraform-provider-alicloud
GOOS=linux  GOARCH=amd64  go build -o bin/terraform-provider-alicloud
cd bin/
ls -al