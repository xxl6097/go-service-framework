package proc

import (
	"errors"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/util"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"net/http"
	"strings"
)

const TOKEN = "xiaxiaoli"

type ProcController struct {
	iframework iface.IFramework
}

func NewController(iframework iface.IFramework) *ProcController {
	return &ProcController{iframework: iframework}
}
func (this *ProcController) restart(w http.ResponseWriter, r *http.Request) {
	name := util.GetRequestParam(r, "name")
	err := this.iframework.RestartProcess(name)
	glog.Error("3---restart process err: ", err)
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

func (this *ProcController) auth(w http.ResponseWriter, r *http.Request) {
	_token := r.Header.Get("accessToken")
	if strings.ToLower(TOKEN) == strings.ToLower(_token) {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(502)
	}
}
