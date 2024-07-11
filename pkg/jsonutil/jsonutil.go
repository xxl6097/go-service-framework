package jsonutil

import (
	"encoding/json"
	"fmt"
)

func ToJson[T any](obj T) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func ToJsonBytes(obj interface{}) []byte {
	b, _ := ToJson(obj)
	return b
}

func ToObj[T any](jsonstr string) (*T, error) {
	var obj T
	err := json.Unmarshal([]byte(jsonstr), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func JsonByte2Object[T any](bytes []byte) (*T, error) {
	var obj T
	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func JsonByte2Any(bytes []byte) (any, error) {
	var obj any
	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func JsonString2Any(json string) (any, error) {
	return JsonByte2Any([]byte(json))
}
func Any2JsonByte(obj any) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func Any2JsonString(obj any) (string, error) {
	bytes, err := Any2JsonByte(obj)
	return string(bytes), err
}

func JsonStrToMap(jsonstr string) map[string]interface{} {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(jsonstr), &obj)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return obj
}
func JsonStrToArray(jsonstr string) []interface{} {
	var obj []interface{}
	err := json.Unmarshal([]byte(jsonstr), &obj)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return obj
}

func JsonToMap(jbyte []byte) map[string]interface{} {
	var obj map[string]interface{}
	err := json.Unmarshal(jbyte, &obj)
	if err != nil {
		return nil
	}
	return obj
}

func ToObjs(jsonstr string) interface{} {
	var obj interface{}
	err := json.Unmarshal([]byte(jsonstr), &obj)
	if err != nil {
		return obj
	}
	return &obj
}

func Respond(code int, msg string, err error) interface{} {
	return map[string]interface{}{"code": code, "msg": fmt.Sprintf("%s:%s", msg, err.Error())}
}
