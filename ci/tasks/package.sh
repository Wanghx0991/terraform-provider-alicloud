#!/usr/bin/env bash

set -e -o pipefail
cd terraform-provider-alicloud
#go env -w GOPROXY=https://goproxy.cn,direct
make build
cd bin/
ls -al