package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"os"
	"path/filepath"
	"sync"
)

// 互斥锁
var mu sync.Mutex

type cache struct {
	config *model.ConfigModel
	binDir string
}

func NewCache(bindir string) iface.ICache {
	return &cache{
		binDir: bindir,
	}
}

func (c *cache) getConfigPath() string {
	//fullexecpath, err := os.Executable()
	//if err != nil {
	//	return "", err
	//}
	//dir, _ := filepath.Split(fullexecpath)
	//ext := filepath.Ext(execname)
	//name := execname[:len(execname)-len(ext)]
	return filepath.Join(c.binDir, "config.json")
}

func (c *cache) HasCache() bool {
	dir := c.getConfigPath()
	if _, err1 := os.Stat(dir); !os.IsNotExist(err1) {
		return true
	}
	return false
}

func (c *cache) Read() *model.ConfigModel {
	confPath := c.getConfigPath()
	if _, err1 := os.Stat(confPath); os.IsNotExist(err1) {
		//glog.Error(err1)
		return nil
	}
	f, err := os.Open(confPath)
	if err != nil {
		glog.Error(err)
		return nil
	}
	defer f.Close()

	var obj model.ConfigModel

	r := json.NewDecoder(f)
	err = r.Decode(&obj)
	if err != nil {
		glog.Error(err)
		return nil
	}
	c.config = &obj
	return &obj
}

func (c *cache) Get() *model.ConfigModel {
	if c.config != nil {
		return c.config
	}
	return c.Read()
}

func (c *cache) Set(data *model.ConfigModel) error {
	if data == nil {
		return errors.New("data is nil")
	}
	c.config = data
	binpath := c.getConfigPath()
	// 打开文件，如果文件不存在则创建它
	file, err := os.Create(binpath)
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
	return nil
}

func (c *cache) Save(data *model.ProcModel) error {
	mu.Lock()         // 进入临界区前获取锁
	defer mu.Unlock() // 使用 defer 确保在函数退出前释放锁
	if data == nil {
		return errors.New("data is nil")
	}
	if c.config == nil {
		return errors.New("config is nil")
	}
	if c.config.Procs == nil {
		c.config.Procs = make([]model.ProcModel, 0)
		c.config.Procs = append(c.config.Procs, *data)
	} else {
		for i := 0; i < len(c.config.Procs); i++ {
			if c.config.Procs[i].Name == data.Name {
				c.config.Procs = append(c.config.Procs[0:i], c.config.Procs[i+1:]...)
				break
			}
		}
		c.config.Procs = append(c.config.Procs, *data)
	}
	return c.Set(c.config)
}

func (c *cache) Delete(name string) error {
	mu.Lock()         // 进入临界区前获取锁
	defer mu.Unlock() // 使用 defer 确保在函数退出前释放锁
	if c.config == nil {
		return errors.New("config is nil")
	}
	if c.config.Procs == nil {
		return errors.New("config Procs is nil")
	}
	for i := 0; i < len(c.config.Procs); i++ {
		if c.config.Procs[i].Name == name {
			c.config.Procs = append(c.config.Procs[0:i], c.config.Procs[i+1:]...)
			break
		}
	}
	return c.Set(c.config)
}
