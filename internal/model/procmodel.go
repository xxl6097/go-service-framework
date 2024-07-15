package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"os"
)

const (
	STOP_NO     = 0
	STOP_EXIT   = 1
	STOP_DELETE = 2
)

type ProcModel struct {
	gorm.Model
	Name        string      `json:"name" gorm:"column:name;unique;not null;comment:'程序名称'"`
	BinUrl      string      `json:"binUrl" gorm:"column:binUrl;not null;comment:'应用程序下载链接'"`
	ConfUrl     string      `json:"confUrl" gorm:"column:confUrl;comment:'应用程序运行配置文件'"`
	Description string      `json:"description" gorm:"column:confUrl;comment:'应用描述信息'"`
	Upgrade     bool        `json:"upgrade" gorm:"column:confUrl;comment:'应用升级标志'"`
	Args        string      `json:"args" gorm:"column:confUrl;comment:'应用运行参数'"`
	Status      string      `json:"status" gorm:"column:confUrl;comment:'应用状态'"`
	Exit        int         `json:"exit" gorm:"column:confUrl;comment:'应用退出码'"`
	Proc        *os.Process `json:"-" gorm:"-"`
}

func (u *ProcModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
func (u *ProcModel) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
