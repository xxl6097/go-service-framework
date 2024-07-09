package iface

import "github.com/xxl6097/go-service-framework/internal/model"

type IFramework interface {
	AddElement(*model.ProcModel)
	TakeElement() *model.ProcModel
	GetAll() []model.ProcModel
	Delete(string) error
	StartProcess(string) error
	StopProcess(string) error
	RestartProcess(string) error
}
