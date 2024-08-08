package framework

import (
	"context"
	"errors"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/cache"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/pkg/file"
	"github.com/xxl6097/go-service-framework/pkg/http"
	"github.com/xxl6097/go-service-framework/pkg/java"
	os2 "github.com/xxl6097/go-service-framework/pkg/os"
	"github.com/xxl6097/go-service-framework/pkg/timer"
	"github.com/xxl6097/go-service-framework/pkg/zip"
	"os"
	"path/filepath"
	"strings"
)

func (f *Framework) loadConfig() {
	bindir, _ := os.Getwd()
	if bindir != "" {
		//if f.db == nil {
		//	f.db = sqlite.InitMysql(filepath.Join(bindir, "sqlite.db"))
		//	f.confRepo = repository.NewConfRepository(f.db)
		//	f.procRepo = repository.NewProcRepository(f.db)
		//	conf, e := f.confRepo.First()
		//	if e == nil {
		//		f.passcode = conf.Password
		//		glog.Debug(conf)
		//	}
		//	if os2.IsMacOs() {
		//		f.OnInstall(bindir)
		//	}
		//	datas, e := f.procRepo.FindAll()
		//	if e == nil && datas != nil && len(datas) > 0 {
		//		for _, v := range datas {
		//			go f.createProcess(&v)
		//		}
		//	}
		//}
		if f.cache == nil {
			f.cache = cache.NewCache(bindir)
		}
		config := f.cache.Get()
		if config == nil {
			config = &model.ConfigModel{}
		} else {
			f.passcode = config.Password
		}
		if os2.IsDebug() {
			f.OnInstall(bindir)
		}
		datas := config.Procs
		if datas != nil && len(datas) > 0 {
			for _, v := range datas {
				go f.createProcess(&v)
			}
		}
	}
	//exes, err := config.Get()
	//if err == nil && exes != nil {
	//	glog.Println(exes)
	//	for _, v := range exes {
	//		go f.createProcess(&v)
	//	}
	//}
}

// checkConfigFile 检查config配置文件
func (this *Framework) checkConfigFile(binDir string, proc *model.ProcModel) (string, error) {
	configUrl := proc.ConfUrl
	if configUrl == "" {
		glog.Debug(proc.Name, "无配置文件.")
		return "", nil
	}
	if !file.IsUrlOrLocalFile(configUrl) {
		errmsg := fmt.Sprintf("无效配置文件地址 %s", configUrl)
		glog.Error(errmsg)
		return "", errors.New(errmsg)
	}
	_, name := filepath.Split(configUrl)
	confPath := filepath.Join(binDir, name)
	configNeedDownload := false
	//判断conf文件是否存在
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		glog.Printf("%s 配置文件不存在", confPath)
		configNeedDownload = true
	} else {
		if proc.Upgrade {
			err1 := os.Remove(confPath)
			if err1 != nil {
				glog.Printf("配置文件删除失败: %s", err1)
			} else {
				glog.Println("配置文件删除成功")
				configNeedDownload = true
			}
		}
	}
	if configNeedDownload {
		err := http.Download(configUrl, confPath)
		if err != nil {
			proc.Status = "配置文件下载失败"
			errMsg := fmt.Sprintf("%s%s", proc.Status, confPath)
			glog.Error(errMsg)
			return "", errors.New(errMsg)
		}
	}
	return confPath, nil
}

