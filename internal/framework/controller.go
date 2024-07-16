package framework

import (
	"errors"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/pkg/os"
	"time"
)

func (f *Framework) GetConfig() *model.ConfigModel {
	if f.cache == nil {
		return nil
	}
	return f.cache.Get()
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
			err := os.Kill(v.Proc)
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
			err := os.Kill(v.Proc)
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

func (f *Framework) Delete(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	v, exist := f.procs[name]
	if exist {
		v.Exit = model.STOP_DELETE
		if v.Proc != nil {
			err := os.Kill(v.Proc)
			if err == nil {
				delete(f.procs, name)
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

func (f *Framework) GetPassCode() string {
	return f.passcode
}
