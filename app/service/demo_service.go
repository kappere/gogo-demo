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
	demoDao *dao.DemoDao
}

var DemoServiceBean = &DemoService{}

func init() {
	context.Inject(func() {
		DemoServiceBean.Service = *component.NewBaseService()
		DemoServiceBean.demoDao = dao.NewDemoDao(DemoServiceBean.Service.Db)
	})
}

func (s *DemoService) GetDemoById(id int) *model.Demo {
	demo := s.DoTransaction(func(tx *gorm.DB) interface{} {
		return dao.NewDemoDao(tx).GetDemoById(id)
	}).(*model.Demo)
	return demo
}
