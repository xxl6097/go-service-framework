package server

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/api"
	"github.com/xxl6097/go-http/server"
	"github.com/xxl6097/go-http/server/token"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/server/controller/assets"
	"github.com/xxl6097/go-service-framework/internal/server/controller/proc"
)

func Listen(framework iface.IFramework) {
	api.GetApi().Add(proc.NewRoute(framework))
	api.GetApi().Add(assets.NewRoute())
	token.TokenUtils.Callback(func(s string) (bool, map[string]interface{}) {
		glog.Println("Callback", s)
		if s == proc.TOKEN {
			return true, nil
		}
		return false, map[string]interface{}{"msg": "msg err"}
	})
	//route.RouterUtil.SetApiPath("/v1/api")
	server.NewServer().Start(":8888")
}
