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

var DemoControllerBean = &DemoController{}

func init() {
	context.Inject(func() {
		DemoControllerBean.Controller = *component.NewBaseController()
		DemoControllerBean.demoService = service.DemoServiceBean
	})
}

// 查询用户
func (u *DemoController) QueryHandler(c *context.Context) interface{} {
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

// 查询用户页面
func (u *DemoController) QueryHandler2(c *context.Context) interface{} {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.Error("%s", err.Error())
		return c.Render("test.html", nil)
	}
	demo := u.demoService.GetDemoById(id)
	return c.Render("test.html", demo)
}
