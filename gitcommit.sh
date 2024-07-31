#!/bin/bash
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

function todir() {
  pwd
}

function pull() {
  todir
  echo "git pull"
  git pull
}

function forcepull() {
  todir
  echo "git fetch --all && git reset --hard origin/main && git pull"
  git fetch --all && git reset --hard origin/main && git pull
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
#  shellcheck disable=SC2120
function gitpush() {
  commit=""
  if [ ! -n "$1" ]; then
    commit="$(date '+%Y-%m-%d %H:%M:%S') by ${USER}"
  else
    commit="$1 by ${USER}"
  fi

  echo $commit
  tag
  git add .
  git commit -m "$commit"
  #  git push -u origin main
  git push
  #stag
}

function test() {
    version=$(getversion)
    echo "current version:${version}"
    echo $version >version.txt
}

function github_release() {
  # 配置
    version=$(getversion)
    # 配置变量
    REPO="xxl6097/go-service-framework"  # 替换为你的GitHub仓库
    TAG="${version}"  # 替换为你的标签
    RELEASE_NAME="${version}"  # 替换为你的发布名称
    DESCRIPTION="This is the release description."  # 替换为你的发布描述
    TOKEN="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"  # 替换为你的GitHub Token
#    FILE_PATH="./dist/AuGoService_0.2.23_windows_arm64.exe"  # 替换为你要附加的文件路径
    #FILES=("./dist/AuGoService_0.2.23_windows_arm64.exe" "./dist/AuGoService_v0.2.23_darwin_amd64")  # 替换为你要附加的文件路径
    # 定义要扫描的目录
    DIRECTORY="./dist"
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
    echo $version >version.txt
}

function m() {
    echo "1. 强制更新"
    echo "2. 普通更新"
    echo "3. 提交项目"
    echo "4. 测试"
    echo "5. github release"
    echo "请输入编号:"
    read index

    case "$index" in
    [1]) (forcepull);;
    [2]) (pull);;
    [3]) (gitpush);;
    [4]) (test);;
    [5]) (github_release);;
    *) echo "exit" ;;
  esac
}

function bootstrap() {
    case $1 in
    pull) (pull) ;;
    m) (m) ;;
      -f) (forcepull) ;;
       *) ( gitpush $1)  ;;
    esac
}

bootstrap m
