package server

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	http2 "github.com/xxl6097/go-service-framework/pkg/http"
	"log"
	"net/http"
)

type httpserver struct {
	iframework iface.IFramework
}

func NewHttpServer(framework iface.IFramework) iface.IHttpServer {
	this := httpserver{
		iframework: framework,
	}
	return &this
}

func (h *httpserver) Listen() {
	http.HandleFunc("/api", http2.BasicAuth(h.helloHandler, "admin", "het002402"))
	// 启动HTTP服务器，监听8088端口
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

// 定义一个简单的处理函数
func (h *httpserver) helloHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	glog.Info(url)
	if url != "" {
		proc := model.ProcModel{
			Name:    "test1",
			BinUrl:  "http://uuxia.cn:8086/files/2024/07/03/AAServiceTest_0.0.1_windows_amd64.exe",
			ConfUrl: "http://uuxia.cn:8086/files/2024/07/03/AAServiceTest_0.0.1_windows_amd64.exe",
			Args:    []string{"-c", "frpc.toml"},
		}
		h.iframework.AddElement(&proc)
		fmt.Fprintf(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "no url")
	}

}
