package main

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/framework"
	"github.com/xxl6097/go-service/svr"
)

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
	glog.Println("hello")
	svr.Run(&framework.Framework{})
}
