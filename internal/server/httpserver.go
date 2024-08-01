package server

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/api"
	"github.com/xxl6097/go-http/api/static"
	"github.com/xxl6097/go-http/server"
	"github.com/xxl6097/go-http/server/token"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/server/controller/assets"
	"github.com/xxl6097/go-service-framework/internal/server/controller/proc"
)

var tokenstring string

func Listen(port int, framework iface.IFramework) {
	api.GetApi().Add(proc.NewRoute(framework))
	api.GetApi().Add(assets.NewRoute())
	api.GetApi().Add(static.NewRoute())
	token.TokenUtils.Callback(func(s string) (bool, map[string]interface{}) {
		glog.Println("Callback", s)
		//tokenstring = framework.GetPassCode()
		//if strings.EqualFold(s, tokenstring) {
		//	return true, nil
		//}
		//return false, map[string]interface{}{"msg": "msg err"}
		return true, nil
	})
	//route.RouterUtil.SetApiPath("/v1/api")
	server.NewServer().Start(fmt.Sprintf(":%d", port))
}
