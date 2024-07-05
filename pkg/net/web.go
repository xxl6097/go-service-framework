package net

import (
	"net/url"
	"os/exec"
	"runtime"
)

// OpenURL 使用系统默认的浏览器打开给定的 URL。
func OpenURL(urls string) error {
	// 将 URL 进行编码以确保它符合 URL 规范
	parsedURL, err := url.Parse(urls)
	if err != nil {
		return err
	}

	// 根据操作系统选择不同的命令
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", parsedURL.String())
	case "darwin":
		cmd = exec.Command("open", parsedURL.String())
	default: // 包括 Linux 和其他 Unix 系统
		cmd = exec.Command("xdg-open", parsedURL.String())
	}

	// 执行命令
	return cmd.Run()
}
