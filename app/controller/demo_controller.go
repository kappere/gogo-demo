package controller

import (
	"gogo-demo/app/service"
	"strconv"

	"wataru.com/gogo/frame/component"
	"wataru.com/gogo/frame/context"
	"wataru.com/gogo/logger"
)

type DemoController struct {
	component.Controller
	demoService *service.DemoService
}

var DemoController_ *DemoController

func init() {
	context.RegistInitializer(func() {
		DemoController_ = &DemoController{
			Controller:  *component.NewBaseController(),
			demoService: service.DemoService_,
		}
	})
}

// 查询用户
func (u *DemoController) QueryHandler(c *context.Context) *context.Response {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.Error("%s", err.Error())
		return c.Error("参数错误")
	}
	demo := u.demoService.GetDemoById(id)
	if demo == nil {
		return c.Success(nil)
	}
	return c.Success(demo)
}
