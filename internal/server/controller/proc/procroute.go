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
		Method: http.MethodPost,
		Path:   "/proc/app/market",
		Fun:    this.controller.appMarket,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/proc/app/list",
		Fun:    this.controller.getAppList,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/proc/info",
		Fun:    this.controller.info,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/proc/getall",
		Fun:    this.controller.getall,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/setting/appstore",
		Fun:    this.controller.settingAppStore,
		NoAuth: false,
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
		Method: http.MethodGet,
		Path:   "/proc/read/log",
		Fun:    this.controller.readLog,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/proc/app/config",
		Fun:    this.controller.appConfig,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/proc/app/config/save",
		Fun:    this.controller.appConfigSave,
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
		Path:   "/login",
		Fun:    this.controller.login,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/auth",
		Fun:    this.controller.auth,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/uninstall",
		Fun:    this.controller.uninstall,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/reboot",
		Fun:    this.controller.reboot,
		NoAuth: false,
	})
}

func NewRoute(iframework iface.IFramework) inter.IRoute {
	opt := &ProcRoute{
		controller: NewController(iframework),
	}
	return opt
}
