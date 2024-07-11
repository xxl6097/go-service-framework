package proc

import (
	"errors"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/util"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/pkg/crypt"
	"github.com/xxl6097/go-service-framework/pkg/jsonutil"
	"github.com/xxl6097/go-service-framework/pkg/os"
	"net/http"
	"runtime"
	"strings"
)

type ProcController struct {
	iframework iface.IFramework
}

func NewController(iframework iface.IFramework) *ProcController {
	return &ProcController{iframework: iframework}
}
func (this *ProcController) restart(w http.ResponseWriter, r *http.Request) {
	name := util.GetRequestParam(r, "name")
	err := this.iframework.RestartProcess(name)
	if err != nil {
		glog.Errorf("restart process %s error: %s", name, err.Error())
	}
	if err == nil {
		Respond(w, Sucessfully())
	} else {
		Respond(w, Errors(err))
	}
}
func (this *ProcController) stop(w http.ResponseWriter, r *http.Request) {
	name := util.GetRequestParam(r, "name")
	err := this.iframework.StopProcess(name)
	if err == nil {
		Respond(w, Sucessfully())
	} else {
		Respond(w, Errors(err))
	}
}
func (this *ProcController) start(w http.ResponseWriter, r *http.Request) {
	name := util.GetRequestParam(r, "name")
	err := this.iframework.StartProcess(name)
	if err == nil {
		Respond(w, Sucessfully())
	} else {
		Respond(w, Errors(err))
	}
}
func (this *ProcController) del(w http.ResponseWriter, r *http.Request) {
	name := util.GetRequestParam(r, "name")
	err := this.iframework.Delete(name)
	if err == nil {
		Respond(w, Sucessfully())
	} else {
		Respond(w, Errors(err))
	}
}
func (this *ProcController) getall(w http.ResponseWriter, r *http.Request) {
	arrays := this.iframework.GetAll()
	glog.Warn("Test---->", arrays)
	Respond(w, Sucess(arrays))
}
func (this *ProcController) info(w http.ResponseWriter, r *http.Request) {
	arrays := os.GetOsInfo()
	Respond(w, Sucess(arrays))
}

func (this *ProcController) new(w http.ResponseWriter, r *http.Request) {
	req := util.GetReqData[model.ProcModel](w, r)
	if req != nil {
		glog.Warn("resp---->", req)
		this.iframework.AddElement(req)
		Respond(w, Sucessfully())
	} else {
		Respond(w, Errors(errors.New("request is nil")))
	}
}

func (this *ProcController) login(w http.ResponseWriter, r *http.Request) {
	password := r.Header.Get("accessToken")
	glog.Debug(password)
	if sucess, token := crypt.IsPasswordOk([]byte(password)); sucess {
		res := Sucess(string(token))
		glog.Debug(res)
		Respond(w, res)
	} else {
		Respond(w, Errors(errors.New("密码错误")))
	}
}

func (this *ProcController) auth(w http.ResponseWriter, r *http.Request) {
	password := r.Header.Get("accessToken")
	glog.Debug(password)
	if crypt.IsHashOk([]byte(password)) {
		Respond(w, Sucessfully())
	} else {
		Respond(w, Errors(errors.New("密码错误")))
	}
}

func (this *ProcController) getAppList(w http.ResponseWriter, r *http.Request) {
	//applist := "[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]"
	//arrays, _ := jsonutil.JsonString2Any(applist)
	//jsonstr := "[{\"windows\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"wechat\",\"args\":[\"-d\",\"conf.toml\"],\"description\":\"微信应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"QQ\",\"args\":[\"-d\",\"qq.toml\"],\"description\":\"QQ应用程序，用于测试\"}]}},{\"linux\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"dingtalk\",\"args\":[\"-d\",\"dingtalk.toml\"],\"description\":\"dingtalk应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]}},{\"darwin\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"dingtalk\",\"args\":[\"-d\",\"dingtalk.toml\"],\"description\":\"dingtalk应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]}}]"
	//array := jsonutil.JsonStrToArray(jsonstr)
	//for k, v := range array {
	//	if strings.Compare(v[k], runtime.GOOS) == 0 {
	//		Respond(w, Sucess(v))
	//	}
	//}

	jsonstr := "{\"windows\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"wechat\",\"args\":[\"-d\",\"conf.toml\"],\"description\":\"微信应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"QQ\",\"args\":[\"-d\",\"qq.toml\"],\"description\":\"QQ应用程序，用于测试\"}]},\"linux\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"dingtalk\",\"args\":[\"-d\",\"dingtalk.toml\"],\"description\":\"dingtalk应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]},\"darwin\":{\"arm64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"dingtalk\",\"args\":[\"-d\",\"dingtalk.toml\"],\"description\":\"dingtalk应用程序，用于测试\"}],\"amd64\":[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]}}"
	maps := jsonutil.JsonStrToMap(jsonstr)
	//for k, v := range maps {
	//	if strings.Compare(k, runtime.GOOS) == 0 {
	//		Respond(w, Sucess(v))
	//		return
	//	}
	//}
	for k, v := range maps {
		if strings.Compare(k, runtime.GOOS) == 0 {
			if s, ok := v.(map[string]interface{}); ok {
				fmt.Println("Interface value is a string:", s, s[runtime.GOARCH])
				Respond(w, Sucess(s[runtime.GOARCH]))
				return
			} else {
				fmt.Println("Interface value is not a string")
			}

		}
	}
	Respond(w, Errors(errors.New("")))
}

func (this *ProcController) appMarket(w http.ResponseWriter, r *http.Request) {
	applist := "[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]"
	arrays, _ := jsonutil.JsonString2Any(applist)
	Respond(w, Sucess(arrays))
}
