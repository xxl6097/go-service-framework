package config

import (
	"encoding/json"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/model"
	"os"
	"path/filepath"
)

func getConfigPath() (string, error) {
	fullexecpath, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir, execname := filepath.Split(fullexecpath)
	ext := filepath.Ext(execname)
	name := execname[:len(execname)-len(ext)]

	return filepath.Join(dir, name+".json"), nil
}

func getConfig(path string) ([]model.ProcModel, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var obj []model.ProcModel

	r := json.NewDecoder(f)
	err = r.Decode(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func Get() ([]model.ProcModel, error) {
	binpath, err := getConfigPath()
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return getConfig(binpath)
}

func saveConfig(fileName string, data interface{}) error {
	// 打开文件，如果文件不存在则创建它
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return err
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 将结构体编码为 JSON 并写入文件
	jsonData, err := json.MarshalIndent(data, "", "    ") // 使用 json.MarshalIndent 来美化输出
	if err != nil {
		fmt.Printf("Error marshalling data: %s\n", err)
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return err
	}

	fmt.Printf("JSON data written to %s\n", fileName)
	return nil
}

func Save(data model.ProcModel) error {
	binpath, err := getConfigPath()
	if err != nil {
		glog.Error(err)
		return err
	}
	arr, _ := getConfig(binpath)
	if arr == nil {
		arr = make([]model.ProcModel, 0)
	}
	arr = append(arr, data)
	return saveConfig(binpath, arr)
}
