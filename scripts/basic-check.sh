#!/usr/bin/env bash
CurrentPath="$(pwd)"
PrevPath="${GOPATH}/src/github.com/aliyun/terraform-provider-alicloud-prev"

rm -rf "${PrevPath}"
if [ ! -d "${GOPATH}/src/github.com/aliyun" ]; then
  mkdir -p "${GOPATH}/src/github.com/aliyun"
fi

function removed() {
  go mod edit -dropreplace=github.com/aliyun/terraform-provider-alicloud-prev
  go mod edit -droprequire=github.com/aliyun/terraform-provider-alicloud-prev
}

trap removed EXIT

git clone "https://github.com/aliyun/terraform-provider-alicloud" "${PrevPath}"
pushd "${CurrentPath}"

go mod edit -require=github.com/aliyun/terraform-provider-alicloud-prev@v0.0.0
go mod edit -replace github.com/aliyun/terraform-provider-alicloud-prev="${PrevPath}"
export PATH=$PATH:$(go env GOPATH)/bin
go get -t github.com/katbyte/terrafmt
go mod tidy

diffFiles=$(git diff --name-only HEAD^ HEAD)
error=false
for doc in ${diffFiles[@]};
do
  dirname=$(dirname "$doc")
  category=$(basename "$dirname")
  case "$category" in
    "d" | "r")
      grep "https://help.aliyun.com/)\.$" "$doc" > /dev/null
      if [[ "$?" == "0" ]]; then
        echo -e "\033[31mDoc :${doc}: Please input the exact link. Currently it is https://help.aliyun.com/. \033[0m"
        error=true
      fi
      terrafmt diff "$doc" -c
      if [[ "$?" != "0" ]]; then
        echo -e "\033[31mDoc :${doc}: Please correct the terraform example \033[0m"
        error=true
      fi
      ;;
    "alicloud")
      grep "fmt.Println" "$doc" > /dev/null
      if [[ "$?" == "0" ]]; then
        echo -e "\033[31mFile :${doc}: Please Remove the fmt.Println Method! \033[0m"
        error=true
      fi
    ;;
  esac
done

if $error; then
  exit 1
fi
echo -e "\033[32m The Basic Check is Success! \033[0m"
exit 0