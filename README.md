
go get github.com/kardianos/service

go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build-trimpath -ldflags "-linkmode internal" -o AAServiceApp.exe main.go

go get -u github.com/xxl6097/go-glog@v0.0.17

go get -u github.com/xxl6097/go-service@v0.0.27

go get -u github.com/xxl6097/go-http@v0.0.10

goversioninfo -manifest versioninfo.json


```azure

package main

import (
	"fmt"
	"os"
)

func main() {
	os.Chdir("/Users/uuxia/Library/Caches/JetBrains/GoLand2024.1/tmp/GoLand/frpc/")
	executable := "/Users/uuxia/Library/Caches/JetBrains/GoLand2024.1/tmp/GoLand/frpc/frpc" // 可执行文件的路径
	args := []string{executable, "-c", "frpc.toml"}                                         // 启动参数列表

	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		//Sys: &syscall.SysProcAttr{
		//	Chroot: "/Users/uuxia/Library/Caches/JetBrains/GoLand2024.1/tmp/GoLand/frpc/",
		//},
	}

	p, err := os.StartProcess(executable, args, attr)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	p.Wait()

	fmt.Println("Process started successfully.")
}

```

## windows打包

1、 main.go 文件中添加标签，如下

```
//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
}
```

2、编译打包

```
go generate
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-linkmode internal $ldflags" -o ${appname}_${version}_windows_amd64.exe
```

3、生成版本信息
resource文件夹