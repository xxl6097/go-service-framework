package main

import (
	"github.com/xxl6097/go-service-framework/internal/framework"
	os1 "github.com/xxl6097/go-service-framework/pkg/os"
	"github.com/xxl6097/go-service/gservice"
	"os"
)

func init() {
	if os1.IsMacOs() {
		//os1.SetDebug(true)
	}
	if os.Getenv("DEBUG") != "" {
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
	//binFilePath := "/usr/local/AuGoService/vnt/vnt-cli"
	//executable, err := os1.IsExecutable(binFilePath)
	//glog.Debug(executable, err)
}
