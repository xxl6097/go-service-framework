package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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

func IsUrlOrLocalFile(fileurl string) bool {
	if IsValidURL(fileurl) || IsLocalPath(fileurl) {
		return true
	}
	return false
}

func IsLocalPath(path string) bool {
	_, err := os.Stat(path)
	// os.IsNotExist(err) returns true if the error is caused by a non-existing file or directory.
	return !os.IsNotExist(err)
}

func IsValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil {
		return false
	}

	return u.Scheme == "http" || u.Scheme == "https"
}

func IsDirOrFileExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// 文件夹/文件存在
		return true
	}
	return false
}

// IsNotExist 判断文件不存在，返回true
func IsNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsDir()
}

func ReadFile(rawPath string) []byte {
	if IsNotExist(rawPath) {
		return nil
	}
	fileContent, err := os.ReadFile(rawPath)
	if err != nil {
		return nil
	}
	return fileContent
}

func SaveFile(relPath string, body []byte) error {
	if IsNotExist(relPath) {
		return errors.New("file not exist" + relPath)
	}
	err := ioutil.WriteFile(relPath, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetFileNameNoExt(path string) string {
	filenameWithExt := filepath.Base(path)
	filename := strings.TrimSuffix(filenameWithExt, filepath.Ext(filenameWithExt))
	return filename
}

func ScanDirectory(dirPath string) ([]string, error) {
	var files []string
	// 使用 os.ReadDir 读取目录中的一级文件和子目录
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	// 遍历目录项并添加到文件列表中
	for _, entry := range entries {
		files = append(files, filepath.Join(dirPath, entry.Name()))
	}
	return files, nil
}

func ScanDirectoryAndFunc(dirPath string, f func(string)) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return
	}
	// 遍历目录项并添加到文件列表中
	for _, entry := range entries {
		if f != nil {
			f(entry.Name())
		}
	}
	return
}
