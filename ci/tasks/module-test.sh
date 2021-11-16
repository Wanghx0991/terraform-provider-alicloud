#!/usr/bin/env bash

set -e

: ${ALICLOUD_ACCESS_KEY:?}
: ${ALICLOUD_SECRET_KEY:?}
: ${ALICLOUD_REGION:?}
: ${terraform_version:?}

export ALICLOUD_ACCESS_KEY=${ALICLOUD_ACCESS_KEY}
export ALICLOUD_SECRET_KEY=${ALICLOUD_SECRET_KEY}
export ALICLOUD_REGION=${ALICLOUD_REGION}

echo "ALICLOUD_ACCESS_KEY=${ALICLOUD_ACCESS_KEY}"
echo "ALICLOUD_SECRET_KEY=${ALICLOUD_SECRET_KEY}"
echo "ALICLOUD_REGION=${ALICLOUD_REGION}"

echo "${PWD}"
cd ./terraform-provider-alicloud
echo "${PWD}"
ls -al
mv bin/terraform-provider-alicloud /usr/bin/
echo "${PWD}"

provider_dir="$(pwd)"
diffFiles=$(git diff --name-only HEAD~ HEAD)
rm -rf ./terraform_test
git clone https://github.com/Wanghx0991/terraform_test
test_dir="$( cd ./terraform_test && pwd )"



terraform init || exit 1



for fileName in ${diffFiles[*]};
do
  if [[ ${fileName} == "alicloud/resource_alicloud"* || ${fileName} == "alicloud/data_source_alicloud"* ]];then
    if [[ ${fileName} == *?_test.go ]]; then
        echo -e "\033[33m[SKIPPED]\033[0m skipping the file $fileName, continue..."
        continue
    fi
    echo "${fileName}"
    fileName=(${fileName//\.go/_test\.go })
    checkResourceName=$(grep "resourceId := \"alicloud_.*.default\""  ${fileName} | grep -Eo 'alicloud[a-z_]*'| head -n +1)
    echo -e "\033[33m[Info]\033[0m file name = ${fileName} Resource Name = ${checkResourceName}"
    cd "${test_dir}" || exit
    make build || exit
    chmod +rx ./bin/terraform_test || exit
    chmod +rx ./scripts/module.sh  || exit
    ./bin/terraform_test  module_test -r="${checkResourceName}" || exit
  fi
done

