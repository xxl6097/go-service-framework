package framework

import (
	"github.com/kardianos/service"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/config"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/internal/server"
	"github.com/xxl6097/go-service-framework/pkg/version"
	"os"
)

type Framework struct {
	queue      chan *model.ProcModel
	httpserver iface.IHttpServer
	procs      map[string]*model.ProcModel
	running    bool
}

// Shutdown 服务结束回调
func (f *Framework) Shutdown(s service.Service) error {
	defer glog.Flush()
	status, err := s.Status()
	glog.Println("Shutdown")
	glog.Println("Status", status, err)
	glog.Println("Platform", s.Platform())
	glog.Println("String", s.String())
	return nil
}

// Start 服务启动回调
func (f *Framework) Start(s service.Service) error {
	defer glog.Flush()
	status, err := s.Status()
	glog.Println("启动服务")
	glog.Println("Status", status, err)
	glog.Println("Platform", s.Platform())
	glog.Println("String", s.String())
	go f.run()
	return nil
}

// Stop 服务停止回调
func (f *Framework) Stop(s service.Service) error {
	defer glog.Flush()
	glog.Println("停止服务")

	if service.Interactive() {
		glog.Println("停止deamon")
		os.Exit(0)
	}
	return nil
}

func (f *Framework) Config() *service.Config {
	return &service.Config{
		Name:        version.AppName,
		DisplayName: version.DisplayName,
		Description: version.Description,
	}
}

func (f *Framework) Version() string {
	return version.Version()
}

func (f *Framework) AddElement(b *model.ProcModel) {
	// 向通道发送数据，如果通道满了，发送操作会阻塞
	f.queue <- b
	glog.Printf("AddQueue: %v\n", b)
}

func (f *Framework) TakeElement() *model.ProcModel {
	v, ok := <-f.queue
	if ok {
		return v
	}
	return nil
}

func (f *Framework) run() {
	glog.Println("run....")
	f.queue = make(chan *model.ProcModel)
	f.procs = make(map[string]*model.ProcModel)
	f.httpserver = server.NewHttpServer(f)
	f.loadConfig()
	go f.httpserver.Listen()
	for {
		glog.Println("run.for....")
		v, ok := <-f.queue
		if !ok {
			// 当通道被关闭并且所有数据都被消费后，退出循环
			glog.Println("Queue is closed, exiting consumer")
			break
		}
		glog.Printf("Received: %v\n", v)
		f.procs[v.Name] = v
		go f.createProcess(v)
		config.Save(*v)
	}
}
