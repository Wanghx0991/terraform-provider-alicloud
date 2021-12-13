#!/bin/bash

CurrentPath="$(pwd)"

pushd "${CurrentPath}"


error=false

diffFiles=$(git diff --name-only HEAD~ HEAD)
for fileName in ${diffFiles[@]};
do
    if [[ ${fileName} == "alicloud/resource_alicloud"* ]];then
        if [[ ${fileName} == *?_test.go ]]; then
            echo -e "\033[33m[SKIPPED]\033[0m skipping the file $fileName, continue..."
            continue
        fi
        resourceName=$(echo ${fileName} | grep -Eo "alicloud_[a-z_]*") || exit 1
        echo -e "\033[33mThe ResourceName = ${resourceName}"
        git diff HEAD^ HEAD  > git_diff.diff
        go test -v ./scripts/git_diff_test.go -run=TestFieldCompatibilityCheck -file_name="../git_diff.diff"
        if [[ "$?" != "0" ]]; then
          echo -e "\033[31m ${resourceName}: Compatibility Error! Please check out the correct schema \033[0m"
          error=true
        fi

        go test -v ./scripts/git_diff_test.go -run=TestFieldCompatibilityCheck -resource="${resourceName}"
        if [[ "$?" != "0" ]]; then
          echo -e "\033[31m ${resourceName}: Compatibility Error! Please check out the correct schema \033[0m"
          error=true
        fi
    fi
done


if $error; then
  exit 1
fi
exit 0