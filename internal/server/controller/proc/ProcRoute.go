package proc

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/server/inter"
	"github.com/xxl6097/go-http/server/route"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"net/http"
)

type ProcRoute struct {
	controller *ProcController
}

func (this *ProcRoute) Setup(router *mux.Router) {
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/proc/getall",
		Fun:    this.controller.getall,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/proc/new",
		Fun:    this.controller.new,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/proc/del",
		Fun:    this.controller.del,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/proc/start",
		Fun:    this.controller.start,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/proc/stop",
		Fun:    this.controller.stop,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/proc/restart",
		Fun:    this.controller.restart,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/auth",
		Fun:    this.controller.auth,
		NoAuth: true,
	})
}

func NewRoute(iframework iface.IFramework) inter.IRoute {
	opt := &ProcRoute{
		controller: NewController(iframework),
	}
	return opt
}
