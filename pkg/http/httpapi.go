package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"io/ioutil"
	"net/http"
	"time"
)

func Post(apihost, param string, headers map[string]interface{}, requestBody []byte, timeout time.Duration) ([]byte, error) {
	return request(apihost, param, http.MethodPost, headers, requestBody, timeout)
}

func Get(apihost string, headers map[string]interface{}, timeout time.Duration) ([]byte, error) {
	return request(apihost, "", http.MethodGet, headers, nil, timeout)
}

func request(apihost, param, method string, headers map[string]interface{}, requestBody []byte, timeout time.Duration) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}
	url := fmt.Sprintf("%s%s", apihost, param)
	var req *http.Request
	var err error
	if requestBody != nil {
		var body *bytes.Buffer
		body = bytes.NewBuffer(requestBody)
		glog.Info("post", url, string(requestBody))
		req, err = http.NewRequest(method, url, body) //http.MethodPost
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		glog.Error(url, err)
		return nil, err
	}

	if headers != nil {
		//req.Header.Set("Content-Type", "application/json")
		for k, v := range headers {
			req.Header.Add(k, v.(string))
		}
	}

	resp, err1 := client.Do(req)
	if err1 != nil {
		glog.Error("do", url, err1)
		return nil, err1
	}
	glog.Info("StatusCode:", resp.StatusCode)
	if resp.StatusCode == 200 {
		// 读取响应体
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			glog.Error("Error reading response:", err2)
			return nil, err2
		}
		// 输出响应体
		glog.Info("Response:", string(body))
		return body, nil
	} else {
		glog.Error("err StatusCode", resp.StatusCode, resp.Body)
		return nil, nil
	}
}

func Get1(url string) ([]byte, error) {
	// 创建一个HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// 打印响应体内容
	fmt.Println("Response body:", string(body))
	return body, nil
}

func GetReqData[T any](w http.ResponseWriter, r *http.Request) *T {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Error("DecodeBody error", err)
		return nil
	}
	return &t
}
