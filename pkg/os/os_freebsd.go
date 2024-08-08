package os

import (
	"io"
	"os"
	"syscall"
)

func IsExecutable(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	file.Stat()
	defer file.Close()

	// 读取文件头信息
	header := make([]byte, 4)
	if _, err := io.ReadFull(file, header); err != nil {
		return false, err
	}

	// 检查 ELF 文件头（Linux 和 macOS）
	if header[0] == 0x7f && header[1] == 'E' && header[2] == 'L' && header[3] == 'F' {
		return true, nil
	}

	// 检查 PE 文件头（Windows）
	if header[0] == 'M' && header[1] == 'Z' {
		return true, nil
	}

	// 使用 os.Stat 获取文件信息
	fileInfo, err := file.Stat() //os.Stat(filePath)
	if err != nil {
		return false, err
	}
	// 检查用户执行权限
	if fileInfo.Mode().Perm()&0111 == 0111 {
		return true, nil
	}
	// 如果当前用户是文件的所有者，检查所有者执行权限
	if fileInfo.Sys().(*syscall.Stat_t).Uid == uint32(os.Getuid()) {
		return (fileInfo.Mode().Perm() & 0100) == 0100, nil
	}

	// 检查组执行权限
	if fileInfo.Sys().(*syscall.Stat_t).Gid == uint32(os.Getgid()) {
		return (fileInfo.Mode().Perm() & 0010) == 0010, nil
	}

	// 检查其他用户的执行权限
	return (fileInfo.Mode().Perm() & 0001) == 0001, nil
}
