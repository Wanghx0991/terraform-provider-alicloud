#!/usr/bin/env bash

set -e -o pipefail
cd terraform-provider-alicloud
GOOS=linux GOARCH=amd64 go build -o bin/terraform-provider-alicloud