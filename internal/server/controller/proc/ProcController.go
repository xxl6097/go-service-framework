package proc

import (
	"errors"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/util"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"github.com/xxl6097/go-service-framework/pkg/crypt"
	"net/http"
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
