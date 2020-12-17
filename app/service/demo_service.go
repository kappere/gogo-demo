package service

import (
	"gogo-demo/app/dao"
	"gogo-demo/app/model"

	"github.com/jinzhu/gorm"
	"wataru.com/gogo/frame/component"
	"wataru.com/gogo/frame/context"
)

type DemoService struct {
	*component.Service `autowired:""`
	DemoDao            *dao.DemoDao `autowired:""`
}

func init() {
	context.RegistBean(&DemoService{})
}

func (s *DemoService) GetDemoById(id int) *model.Demo {
	demo := s.DoTransaction(func(tx *gorm.DB) interface{} {
		return dao.NewDemoDao(tx).GetDemoById(id)
	}).(*model.Demo)
	return demo
}
