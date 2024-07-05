package iface

import "github.com/xxl6097/go-service-framework/internal/model"

type IFramework interface {
	AddElement(*model.ProcModel)
	TakeElement() *model.ProcModel
}
