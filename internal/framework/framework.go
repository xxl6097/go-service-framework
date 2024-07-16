package framework

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/internal/repository"
	"github.com/xxl6097/go-service-framework/internal/server"
	"github.com/xxl6097/go-service-framework/pkg/crypt"
	os2 "github.com/xxl6097/go-service-framework/pkg/os"
	"github.com/xxl6097/go-service-framework/pkg/version"
	"github.com/xxl6097/go-sqlite/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Framework struct {
	queue    chan *model.ProcModel
	procs    map[string]*model.ProcModel
	running  bool
	db       *gorm.DB
	confRepo iface.ISqlite[model.ConfigModel]
	procRepo iface.ISqlite[model.ProcModel]
	passcode string
}

func (f *Framework) inputArgs() []string {
	for {
		var port int
		fmt.Print("设置服务端口，请输入：")
		fmt.Scan(&port)
		if port < 0 {
			fmt.Println("端口输入错误，请重新输入！")
		} else {
			return []string{strconv.Itoa(port), "by", "uuxia"}
		}
	}
}

func (f *Framework) getFirstConfigCache() *model.ConfigModel {
	data, err := f.confRepo.First()
	if err != nil || data == nil {
		return nil
	}
	return data
}
func (f *Framework) inputAuthCode(installPath string) ([]byte, string) {
	for {
		var password string
		fmt.Print("设置授权密码，请输入：")
		fmt.Scan(&password)
		password = strings.TrimSpace(password)
		passcode, err := crypt.SavePassword(installPath, []byte(password))
		if err != nil {
			fmt.Println("授权码设置失败，请重新设置！")
		} else {
			fmt.Println("授权码设置成功", installPath)
			f.passcode = string(passcode)
			return passcode, password
		}
	}
}

func (f *Framework) initData(args []string, pass string) {
	err := f.confRepo.Add(&model.ConfigModel{
		Password: pass,
		Args:     strings.Join(args, ","),
	})
	glog.Error(err)
}

func (f *Framework) hasConfig(installPath string) []string {
	if f.db == nil {
		f.db = sqlite.InitMysql(filepath.Join(installPath, "sqlite.db"))
		f.confRepo = repository.NewConfRepository(f.db)
		f.procRepo = repository.NewProcRepository(f.db)
	}
	data := f.getFirstConfigCache()
	if data != nil {
		return strings.Split(data.Args, ",")
	}
	return nil
}

func (f *Framework) OnInstall(installPath string) []string {
	data := f.hasConfig(installPath)
	if data != nil {
		return data
	}
	args := f.inputArgs()
	hashcode, _ := f.inputAuthCode(installPath)
	f.initData(args, string(hashcode))
	return args
}

// Shutdown 服务结束回调
func (f *Framework) Shutdown(s service.Service) error {
	defer glog.Flush()
	f.running = false
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
	f.running = true
	go f.run()
	return nil
}

// Stop 服务停止回调
func (f *Framework) Stop(s service.Service) error {
	defer glog.Flush()
	glog.Println("停止服务")
	f.running = false
	for k, v := range f.procs {
		if v.Proc != nil {
			glog.Debug("停止worker进程", k)
			err := os2.Kill(v.Proc)
			glog.Debugf("kill %s %v", k, err)
		}
	}
	time.Sleep(1 * time.Second)
	if service.Interactive() {
		glog.Println("停止deamon")
		os.Exit(0)
	}
	return nil
}

func (f *Framework) Config() *service.Config {
	if os2.IsMacOs() {
		version.AppName = "AAFrameWork"
		version.DisplayName = "AAFrameWork"
		version.Description = "A Test AAFrameWork"
	}
	return &service.Config{
		Name:        version.AppName,
		DisplayName: version.DisplayName,
		Description: version.Description,
	}
}

func (f *Framework) Version() string {
	return version.Version()
}

func (f *Framework) AddElement(v *model.ProcModel) {
	// 向通道发送数据，如果通道满了，发送操作会阻塞
	p, exists := f.procs[v.Name]
	if exists {
		if v.Upgrade {
			p.Exit = model.STOP_EXIT
			if p.Proc != nil {
				glog.Debugf("%s 停止worker进程", p.Name)
				err := os2.Kill(p.Proc)
				glog.Debugf("kill %s %v", p.Name, err)
			}
		} else {
			glog.Error("程序已经存在", v.Name)
			return
		}
	}

	f.queue <- v
	glog.Printf("AddQueue: %v\n", v)
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
	f.loadConfig()
	port := 8888
	if len(os.Args) > 2 {
		pstr := os.Args[2]
		if pstr != "" {
			_port, err := strconv.Atoi(pstr)
			if err == nil && _port > 0 {
				port = _port
			}
		}
	}

	go server.Listen(port, f)
	for {
		glog.Println("run.for....")
		v, ok := <-f.queue
		if !ok {
			// 当通道被关闭并且所有数据都被消费后，退出循环
			glog.Println("Queue is closed, exiting consumer")
			break
		}
		glog.Printf("Received: %v", v)
		go f.createProcess(v)
	}
}
