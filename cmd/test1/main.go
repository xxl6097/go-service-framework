package main

import (
	"fmt"
	"github.com/xxl6097/go-service-framework/pkg/file"
	"github.com/xxl6097/go-service-framework/pkg/java"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func findJavaPath() (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("where", "java")
	} else {
		cmd = exec.Command("which", "java")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	} else {
		// Trim any newlines from output
		return strings.TrimSpace(string(out)), nil
	}
}
func hasSuffix(filename, suffix string) bool {
	return filepath.Ext(filename) == suffix
}
func main() {
	urls := "java"
	fmt.Println(file.IsUrlOrLocalFile(urls))
	fpath, _ := os.Executable()
	fpath = "http://uuxia.cn:8086/files/2024/07/12/gtbx-tcp-server-0.0.0.jar"
	fmt.Println("扩展名:", filepath.Ext(fpath))

	jpath, _ := java.FindJavaPath()

	fmt.Println("jpath:", jpath)
	jpatha, _ := findJavaPath()
	fmt.Println("jpatha:", jpatha)
}
