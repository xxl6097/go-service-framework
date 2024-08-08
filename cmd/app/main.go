package main

import (
	"github.com/xxl6097/go-service-framework/internal/framework"
	"github.com/xxl6097/go-service-framework/pkg/util"
	"github.com/xxl6097/go-service/gservice"
	"os"
)

func init() {
	if util.IsMacOs() {
		//os1.SetDebug(true)
	}
	if os.Getenv("DEBUG") != "" {
		util.SetDebug(true)
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
