package dao

import (
	"gogo-demo/app/model"

	"github.com/jinzhu/gorm"
	"wataru.com/gogo/frame/component"
)

type DemoDao struct {
	component.Dao
}

func NewDemoDao(db *gorm.DB) *DemoDao {
	return &DemoDao{
		component.Dao{
			Db: db,
		},
	}
}

func (dao *DemoDao) GetDemoById(id int) *model.Demo {
	demo := model.Demo{}
	dao.Db.Take(&demo, id)
	nameNew := "2123333346"
	dao.Db.Create(&model.Demo{
		Name: &nameNew,
	})
	return &demo
}
