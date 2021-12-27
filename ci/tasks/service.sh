#!/usr/bin/env bash

set -e

: "${ALICLOUD_ACCESS_KEY:?}"
: "${ALICLOUD_SECRET_KEY:?}"
: "${ALICLOUD_REGION:?}"

export ALICLOUD_ACCESS_KEY=${ALICLOUD_ACCESS_KEY}
export ALICLOUD_SECRET_KEY=${ALICLOUD_SECRET_KEY}
export ALICLOUD_REGION=${ALICLOUD_REGION}


CURRENT_PATH=${PWD}
echo -e "ls -l CURRENT_PATH"
serviceFiles=ls alicloud | grep 'data_[a-z_]*service.go'

provider="terraform-provider-alicloud"

cd terraform-provider-alicloud

for fileName in ${serviceFiles[@]};
do
    echo -e "\n\033[37mchecking diff file $fileName ... \033[0m"
    if [[ ${fileName} == "alicloud/resource_alicloud"* || ${fileName} == "alicloud/data_source_alicloud"* ]];then
        fileName=(${fileName//\.go/_test\.go })
        checkFuncs=$(grep "func TestAcc.*" ${fileName})
        echo -e "found the test funcs:\n${checkFuncs}\n"
        funcs=(${checkFuncs//"(t *testing.T) {"/ })
        for func in ${funcs[@]};
        do
          if [[ ${func} != "TestAcc"* ]]; then
            continue
          fi
          go clean -cache -modcache -i -r
          echo -e "\033[34m################################################################################\033[0m"
          echo -e "\033[34mTF_ACC=1 go test ./alicloud -v -run=${func} -timeout=1200m\033[0m"
          TF_ACC=1 go test ./alicloud -v -run=${func} -timeout=1200m | {
          while read LINE
          do
              echo -e "$LINE"
              if [[ $LINE == "--- FAIL: "* || ${LINE} == "FAIL"* ]]; then
                  FAILED_COUNT=$((${FAILED_COUNT}+1))
              fi
              if [[ $LINE == "panic: "* ]]; then
                  FAILED_COUNT=$((${FAILED_COUNT}+1))
                  break
              fi
          done
          # send child var to an failed file
          if [[ $FAILED_COUNT -gt 0 ]]; then
            echo -e "\033[31mrecord the failed count $FAILED_COUNT into a temp file\033[0m"
            echo $FAILED_COUNT > failed.txt
          fi
          }
        done
        echo -e "\033[34mFinished\033[0m"
    fi
done