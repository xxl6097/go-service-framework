package proc

import (
	"encoding/json"
	"net/http"
)

func Allow(super bool) map[string]interface{} {
	return map[string]interface{}{"is_superuser": super, "result": "allow"}
}

func Deny(super bool) map[string]interface{} {
	return map[string]interface{}{"is_superuser": super, "result": "deny"}
}

func Ignore(super bool) map[string]interface{} {
	return map[string]interface{}{"is_superuser": super, "result": "ignore"}
}

func Error(code int, msg string) map[string]interface{} {
	return map[string]interface{}{"code": code, "msg": msg}
}

func Errors(msg error) map[string]interface{} {
	return map[string]interface{}{"code": -1, "msg": msg}
}

func Sucess(data interface{}) map[string]interface{} {
	return map[string]interface{}{"code": 0, "data": data}
}

func Sucessfully() map[string]interface{} {
	return map[string]interface{}{"code": 0}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	//w.Header().Add("Access-Control-Allow-Origin", "*")
	//w.Header().Add("Access-Control-Allow-Headers", "content-type,Authorization")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	//w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	//buf, _ := json.Marshal(data)
	//glog.Info(string(buf))
	if json.NewEncoder(w).Encode(data) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//w.WriteHeader(400)
}
