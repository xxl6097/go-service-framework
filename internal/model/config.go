package model

import "gorm.io/gorm"

type ConfigModel struct {
	gorm.Model
	Password string `json:"password" gorm:"column:password;unique;not null;comment:'密码'"`
	Args     string `json:"args" gorm:"column:args;unique;not null;comment:'运行参数'"`
}
