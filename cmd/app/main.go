package main

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/framework"
	"github.com/xxl6097/go-service/gservice"
)

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
	test()
	glog.SetLogFile("./", "install.log")
	gservice.Run(&framework.Framework{})
}

func test() {
	//glog.Debug("倒计时开始...")
}
