#!/bin/bash
appname=test1
version=0.0.0

function build_windows_amd64() {
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${appname}.exe
  bash <(curl -s -S -L http://uuxia.cn:8086/up) ${appname}.exe
}

function build_linux_amd64() {
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${appname}_${version}_linux_amd64
  bash <(curl -s -S -L http://uuxia.cn:8086/up) ${appname}_${version}_linux_amd64
}

function build_linux_arm64() {
  CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build  -o ${appname}
  bash <(curl -s -S -L http://uuxia.cn:8086/up) ${appname}
}

function build_darwin_arm64() {
  CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build  -o ${appname}
  bash <(curl -s -S -L http://uuxia.cn:8086/up) ${appname}
}


function menu() {
  echo -e "\r\n0. 编译 Windows amd64"
  echo "1. 编译 Linux amd64"
  echo "2. 编译 Linux arm64"
  echo "3. 编译 MacOS"
  echo "4. 打包多平台镜像->DockerHub"
  echo "5. 打包多平台镜像->Coding"
  echo "6. 打包多平台镜像->Tencent"
  echo "7. go mod tidy"
  echo "请输入编号:"
  read index
  case "$index" in
  [0]) (build_windows_amd64) ;;
  [1]) (build_linux_amd64) ;;
  [2]) (build_linux_arm64) ;;
  [3]) (build_darwin_arm64) ;;
  [4]) (build_images_to_hubdocker) ;;
  [5]) (build_images_to_conding) ;;
  [6]) (build_images_to_tencent) ;;
  [7]) (gomodtidy) ;;
  *) echo "exit" ;;
  esac
}

function main() {
  menu
}
main
