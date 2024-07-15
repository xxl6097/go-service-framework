package repository

import (
	"fmt"
	"github.com/xxl6097/go-service-framework/internal/iface"
	"github.com/xxl6097/go-service-framework/internal/model"
	"gorm.io/gorm"
)

type passrepository struct {
	db *gorm.DB
}

func (p passrepository) AddJson(v map[string]interface{}) error {
	return p.db.Debug().Model(model.ConfigModel{}).Create(&v).Error
}

func (p passrepository) Add(v *model.ConfigModel) error {
	return p.db.Create(v).Error
}

func (p passrepository) Update(m map[string]interface{}) error {
	return nil
}

func (p passrepository) Delete(v *model.ConfigModel) error {
	return p.db.Where("password = ?", v.Password).Unscoped().Delete(v).Error
}

func (p passrepository) Find(v *model.ConfigModel) (*model.ConfigModel, error) {
	var mm model.ConfigModel
	err := p.db.
		Debug().
		Model(&model.ConfigModel{}).
		Where(v).
		Take(&mm).Error
	return &mm, err
}

func (p passrepository) First() (*model.ConfigModel, error) {
	var conf model.ConfigModel
	return &conf, p.db.First(&conf, 1).Error // 根据整型主键查找
}

func (p passrepository) FindAll() ([]model.ConfigModel, error) {
	var arr []model.ConfigModel
	err := p.db.Find(&arr).Error
	return arr, err
}

func (p passrepository) DeleteByUniqueKey(s string) error {
	return nil
}
func (p passrepository) Save(models *model.ConfigModel) error {
	return p.db.Save(models).Error // 根据整型主键查找
}
func NewConfRepository(db *gorm.DB) iface.ISqlite[model.ConfigModel] {
	//if db.Migrator().HasTable(&TestModel{}) {
	//	db.Debug().Migrator().CreateTable(&TestModel{})
	//}
	err := db.Debug().AutoMigrate(&model.ConfigModel{})
	//db.Migrator().DropTable(&model.ConfigModel{})
	if err != nil {
		fmt.Println("TestModel created failed", err)
	} else {
		fmt.Println("TestModel created")
	}
	return &passrepository{
		db: db,
	}
}
