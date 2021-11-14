#!/usr/bin/env bash

set -e -o pipefail
cd terraform-provider-alicloud
go get golang.org/x/tools/cmd/goimports
#go env -w GOPROXY=https://goproxy.cn,direct

echo "Starting Go mod tidy"
go mod tidy
echo "Ending Go mod tidy"
go build -o bin/terraform-provider-alicloud
cd bin/
ls -al