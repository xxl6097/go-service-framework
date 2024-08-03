package main

import (
	"github.com/xxl6097/go-service-framework/internal/framework"
	os1 "github.com/xxl6097/go-service-framework/pkg/os"
	"github.com/xxl6097/go-service/gservice"
)

func init() {
	if os1.IsMacOs() {
		os1.SetDebug(true)
	}
	//os1.SetDebug(true)
}

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
	test()
	gservice.Run(&framework.Framework{})
}

func test() {
	//binDir := "/Users/uuxia/Downloads/sun-panel"
	//binPath := "/Users/uuxia/Downloads/sun-panel/sun-panel.zip"
	//isZip := zip.IsZip(binPath)
	//if isZip {
	//	//确定解压成功
	//	zipDir, err := zip.GetRootDir(binDir)
	//	if err == nil && zipDir != "" {
	//		//确定zip有一级目录
	//		binDir = filepath.Join(binDir, zipDir)
	//	}
	//	fileName := "sun-panel"
	//	if zipDir != "" {
	//		fileName = zipDir
	//	}
	//	file.ScanDirectoryAndFunc(binDir, func(fName string) {
	//		if strings.HasPrefix(strings.ToLower(fileName), strings.ToLower(fName)) {
	//			binFilePath := filepath.Join(binDir, fName)
	//			executable, err := os1.IsExecutable(binFilePath)
	//			if err == nil && executable {
	//				binPath = binFilePath
	//			}
	//		}
	//	})
	//}
}
