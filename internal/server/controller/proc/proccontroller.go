package proc

import (
	"errors"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/util"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/pkg/crypt"
	http2 "github.com/xxl6097/go-service-framework/pkg/http"
	"github.com/xxl6097/go-service-framework/pkg/jsonutil"
	"github.com/xxl6097/go-service-framework/pkg/os"
	"github.com/xxl6097/go-service-framework/pkg/version"
	"github.com/xxl6097/go-service/gservice"
	"net/http"
	"runtime"
	"strings"
	"time"
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
	//glog.Warn("Test---->", arrays)
	Respond(w, Sucess(arrays))
}
func (this *ProcController) info(w http.ResponseWriter, r *http.Request) {
	arrays := os.GetOsInfo()
	Respond(w, Sucess(arrays))
}

func (this *ProcController) settingAppStore(w http.ResponseWriter, r *http.Request) {
	url := util.GetRequestParam(r, "url")
	if url != "" {
		glog.Warn("resp---->", url)
		if this.iframework != nil {
			this.iframework.SetAppStore(url)
		}
		Respond(w, Sucessfully())
	} else {
		Respond(w, Errors(errors.New("url is nil")))
	}
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
	hashcode := this.iframework.GetPassCode()
	if crypt.ComparePassword([]byte(hashcode), []byte(password)) == nil { //sucess, token := crypt.IsPasswordOk([]byte(password)); sucess
		res := Sucess(hashcode)
		glog.Debug(res)
		Respond(w, res)
	} else {
		Respond(w, Errors(errors.New("密码错误")))
	}
}

func (this *ProcController) uninstall(w http.ResponseWriter, r *http.Request) {
	gservice.Uninstall()
}

func (this *ProcController) auth(w http.ResponseWriter, r *http.Request) {
	password := r.Header.Get("accessToken")
	//glog.Debug(password)
	if strings.EqualFold(this.iframework.GetPassCode(), password) { //crypt.IsHashOk([]byte(password))
		Respond(w, Sucess(version.VersionJson()))
	} else {
		Respond(w, Errors(errors.New("密码错误")))
	}
}

func (this *ProcController) getAppList(w http.ResponseWriter, r *http.Request) {
	if this.iframework == nil {
		Respond(w, Errors(errors.New("framework is nil")))
		return
	}
	config := this.iframework.GetConfig()
	if config == nil {
		Respond(w, Errors(errors.New("config is nil")))
		return
	}
	if config.AppStoreUrl == "" {
		Respond(w, Errors(errors.New("AppStoreUrl is nil")))
		return
	}
	response, err := http2.Get(config.AppStoreUrl, nil, time.Second*5)
	if err != nil || response == nil {
		Respond(w, Errors(errors.New("http get failed")))
		return
	}
	maps := jsonutil.JsonToMap(response)
	if maps == nil {
		Respond(w, Errors(errors.New("json parse failed")))
		return
	}
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
	Respond(w, Errors(errors.New("no app")))
}

func (this *ProcController) getAppList1(w http.ResponseWriter, r *http.Request) {
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
	Respond(w, Errors(errors.New("no app")))
}

func (this *ProcController) appMarket(w http.ResponseWriter, r *http.Request) {
	applist := "[{\"name\":\"frpc\",\"args\":[\"-c\",\"frpc.toml\"],\"description\":\"frp测试描述信息\"},{\"name\":\"surge\",\"args\":[\"-d\",\"config.toml\"],\"description\":\"surge应用程序，用于测试\"}]"
	arrays, _ := jsonutil.JsonString2Any(applist)
	Respond(w, Sucess(arrays))
}
