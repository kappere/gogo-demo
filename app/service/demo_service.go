package service

import (
	"gogo-demo/app/dao"
	"gogo-demo/app/model"

	"github.com/jinzhu/gorm"
	"wataru.com/gogo/frame/component"
	"wataru.com/gogo/frame/context"
)

type DemoService struct {
	component.Service
}

var DemoService_ *DemoService

func init() {
	context.RegistInitializer(func() {
		DemoService_ = &DemoService{
			*component.NewBaseService(),
		}
	})
}

func (s *DemoService) GetDemoById(id int) *model.Demo {
	demo := s.DoTransaction(func(tx *gorm.DB) interface{} {
		userDao := dao.NewDemoDao(tx)
		return userDao.GetDemoById(id)
	}).(*model.Demo)
	return demo
}
