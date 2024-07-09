package framework

import (
	"errors"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/model"
)

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
			err := v.Proc.Kill()
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
		if v.Proc != nil && v.Status == "running" {
			return nil
		}
		v.Exit = false
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
		if v.Proc != nil {
			err := v.Proc.Kill()
			glog.Error(err)
		}
		v.Exit = false
		go f.createProcess(v)
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
		v.Exit = true
		if v.Proc != nil {
			err := v.Proc.Kill()
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
