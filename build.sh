#!/bin/bash
#修改为自己的应用名称
appname=AuGoService
DisplayName=AuGoService
Description="基于Go语言的服务程序，可安装和管理第三方应用程序，可运行于Windows、Linux、Macos、Openwrt等各类操作系统。"
version=0.0.0
versionDir="github.com/xxl6097/go-service-framework/pkg/version"

function getversion() {
  version=$(cat version.txt)
  if [ "$version" = "" ]; then
    version="0.0.0"
    echo $version
  else
    v3=$(echo $version | awk -F'.' '{print($3);}')
    v2=$(echo $version | awk -F'.' '{print($2);}')
    v1=$(echo $version | awk -F'.' '{print($1);}')
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

function tag() {
    version=$(getversion)
    echo "current version:${version}"
    git add .
    git commit -m "release v${version}"
    git tag -a v$version -m "release v${version}"
    git push origin v$version
    echo $version >version.txt
}


function GetLDFLAGS() {
  os_name=$(uname -s)
  #echo "os type $os_name"
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
  #echo "$ldflags"
}

function build_linux_mips_opwnert_REDMI_AC2100() {
  rm -rf bin
  rm -rf ./cmd/app/resource.syso
  GetLDFLAGS
  CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "$ldflags -s -w -linkmode internal" -o ./bin/${appname}_${version}_linux_mipsle ./cmd/app
  bash <(curl -s -S -L http://uuxia.cn:8086/up) ./bin/${appname}_${version}_linux_mipsle
}

function build() {
  rm -rf bin
  rm -rf ./cmd/app/resource.syso
  os=$1
  arch=$2
  GetLDFLAGS
  CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -ldflags "$ldflags -s -w -linkmode internal" -o ./bin/${appname}_v${version}_${os}_${arch} ./cmd/app
  bash <(curl -s -S -L http://uuxia.cn:8086/up) ./bin/${appname}_v${version}_${os}_${arch}
}

function build_win() {
  rm -rf bin
  rm -rf ./cmd/app/resource.syso
  os=$1
  arch=$2
  GetLDFLAGS
  go generate ./cmd/app
  CGO_ENABLED=1 GOOS=${os} GOARCH=${arch} go build -ldflags "$ldflags -s -w -linkmode internal" -o ./bin/${appname}_v${version}_${os}_${arch}.exe ./cmd/app
  bash <(curl -s -S -L http://uuxia.cn:8086/up) ./bin/${appname}_v${version}_${os}_${arch}.exe
}


function build_windows_arm64() {
  rm -rf bin
  rm -rf ./cmd/app/resource.syso
  GetLDFLAGS
  CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "$ldflags -s -w -linkmode internal" -o ./bin/${appname}_${version}_windows_arm64.exe ./cmd/app
  bash <(curl -s -S -L http://uuxia.cn:8086/up) ./bin/${appname}_${version}_windows_arm64.exe
}

function menu() {
  echo "1. 编译 Windows amd64"
  echo "2. 编译 Windows arm64"
  echo "3. 编译 Linux amd64"
  echo "4. 编译 Linux arm64"
  echo "5. 编译 Linux mips"
  echo "6. 编译 Darwin arm64"
  echo "7. 编译 Darwin amd64"
  echo "请输入编号:"
  read index
  tag
  case "$index" in
  [1]) (build_win windows amd64) ;;
  [2]) (build_windows_arm64) ;;
  [3]) (build linux amd64) ;;
  [4]) (build linux arm64) ;;
  [5]) (build_linux_mips_opwnert_REDMI_AC2100) ;;
  [6]) (build darwin arm64) ;;
  [7]) (build darwin amd64) ;;
  *) echo "exit" ;;
  esac
}
menu

