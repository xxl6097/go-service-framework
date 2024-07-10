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
	//var password string
	//fmt.Print("设置授权码,请输入:")
	//fmt.Scan(&password)
	//fullexecpath, _ := os.Executable()
	//installPath, _ := filepath.Split(fullexecpath)
	//err := crypt.SavePassword(installPath, []byte(password))
	//if err != nil {
	//	fmt.Println("授权码设置失败，请重新设置！")
	//} else {
	//	fmt.Println("授权码设置成功", installPath)
	//	return
	//}
}
