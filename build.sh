#!/bin/bash
#修改为自己的应用名称
appname=AAFrameworkService
DisplayName=AAFrameworkService
Description=基于Go语言的服务程序框架
version=0.0.0
versionDir="github.com/xxl6097/go-service-framework/pkg/version"

function getversion() {
  appversion=$(cat version.txt)
  if [ "$appversion" = "" ]; then
    appversion="0.0.0"
    echo $appversion
  else
    v3=$(echo $appversion | awk -F'.' '{print($3);}')
    v2=$(echo $appversion | awk -F'.' '{print($2);}')
    v1=$(echo $appversion | awk -F'.' '{print($1);}')
    if [[ $(expr $v3 \>= 99) == 1 ]]; then
      v3=0
      if [[ $(expr $v2 \>= 99) == 1 ]]; then
        v2=0
        v1=$(expr $v1 + 1)
      else
        v2=$(expr $v2 + 1)
      fi
    else
      v3=$(expr $v3 + 1)
    fi
    ver="$v1.$v2.$v3"
    echo $ver
  fi
}


function GetLDFLAGS() {
  os_name=$(uname -s)
  echo "os type $os_name"
  APP_NAME=${appname}
  APP_VERSION=${appversion}
  BUILD_VERSION=$(if [ "$(git describe --tags --abbrev=0 2>/dev/null)" != "" ]; then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
  BUILD_TIME=$(TZ=Asia/Shanghai date +%FT%T%z)
  GIT_REVISION=$(git rev-parse --short HEAD)
  GIT_BRANCH=$(git name-rev --name-only HEAD)
  GO_VERSION=$(go version)
  ldflags="-s -w\
 -X '${versionDir}.AppName=${APP_NAME}'\
 -X '${versionDir}.DisplayName=${DisplayName}'\
 -X '${versionDir}.Description=${Description}'\
 -X '${versionDir}.AppVersion=${APP_VERSION}'\
 -X '${versionDir}.BuildVersion=${BUILD_VERSION}'\
 -X '${versionDir}.BuildTime=${BUILD_TIME}'\
 -X '${versionDir}.GitRevision=${GIT_REVISION}'\
 -X '${versionDir}.GitBranch=${GIT_BRANCH}'\
 -X '${versionDir}.GoVersion=${GO_VERSION}'"
  echo "$ldflags"
}

function build_windows_amd64() {
  #goversioninfo -manifest versioninfo.json
  rm -rf ${appname}_${version}_windows_amd64.exe
  GetLDFLAGS
  go generate
  #CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-linkmode internal" -o ${appname}_${version}_windows_amd64.exe
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "$ldflags -s -w -linkmode internal" -o ./bin/${appname}_${version}_windows_amd64.exe ./cmd/app
}


function menu() {
  echo -e "\r\n0. 编译 Windows amd64"
  echo "请输入编号:"
  read index
  case "$index" in
  [0]) (build_windows_amd64) ;;
  *) echo "exit" ;;
  esac

  if ((index >= 4 && index <= 6)); then
    # 获取命令的退出状态码
    exit_status=$?
    # 检查退出状态码
    if [ $exit_status -eq 0 ]; then
      echo "成功推送Docker"
      echo $appversion >version.txt
    else
      echo "失败"
      echo "【$docker_push_result】"
    fi
  fi
}
menu

