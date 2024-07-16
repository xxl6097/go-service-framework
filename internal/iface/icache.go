package iface

import "github.com/xxl6097/go-service-framework/internal/model"

type ICache interface {
	HasCache() bool
	Get() *model.ConfigModel
	Set(*model.ConfigModel) error
	Save(data *model.ProcModel) error
	Delete(name string) error
}
