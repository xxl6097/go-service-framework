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
	//glog.SetLogFile("./", "install.log")
	gservice.Run(&framework.Framework{})
}

func test() {
	//ctx, cancel := context.WithCancel(context.Background())
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	cancel()
	//}()
	//timer.Countdown(10, ctx, func() {
	//	glog.Error("done")
	//}, func(s string) {
	//	glog.Debug(s)
	//})

}