// checkBinFile 检查可执行程序是否存在
func (this *Framework) checkBinFile(binDir string, proc *model.ProcModel) (string, error) {
	binUrl := proc.BinUrl
	if binUrl == "" {
		errMsg := "bin可执行文件地址不正确"
		glog.Debug(errMsg)
		return "", errors.New(errMsg)
	}
	if !file.IsUrlOrLocalFile(binUrl) {
		errMsg := fmt.Sprintf("可执行文件文件地址无效 %s", binUrl)
		glog.Error(errMsg)
		return "", errors.New(errMsg)
	}

	binNotExist := false
	proc.Status = "创建中"
	//_, binName := filepath.Split(binUrl)
	//binName := proc.Name

	_, binName := filepath.Split(proc.BinUrl)
	binPath := filepath.Join(binDir, binName)
	//判断bin文件是否存在
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		glog.Printf("程序未下载 %s", binPath)
		binNotExist = true
	} else {
		glog.Printf("程序下载过 %s", binPath)
		if proc.Upgrade {
			err1 := os.Remove(binPath)
			if err1 != nil {
				glog.Printf("可执行文件删除失败: %s", err1)
			} else {
				glog.Println("可执行文件删除成功")
				binNotExist = true
			}
		}
	}

	glog.Debug("binPath", binPath)
	if binNotExist {
		proc.Status = "下载中"
		glog.Debug("下载中", binUrl)
		err := http.Download(binUrl, binPath)
		if err != nil {
			proc.Status = "可执行文件下载失败"
			eMsg := fmt.Sprintf("%s%s", proc.Status, binPath)
			glog.Error(eMsg)
			return "", errors.New(eMsg)
		} else {
			glog.Debug("下载完成，解压", binDir)
			isZip, err := zip.UnPack(binPath, binDir)
			if err == nil && isZip {
				//确定解压成功
				zipDir, err := zip.GetRootDir(binDir, proc.Name)
				if err == nil && zipDir != "" {
					//确定zip有一级目录
					binDir = filepath.Join(binDir, zipDir)
				}
				fileName := proc.Name
				if zipDir != "" {
					fileName = zipDir
				}
				glog.Debugf("解压成功,zipDir:%s,binDir:%s,fileName:%s,err:%v", zipDir, binDir, fileName, err)
				file.ScanDirectoryAndFunc(binDir, func(fName string) {
					isZip = zip.IsZip(fName)
					hasprefix := strings.HasPrefix(strings.ToLower(fileName), strings.ToLower(fName))
					glog.Debugf("扫描目:%s,zip:%v,是否匹配[%s]:%v", fName, isZip, fileName, hasprefix)
					if hasprefix && !isZip {
						binFilePath := filepath.Join(binDir, fName)
						executable, err := os2.IsExecutable(binFilePath)
						glog.Debugf("检测可执行程序:err:%v executable:%v binFilePath:%s", binFilePath)
						if err == nil && executable {
							binPath = binFilePath
						}
					}
				})
			}
		}
	} else {
		isZip := zip.IsZip(binPath)
		if isZip {
			glog.Debug("判断程序是压缩包", binPath)
			//确定解压成功
			zipDir, err := zip.GetRootDir(binDir, proc.Name)
			if err == nil && zipDir != "" {
				//确定zip有一级目录
				binDir = filepath.Join(binDir, zipDir)
			}
			fileName := proc.Name
			if zipDir != "" {
				fileName = zipDir
			}
			glog.Debugf("解压成功,zipDir:%s,binDir:%s,fileName:%s,err:%v", zipDir, binDir, fileName, err)
			file.ScanDirectoryAndFunc(binDir, func(fName string) {
				isZip = zip.IsZip(fName)
				hasprefix := strings.HasPrefix(strings.ToLower(fileName), strings.ToLower(fName))
				glog.Debugf("扫描目:%s,zip:%v,是否匹配[%s]:%v", fName, isZip, fileName, hasprefix)
				if hasprefix && !isZip {
					binFilePath := filepath.Join(binDir, fName)
					executable, err := os2.IsExecutable(binFilePath)
					glog.Debugf("检测可执行程序:err:%v executable:%v binFilePath:%s", binFilePath)
					if err == nil && executable {
						binPath = binFilePath
					}
				}
			})
		}
	}
	glog.Info("可执行文件路径", binPath)
	return binPath, nil
}

// checkBinDir 检查bin程序目录是否存在
func (this *Framework) checkBinDir(proc *model.ProcModel) (string, error) {
	if proc.Name == "" {
		errmsg := "bin Name is nil"
		glog.Error(errmsg)
		return "", errors.New(errmsg)
	}
	exePath, err := os.Executable()
	if err != nil {
		glog.Error(err)
		return "", err
	}
	rootDir, _ := filepath.Split(exePath)
	//rootDir, err := os.Getwd()
	//if err != nil {
	//	glog.Error(err)
	//	return "", err
	//}
	binDir := filepath.Join(rootDir, proc.Name)
	//判断bin文件夹是否存在
	if _, err1 := os.Stat(binDir); os.IsNotExist(err1) {
		// 文件夹不存在
		err2 := os.MkdirAll(binDir, 0775)
		if err2 != nil {
			glog.Printf("MkdirAll %s error:%s", binDir, err2)
			return binDir, err2
		}
	}
	return binDir, nil
}

// checkLogDir 检查log目录是否存在
func (this *Framework) checkLogDir(binDir string) (string, error) {
	logDir := filepath.Join(binDir, "logs")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		// 文件夹不存在
		err = os.MkdirAll(logDir, 0775)
		if err != nil {
			eMsg := fmt.Sprintf("MkdirAll %s%s", logDir, err.Error())
			glog.Error(eMsg)
			return "", errors.New(eMsg)
		}
	}
	return logDir, nil
}

// 1. 检查bin目录是否存在，不存在：创建并记录，存在：记录
// 2. 检查config配置文件是否存在，不存在下载，下载失败返回错误；
// 3. 检查bin是不是java，java：走java流程，可执行程序：判断是否存在，不存在down
func (this *Framework) createProcess(proc *model.ProcModel) {
	if proc == nil {
		glog.Errorf("程序运行失败 proc is nil")
		return
	}

	binDir, err := this.checkBinDir(proc)
	if err != nil {
		glog.Errorf("【%s】程序运行失败 %v", proc.Name, err)
		return
	}
	_, err1 := this.checkConfigFile(binDir, proc)
	if err1 != nil {
		glog.Errorf("【%s】程序运行失败 %v", proc.Name, err1)
		return
	}

	binPath, err2 := this.checkBinFile(binDir, proc)
	if err2 != nil {
		glog.Errorf("【%s】程序运行失败 %v", proc.Name, err2)
		return
	}

	logDir, err3 := this.checkLogDir(binDir)
	if err3 != nil {
		glog.Errorf("【%s】程序运行失败 %v", proc.Name, err3)
		return
	}

	proc.Upgrade = false
	this.procs[proc.Name] = proc
	this.startProcess(binDir, binPath, logDir, proc)
}

