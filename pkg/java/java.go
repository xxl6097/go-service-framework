package java

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func HasSuffix(filename, suffix string) bool {
	return filepath.Ext(filename) == suffix
}
func IsJar(file string) bool {
	if HasSuffix(file, ".jar") {
		return true
	}
	return false
}

func IsJava(name string) bool {
	if strings.EqualFold(name, "java") {
		return true
	}
	return false
}
func WhereIsJava() (string, error) {
	cmd := exec.Command("which", "java")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	javaPath := strings.TrimSpace(string(output))
	// 根据java命令的路径找到Java的安装路径
	// 这可能需要一些字符串操作，取决于你的具体需求
	return javaPath, nil
}

func IsExist(file string) bool {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		return true
	}
	return false
}
func FindJavaPath() (string, error) {
	pathEnv, _ := exec.LookPath("java")
	if IsExist(pathEnv) {
		return pathEnv, nil
	}
	javaHome := os.Getenv("JAVA_HOME")
	if javaHome != "" {
		fmt.Println("Using JAVA_HOME: " + javaHome)
		javaPath := filepath.Join(javaHome, "bin")
		if IsExist(javaPath) {
			return pathEnv, nil
		}
	}

	path, err := WhereIsJava()
	if err == nil && path != "" {
		if IsExist(path) {
			return pathEnv, nil
		}
	}
	return "", errors.New("not found")
}
