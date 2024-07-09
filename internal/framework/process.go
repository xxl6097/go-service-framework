package framework

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/config"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/pkg/http"
	"os"
	"path/filepath"
	"time"
)

func (f *Framework) loadConfig() {
	exes, err := config.Get()
	if err == nil && exes != nil {
		glog.Println(exes)
		for _, v := range exes {
			go f.createProcess(&v)
		}
	}
}

func (this *Framework) createProcess(proc *model.ProcModel) {
	if proc == nil {
		glog.Println("proc is nil")
		return
	}
	if proc.BinUrl == "" {
		glog.Error("binurl is nil")
		return
	}
	if proc.Name == "" {
		glog.Error("Name is nil")
		return
	}
	rootDir, err1 := os.Getwd()
	if err1 != nil {
		glog.Println(err1)
		return
	}
	var binIsNotExist bool
	binDir := filepath.Join(rootDir, proc.Name)
	//判断bin文件夹是否存在
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		// 文件夹不存在
		binIsNotExist = true
		err = os.MkdirAll(binDir, 0775)
		if err != nil {
			glog.Printf("MkdirAll %s error:%s", binDir, err)
			return
		}
	}

	if proc.ConfUrl != "" {
		_, confName := filepath.Split(proc.ConfUrl)
		confPath := filepath.Join(binDir, confName)
		var isConfExist bool
		//判断conf文件是否存在
		if _, err := os.Stat(confPath); os.IsNotExist(err) {
			glog.Printf("%s 文件不存在", confPath)
			isConfExist = true
		} else {
			if proc.Upgrade {
				err11 := os.Remove(confPath)
				if err11 != nil {
					glog.Printf("删除文件失败: %s", err11)
				} else {
					glog.Println("文件删除成功")
					isConfExist = true
				}
			}
		}
		if isConfExist {
			err := http.Download(proc.ConfUrl, confPath)
			if err != nil {
				proc.Status = "dowloading failed"
				glog.Error(confPath, proc.Status)
				return
			}
		}
	}

	proc.Status = "creatting"
	_, binName := filepath.Split(proc.BinUrl)
	binPath := filepath.Join(binDir, binName)
	//判断bin文件是否存在
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		glog.Printf("%s 文件夹不存在", binPath)
		binIsNotExist = true
	} else {
		if proc.Upgrade {
			err11 := os.Remove(binPath)
			if err11 != nil {
				glog.Printf("删除文件失败: %s", err11)
			} else {
				glog.Println("文件删除成功")
				binIsNotExist = true
			}
		}
	}

	if binIsNotExist {
		proc.Status = "dowloading"
		err := http.Download(proc.BinUrl, binPath)
		if err != nil {
			proc.Status = "dowloading failed"
			glog.Error(proc.Status)
			return
		}
		glog.Info("download sucess.", binPath, err)
	}

	logDir := filepath.Join(binDir, "logs")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		// 文件夹不存在
		err = os.MkdirAll(logDir, 0775)
		if err != nil {
			glog.Printf("MkdirAll %s error:%s", logDir, err)
			return
		}
	}
	proc.Upgrade = false
	this.procs[proc.Name] = proc
	this.startProcess(binDir, binPath, logDir, proc)
}

func (this *Framework) startProcess(bindir, binPath, logDir string, proc *model.ProcModel) {
	err := os.Chdir(bindir)
	glog.Debug("chdir", bindir, err)
	for {
		tmpDump := filepath.Join(logDir, "dump.log.tmp")
		dumpFile := filepath.Join(logDir, "dump.log")
		f, err2 := os.Create(filepath.Join(tmpDump))
		if err2 != nil {
			glog.Printf("start worker error:%s", err2)
			return
		}
		defer f.Close()
		err1 := os.Chmod(binPath, 0755)
		if err1 == nil {
			glog.Debug(binPath, "赋予0755权限成功")
		} else {
			glog.Error(binPath, "赋予0755权限失败")
		}

		var args = []string{binPath}
		if proc.Args != nil && len(proc.Args) > 0 {
			args = append(args, proc.Args...)
		}
		glog.Println("启动worker进程", binPath, args)
		execSpec := &os.ProcAttr{
			Env:   append(os.Environ(), "GOTRACEBACK=crash"),
			Files: []*os.File{os.Stdin, os.Stdout, f},
			//Sys: &syscall.SysProcAttr{
			//	Chroot: bindir,
			//},
		}
		p, err3 := os.StartProcess(binPath, args, execSpec)
		if err3 != nil {
			glog.Printf("启动worker进程失败，错误信息：%s", err3)
			return
		}
		proc.Status = "running"
		proc.Proc = p
		config.Save(*proc)
		status, err4 := p.Wait()
		if err4 == nil {
			glog.Debug("成功", status.String(), err4)
		} else {
			glog.Error("失败", status.String(), err4)
		}
		proc.Status = "stopped"
		time.Sleep(time.Second)
		err := os.Rename(tmpDump, dumpFile)
		glog.Debugf("rename dump %v", err)
		glog.Debugf("this.running:%v,proc.Exit:%v", this.running, proc.Exit)
		if !this.running {
			return
		}
		if proc.Exit {
			return
		}
		glog.Info("worker进程停止,10秒后重新启动")
		time.Sleep(time.Second * 10)
	}
}
