#!/bin/bash

# 配置变量
 REPO="xxl6097/go-service-framework"  # 替换为你的GitHub仓库
TAG="v2.1.4"  # 替换为你的标签
RELEASE_NAME="Release v2.1.4"  # 替换为你的发布名称
DESCRIPTION="This is the release description."  # 替换为你的发布描述
TOKEN=$(cat .token)
FILES=("./dist/AuGoService_0.2.63_linux_arm64")  # 替换为你要附加的文件路径

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
