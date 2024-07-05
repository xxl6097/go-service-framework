package framework

import (
	"github.com/kardianos/service"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/pkg/version"
	"os"
	"time"
)

type Framework struct{}

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

func (f *Framework) run() {
	for {
		glog.Info("Running...")
		time.Sleep(time.Second * 5)
	}
}