func (this *Framework) startProcess(binDir, binPath, logDir string, proc *model.ProcModel) {
	var args = []string{binPath}
	if java.IsJar(binPath) {
		javaPath, err2 := java.FindJavaPath()
		if err2 != nil {
			glog.Error("您的电脑没有Java运行环境，请安装JDK")
			return
		}
		_, jar := filepath.Split(binPath)
		binPath = javaPath
		args = []string{binPath, "-jar", jar}
	}

	if proc.Args != nil && len(proc.Args) > 0 {
		args = append(args, proc.Args...)
	}
	//glog.Debug("===>", binPath, args)
	err1 := os.Chmod(binPath, 0755)
	if err1 == nil {
		glog.Debug(binPath, "赋予0755权限成功")
	} else {
		glog.Error(binPath, "赋予0755权限失败")
	}
	//outFile := filepath.Join(logDir, "out.log")
	//out, err2 := os.Create(filepath.Join(outFile))
	//if err2 != nil {
	//	glog.Errorf("【%s】程序运行失败 %v", proc.Name, err2)
	//	return
	//}
	//defer out.Close()

	this.wg.Add(1)
	defer this.wg.Done()
	for {
		tmpDump := filepath.Join(logDir, "dump.tmp.log")
		dumpFile := filepath.Join(logDir, "dump.log")
		stderr, err21 := os.Create(filepath.Join(tmpDump))
		if err21 != nil {
			glog.Errorf("【%s】程序运行失败 %v", proc.Name, err21)
			return
		}
		defer stderr.Close()
		glog.Println("启动worker进程", binPath, args)
		execSpec := &os.ProcAttr{
			Dir: binDir,
			Env: append(os.Environ(), "GOTRACEBACK=crash"),
			//Files: []*os.File{os.Stdin, os.Stdout, f},
			//Files: []*os.File{os.Stdin, out, stderr},
			Files: []*os.File{os.Stdin, stderr, stderr},
			//Sys: &syscall.SysProcAttr{
			//	Chroot: binDir,
			//},
		}
		p, err3 := os.StartProcess(binPath, args, execSpec)
		proc.Proc = p
		if err3 != nil {
			errMsg := fmt.Sprintf("启动worker进程失败，错误信息：%s", err3)
			glog.Printf(errMsg)
			proc.Status = errMsg
			return
		} else {
			proc.Status = "运行中"
			//config.Save(*proc)
			//this.procRepo.Save(proc)
			this.cache.Save(proc)
		}
		glog.Debugf("【程序启动成功】%s", proc.Name)
		status, err4 := p.Wait()
		if err4 == nil {
			glog.Debugf("Wait正常停止 %s", proc.Name)
		} else {
			glog.Errorf("Wait异常停止 %v %s %v", err4, binPath, status.String())
		}
		proc.Status = "已停止"
		//err5 := p.Release()
		//glog.Debugf("【%s】释放资源 %v", proc.Name, err5)
		//time.Sleep(time.Second)
		os.Rename(tmpDump, dumpFile)
		if !this.running {
			glog.Error("进程结束", this.running)
			return
		}
		glog.Warnf("【%s】进程停止,10秒后重新启动", proc.Name)
		//for i := 10; i > 0; i-- {
		//	proc.Status = fmt.Sprintf("【%s】%d秒后重新启动..", proc.Name, i)
		//	fmt.Printf("\r%s", proc.Status)
		//	time.Sleep(1 * time.Second)
		//}

		if proc.Exit == model.STOP_EXIT {
			glog.Error("进程结束", proc.Name)
			return
		} else if proc.Exit == model.STOP_DELETE {
			//err5 := os.RemoveAll(binDir)
			//if err5 != nil {
			//	glog.Error("进程结束，删除程序失败", binDir)
			//} else {
			//	glog.Debug("进程结束，程序删除成功", binDir)
			//}
			////config.Delete(proc.Name)
			////this.procRepo.Delete(proc)
			//delete(this.procs, proc.Name)
			//this.cache.Delete(proc.Name)
			this.deleteApplication(proc.Name)
			return
		} else {
			timer.Countdown(10, func(ctx context.Context, cancel context.CancelFunc) {
				proc.Context = ctx
				proc.Cancel = cancel
			}, func(i int) {
				proc.Status = fmt.Sprintf("【%s】%d秒后重新启动..", proc.Name, i)
				fmt.Printf("\r%s", proc.Status)
			})
		}
	}
}

func (this *Framework) deleteApplication(name string) {
	baseDir, err := os.Getwd()
	if err != nil {
		return
	}
	binDir := filepath.Join(baseDir, name)
	err5 := os.RemoveAll(binDir)
	if err5 != nil {
		glog.Error("进程结束，删除程序失败", binDir)
	} else {
		glog.Debug("进程结束，程序删除成功", binDir)
	}
	//config.Delete(proc.Name)
	//this.procRepo.Delete(proc)
	delete(this.procs, name)
	this.cache.Delete(name)
}
