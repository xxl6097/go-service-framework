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
	os1.SetDebug(true)
}

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
	test()
	//glog.SetLogFile("./", "install.log")
	gservice.Run(&framework.Framework{})
}

func test() {
	//jsonstr := "{\"windows\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"wechat\",\"args\":[\"-d\",\"conf.toml\"],\"description\":\"微信应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"QQ\",\"args\":[\"-d\",\"qq.toml\"],\"description\":\"QQ应用程序，用于测试\"}]},\"linux\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"dingtalk\",\"args\":[\"-d\",\"dingtalk.toml\"],\"description\":\"dingtalk应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]},\"darwin\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"dingtalk\",\"args\":[\"-d\",\"dingtalk.toml\"],\"description\":\"dingtalk应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]}}"
	//maps := jsonutil.JsonStrToMap(jsonstr)
	//for k, v := range maps {
	//	if strings.Compare(k, runtime.GOOS) == 0 {
	//		if s, ok := v.(map[string]interface{}); ok {
	//			fmt.Println("Interface value is a string:", s, s[runtime.GOARCH])
	//		} else {
	//			fmt.Println("Interface value is not a string")
	//		}
	//		return
	//	}
	//}
	//glog.Debug(maps)

	//base64.StdEncoding.DecodeString(auth[len(basicScheme):])
	//str := base64.StdEncoding.EncodeToString([]byte("admin:admin"))
	//fmt.Println(str)
	//
	//path := "path-admin:het002402-/data/z4"
	//paths := strings.Split(path, "-")
	//if paths != nil && len(paths) >= 3 {
	//	fmt.Println(paths)
	//}

}
