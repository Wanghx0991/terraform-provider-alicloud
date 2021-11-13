#!/usr/bin/env bash

set -e -o pipefail
cd terraform-provider-alicloud
go env -w GOPROXY=https://goproxy.cn,direct
ls -al
GOOS=linux GOARCH=amd64 go build -o bin/terraform-provider-alicloud