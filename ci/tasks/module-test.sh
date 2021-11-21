#!/usr/bin/env bash

set -e

: ${ALICLOUD_ACCESS_KEY:?}
: ${ALICLOUD_SECRET_KEY:?}
: ${ALICLOUD_REGION:?}
: ${terraform_version:?}

export ALICLOUD_ACCESS_KEY=${ALICLOUD_ACCESS_KEY}
export ALICLOUD_SECRET_KEY=${ALICLOUD_SECRET_KEY}
export ALICLOUD_REGION=${ALICLOUD_REGION}


CURRENT_PATH=${PWD}
echo -e "ls -l CURRENT_PATH"
ls -al "${CURRENT_PATH}"

TF_PLUGIN_CACHE_DIR=${PWD}/cache/.terraform/plugins
echo -e "mkdir -p $TF_PLUGIN_CACHE_DIR"
mkdir -p $TF_PLUGIN_CACHE_DIR
export TF_PLUGIN_CACHE_DIR=${TF_PLUGIN_CACHE_DIR}

TERRAFORM_SOURCE_PATH=$CURRENT_PATH/terraform-provider-alicloud
TERRAFORM_TEST_PATH=$CURRENT_PATH/terraform_module_test
TF_NEXT_PROVIDER=$CURRENT_PATH/next-provider/terraform-provider-alicloud

echo "TERRAFORM_SOURCE_PATH = ${TERRAFORM_SOURCE_PATH}"
echo "TERRAFORM_TEST_PATH = ${TERRAFORM_TEST_PATH}"
echo "TF_NEXT_PROVIDER = ${TF_NEXT_PROVIDER}"

apt-get update && apt-get install -y zip
wget -qN https://releases.hashicorp.com/terraform/${terraform_version}/terraform_${terraform_version}_linux_amd64.zip
unzip -o terraform_${terraform_version}_linux_amd64.zip -d /usr/bin

pushd ${TERRAFORM_SOURCE_PATH}
cp "${TERRAFORM_TEST_PATH}" ${TERRAFORM_SOURCE_PATH}
ls -al
echo "cp -r"
cp -r "${TERRAFORM_TEST_PATH}" ${TERRAFORM_SOURCE_PATH}
diffFiles=$(git diff --name-only HEAD~ HEAD)


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
    cd "${TERRAFORM_TEST_PATH}" || exit
    make build || exit
    chmod +rx ./bin/terraform_test || exit
    chmod +rx ./scripts/module.sh  || exit
    ./bin/terraform_test  module_test -r="${checkResourceName}" || exit
  fi
done
