package framework

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/config"
	"github.com/xxl6097/go-service-framework/internal/model"
	"os"
	"path/filepath"
	"time"
)

func (this *Framework) loadConfig() {
	exes, err := config.Get()
	if err == nil && exes != nil {
		for _, v := range exes {
			go this.createProcess(&v)
		}
	}
}

func (this *Framework) createProcess(proc *model.ProcModel) {
	if proc == nil {
		glog.Println("proc is nil")
		return
	}
	rootdir, err1 := os.Getwd()
	if err1 != nil {
		glog.Println(err1)
		return
	}
	bindir := filepath.Join(rootdir, proc.Name)
	// 使用 os.Stat 检查文件夹
	if _, err := os.Stat(bindir); os.IsNotExist(err) {
		// 文件夹不存在
		err = os.MkdirAll(bindir, 0775)
		if err != nil {
			glog.Printf("MkdirAll %s error:%s", bindir, err)
			return
		}
	}

	err := os.Chdir(bindir)
	if err != nil {
		glog.Println("cd error:", err)
		return
	}
}

func (this *Framework) newProcess(bindir, binpath string, proc *model.ProcModel) {
	for {
		// start worker
		tmpDump := filepath.Join(bindir+string(filepath.Separator)+"logs", "dump.log.tmp")
		dumpFile := filepath.Join(bindir+string(filepath.Separator)+"logs", "dump.log")
		f, err2 := os.Create(filepath.Join(tmpDump))
		if err2 != nil {
			glog.Printf("start worker error:%s", err2)
			return
		}
		glog.Println("启动worker进程，参数：", proc.Args)
		execSpec := &os.ProcAttr{Env: append(os.Environ(), "GOTRACEBACK=crash"), Files: []*os.File{os.Stdin, os.Stdout, f}}
		p, err3 := os.StartProcess(binpath, proc.Args, execSpec)
		if err3 != nil {
			glog.Printf("启动worker进程失败，错误信息：%s", err3)
			return
		}
		proc.Proc = p
		_, _ = p.Wait()
		f.Close()
		time.Sleep(time.Second)
		err := os.Rename(tmpDump, dumpFile)
		if err != nil {
			glog.Printf("rename dump error:%s", err)
		}
		if !this.running {
			return
		}
		glog.Printf("worker进程停止,10秒后重新启动")
		time.Sleep(time.Second * 10)
	}
}
