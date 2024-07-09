package file

import (
	"fmt"
	"net/url"
	"path/filepath"
)

func GetFileAndExtensionFromURL(rawurl string) (string, string) {
	// 解析URL
	u, err := url.Parse(rawurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", ""
	}

	// 获取URL的路径部分
	path := u.Path

	// 从路径中提取文件名,包含了后缀
	filename := filepath.Base(path)

	// 获取文件扩展名
	ext := filepath.Ext(filename)

	return filename, ext
}
