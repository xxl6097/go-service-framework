package repository

import (
	"fmt"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"gorm.io/gorm"
)

type procrepository struct {
	db *gorm.DB
}

func (p procrepository) AddJson(v map[string]interface{}) error {
	return p.db.Debug().Model(model.ProcModel{}).Create(&v).Error
}

func (p procrepository) Add(v *model.ProcModel) error {
	return p.db.Create(v).Error
}

func (p procrepository) Update(m map[string]interface{}) error {
	return p.db.Debug().Model(model.ProcModel{}).Omit("name").Where("name = ?", m["name"]).Updates(m).Error
}

func (p procrepository) Delete(v *model.ProcModel) error {
	return p.db.Where("name = ?", v.Name).Unscoped().Delete(v).Error
}

func (p procrepository) Find(v *model.ProcModel) (*model.ProcModel, error) {
	var mm model.ProcModel
	err := p.db.
		Debug().
		Model(&model.ProcModel{}).
		Where(v).
		Take(&mm).Error
	return &mm, err
}

func (p procrepository) FindAll() ([]model.ProcModel, error) {
	var arr []model.ProcModel
	err := p.db.Find(&arr).Error
	return arr, err
}

func (p procrepository) DeleteByUniqueKey(s string) error {
	return p.db.Where("name = ?", s).Unscoped().Delete(&model.ProcModel{}).Error
}

func (p procrepository) Save(procModel *model.ProcModel) error {
	return p.db.Save(procModel).Error // 根据整型主键查找
}

func (p procrepository) First() (*model.ProcModel, error) {
	var conf model.ProcModel
	return &conf, p.db.First(&conf, 1).Error // 根据整型主键查找
}
func NewProcRepository(db *gorm.DB) iface.ISqlite[model.ProcModel] {
	//if db.Migrator().HasTable(&TestModel{}) {
	//	db.Debug().Migrator().CreateTable(&TestModel{})
	//}
	err := db.Debug().AutoMigrate(&model.ProcModel{})
	if err != nil {
		fmt.Println("TestModel created failed", err)
	} else {
		fmt.Println("TestModel created")
	}
	return &procrepository{
		db: db,
	}
}
