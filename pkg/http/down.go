package http

import (
	"fmt"
	"github.com/xxl6097/go-service-framework/pkg/file"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadFile 下载文件并保存到本地
func DownloadFile(filedir string, url string) (string, error) {
	// 创建HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 检查HTTP请求是否成功
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned HTTP status %v", resp.StatusCode)
	}

	// 获取文件名和扩展名
	fileName, _ := file.GetFileAndExtensionFromURL(url)
	//fmt.Println(filedir, string(filepath.Separator), fileName, extension)
	filePath := fmt.Sprintf("%s%s%s", filedir, string(filepath.Separator), fileName)
	//fmt.Println(filePath)
	// 打开文件准备写入
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 将下载的数据写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func Download(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned HTTP status %v", resp.StatusCode)
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
