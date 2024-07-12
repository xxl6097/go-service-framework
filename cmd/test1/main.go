package main

import (
	"fmt"
	"github.com/xxl6097/go-service-framework/pkg/file"
	"os"
	"path/filepath"
)

func hasSuffix(filename, suffix string) bool {
	return filepath.Ext(filename) == suffix
}
func main() {
	urls := "java"
	fmt.Println(file.IsUrlOrLocalFile(urls))
	fpath, _ := os.Executable()
	fpath = "http://uuxia.cn:8086/files/2024/07/12/gtbx-tcp-server-0.0.0.jar"
	fmt.Println("扩展名:", filepath.Ext(fpath))
}
