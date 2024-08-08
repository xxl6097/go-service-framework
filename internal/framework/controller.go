package framework

import (
	"errors"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/pkg/file"
	"github.com/xxl6097/go-service-framework/pkg/util"
	"os"
	"path/filepath"
	"time"
)

func (f *Framework) SetAppStore(url string) {
	if url == "" || f.cache == nil {
		return
	}
	config := f.cache.Get()
	if config == nil {
		return
	}
	config.AppStoreUrl = url
	glog.Debug("appstore", url)
	f.cache.Set(config)
}
func (f *Framework) GetConfig() *model.ConfigModel {
	if f.cache == nil {
		return nil
	}
	return f.cache.Read()
}

func (f *Framework) SetConfig(config *model.ConfigModel) error {
	if f.cache == nil {
		return fmt.Errorf("cache is nil")
	}
	if config == nil {
		return fmt.Errorf("config is nil")
	}
	return f.cache.Set(config)
}

func (f *Framework) GetAll() []model.ProcModel {
	procs := make([]model.ProcModel, 0)
	for _, v := range f.procs {
		procs = append(procs, *v)
	}
	return procs
}

func (f *Framework) StopProcess(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	v, exist := f.procs[name]
	if exist {
		if v.Proc != nil {
			err := util.Kill(v.Proc)
			if err == nil {
				return nil
			}
			glog.Error(err)
			return err
		} else {
			return errors.New(fmt.Sprintf("%s proc is nil", name))
		}
	}
	return errors.New(fmt.Sprintf("%s proc is not exist", name))
}

func (f *Framework) StartProcess(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	v, exist := f.procs[name]
	if exist {
		if v.Proc != nil {
			return nil
		}
		v.Exit = model.STOP_NO
		go f.createProcess(v)
		return nil
	}
	return errors.New(fmt.Sprintf("%s proc is not exist", name))
}

func (f *Framework) RestartProcess(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	v, exist := f.procs[name]
	if exist {
		v.Exit = model.STOP_EXIT
		if v.Proc != nil {
			err := util.Kill(v.Proc)
			if err != nil {
				glog.Errorf("%s proc kill error: %s", name, err.Error())
			}
		}

		time.AfterFunc(time.Second*3, func() {
			v.Exit = model.STOP_NO
			go f.createProcess(v)
		})
		return nil
	}
	return errors.New(fmt.Sprintf("%s proc is not exist", name))
}

// Delete 卸载程序
func (f *Framework) Delete(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	v, exist := f.procs[name]
	if exist {
		v.Exit = model.STOP_DELETE
		if v.Proc != nil {
			err := util.Kill(v.Proc)
			if err == nil {
				delete(f.procs, name)
				glog.Debug("delete proc success", f.procs)
				return nil
			}
			if v.Cancel != nil {
				v.Cancel()
			}
			glog.Error(err)
			return err
		} else {
			//TODO 也要删除源文件
			f.deleteApplication(name)
			return errors.New(fmt.Sprintf("%s proc is nil", name))
		}
	}
	return errors.New(fmt.Sprintf("%s proc is not exist", name))
}

func (f *Framework) GetPassCode() string {
	return f.passcode
}

func (f *Framework) GetAppConfig(appName string) []byte {
	if appName == "" {
		return nil
	}
	p, exists := f.procs[appName]
	if exists && p != nil && p.ConfUrl != "" {
		//p.Args
		exeDir, err := os.Getwd()
		if err == nil && exeDir != "" {
			_, confFileName := filepath.Split(p.ConfUrl)
			if confFileName != "" {
				confPath := filepath.Join(exeDir, appName, confFileName)
				return file.ReadFile(confPath)
			}
			glog.Error(exeDir, confFileName, err)
		}

	}
	return nil
}

func (f *Framework) SaveAppConfig(appName string, body []byte) error {
	if body == nil {
		return errors.New("body is nil")
	}
	p, exists := f.procs[appName]
	if exists && p != nil && p.ConfUrl != "" {
		//p.Args
		exeDir, err := os.Getwd()
		if err == nil && exeDir != "" {
			_, confFileName := filepath.Split(p.ConfUrl)
			if confFileName != "" {
				confPath := filepath.Join(exeDir, appName, confFileName)
				return file.SaveFile(confPath, body)
			}
			glog.Error(exeDir, confFileName, err)
		}

	}
	return nil
}

func (f *Framework) GetLogPath(name string) string {
	dir, err := os.Getwd()
	if err == nil && dir != "" {
		return filepath.Join(dir, name, "logs", "dump.tmp.log")
	}
	return ""
}
