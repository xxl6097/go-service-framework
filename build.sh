#!/bin/bash
#修改为自己的应用名称
appname=AuGoService
DisplayName=基于Golang后台服务管理程序
Description="基于Go语言的服务程序，可安装和管理第三方应用程序，可运行于Windows、Linux、Macos、Openwrt等各类操作系统。"
version=0.0.0
versionDir="github.com/xxl6097/go-service-framework/pkg/version"
appdir="./cmd/app"

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


function github_release() {
  echo "开发发布github release"
    # 配置变量
    REPO="xxl6097/go-service-framework"  # 替换为你的GitHub仓库
    TAG="${version}"  # 替换为你的标签
    RELEASE_NAME="${version}"  # 替换为你的发布名称
    DESCRIPTION="This is the release description."  # 替换为你的发布描述
    TOKEN=$(cat .token)  # 替换为你的GitHub Token
#    FILE_PATH="./dist/AuGoService_0.2.23_windows_arm64.exe"  # 替换为你要附加的文件路径
    #FILES=("./dist/AuGoService_0.2.23_windows_arm64.exe" "./dist/AuGoService_v0.2.23_darwin_amd64")  # 替换为你要附加的文件路径
    # 定义要扫描的目录
    DIRECTORY="./dist"
    echo "token=>$TOKEN"
    # 初始化一个空数组
    FILES=()
    # 使用find命令扫描目录，并将结果添加到数组中
    while IFS= read -r file; do
        FILES+=("$file")
    done < <(find "$DIRECTORY" -type f)
    # 打印数组内容
    echo "Found files:"
    printf '%s\n' "${FILES[@]}"

    # 创建一个新的release
    response=$(curl -s -X POST \
      -H "Authorization: token $TOKEN" \
      -H "Accept: application/vnd.github.v3+json" \
      https://api.github.com/repos/$REPO/releases \
      -d "{
        \"tag_name\": \"$TAG\",
        \"target_commitish\": \"main\",
        \"name\": \"$RELEASE_NAME\",
        \"body\": \"$DESCRIPTION\",
        \"draft\": false,
        \"prerelease\": false
      }")

    # 提取release的上传URL
    upload_url=$(echo "$response" | jq -r .upload_url | sed -e "s/{?name,label}//")

    # 检查创建release是否成功
    if [ "$upload_url" == "null" ]; then
      echo "Failed to create release"
      echo "$response"
      exit 1
    fi

    # 上传附件文件
    for FILE_PATH in "${FILES[@]}"; do
      FILE_NAME=$(basename "$FILE_PATH")
      echo "Uploading $FILE_NAME..."
      curl -s -X POST \
        -H "Authorization: token $TOKEN" \
        -H "Content-Type: $(file -b --mime-type "$FILE_PATH")" \
        --data-binary @"$FILE_PATH" \
        "$upload_url?name=$FILE_NAME"
      echo "$FILE_NAME uploaded successfully."
    done

    echo "All files uploaded successfully."
    #echo $version >version.txt
}

function build_linux_mips_opwnert_REDMI_AC2100() {
  echo "开始编译 linux mipsle ${appname}_${version}"
  CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "$ldflags -s -w -linkmode internal" -o ./dist/${appname}_${version}_linux_mipsle ${appdir}
}

function build() {
  os=$1
  arch=$2
  echo "开始编译 ${os} ${arch} ${appname}_${version}"
  CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -ldflags "$ldflags -s -w -linkmode internal" -o ./dist/${appname}_${version}_${os}_${arch} ${appdir}
}

function build_win() {
  os=$1
  arch=$2
  echo "开始编译 ${os} ${arch} ${appname}_${version}"
  go generate ${appdir}
  CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -ldflags "$ldflags -s -w -linkmode internal" -o ./dist/${appname}_${version}_${os}_${arch}.exe ${appdir}
  rm -rf ${appdir}/resource.syso
}


function build_windows_arm64() {
  echo "开始编译 windows arm64 ${appname}_${version}"
  CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "$ldflags -s -w -linkmode internal" -o ./dist/${appname}_${version}_windows_arm64.exe ${appdir}
}

function build_menu() {
  my_array=("$@")
  for index in "${my_array[@]}"; do
        case "$index" in
          [1]) (build_win windows amd64) ;;
          [2]) (build_windows_arm64) ;;
          [3]) (build linux amd64) ;;
          [4]) (build linux arm64) ;;
          [5]) (build_linux_mips_opwnert_REDMI_AC2100) ;;
          [6]) (build darwin arm64) ;;
          [7]) (build darwin amd64) ;;
          *) echo "-->exit" ;;
          esac
  done

  bash <(curl -s -S -L http://uuxia.cn:8087/up) ./dist /soft/${appname}/${version}
  github_release
}

function buildArgs() {
  os_name=$(uname -s)
  #echo "os type $os_name"
  APP_NAME=${appname}
  BUILD_VERSION=$(if [ "$(git describe --tags --abbrev=0 2>/dev/null)" != "" ]; then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
  BUILD_TIME=$(TZ=Asia/Shanghai date +%FT%T%z)
  GIT_REVISION=$(git rev-parse --short HEAD)
  GIT_BRANCH=$(git name-rev --name-only HEAD)
  GO_VERSION=$(go version)
  ldflags="-s -w\
 -X '${versionDir}.AppName=${APP_NAME}'\
 -X '${versionDir}.DisplayName=${DisplayName}'\
 -X '${versionDir}.Description=${Description}'\
 -X '${versionDir}.AppVersion=${BUILD_VERSION}'\
 -X '${versionDir}.BuildVersion=${BUILD_VERSION}'\
 -X '${versionDir}.BuildTime=${BUILD_TIME}'\
 -X '${versionDir}.GitRevision=${GIT_REVISION}'\
 -X '${versionDir}.GitBranch=${GIT_BRANCH}'\
 -X '${versionDir}.GoVersion=${GO_VERSION}'"
  #echo "$ldflags"
}

function initArgs() {
  version=$(getversion)
  echo "version:${version}"
  rm -rf dist
  tagAndGitPush
  buildArgs
}

function tagAndGitPush() {
    git add .
    git commit -m "release ${version}"
    git tag -a v$version -m "release ${version}"
    git push origin v$version
    echo $version >version.txt
}

# shellcheck disable=SC2120
function menu() {
  echo "1. 编译 Windows amd64"
  echo "2. 编译 Windows arm64"
  echo "3. 编译 Linux amd64"
  echo "4. 编译 Linux arm64"
  echo "5. 编译 Linux mips"
  echo "6. 编译 Darwin arm64"
  echo "7. 编译 Darwin amd64"
  echo "8. 编译全平台"
  echo "9. github release"
  echo "请输入编号:"
  read -r -a inputData "$@"
  initArgs
  if (( inputData[0] == 8 )); then
     array=(1 2 3 4 5 6 7)
     (build_menu "${array[@]}")
  elif (( inputData[0] == 9 )); then
       github_release
  else
     (build_menu "${inputData[@]}")
  fi
  echo $version >version.txt
}
menu

