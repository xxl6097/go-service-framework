package iface

import "github.com/xxl6097/go-service-framework/internal/model"

type IFramework interface {
	AddElement(*model.ProcModel)
	TakeElement() *model.ProcModel
	GetAll() []model.ProcModel
	GetConfig() *model.ConfigModel
	SetAppStore(string)
	Delete(string) error
	StartProcess(string) error
	StopProcess(string) error
	RestartProcess(string) error
	GetPassCode() string
	GetAppConfig(string) []byte
	SaveAppConfig(string, []byte) error
	GetLogPath(string) string
}
