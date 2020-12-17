package controller

import (
	"gogo-demo/app/service"
	"strconv"

	"wataru.com/gogo"
	"wataru.com/gogo/frame/component"
	"wataru.com/gogo/frame/context"
	"wataru.com/gogo/frame/router"
	"wataru.com/gogo/logger"
)

type DemoController struct {
	*component.Controller `autowired:""`
	DemoService           *service.DemoService `autowired:""`
}

func init() {
	context.RegistBean(&DemoController{})
}

func (u *DemoController) Initialize() {
	gogo.HttpServer().Router().Group("/demo", func(group *router.RouterGroup) {
		group.Get("/query", u.QueryHandler)
		group.Get("/queryPage", u.QueryHandler2)
		group.Post("/postQuery", u.QueryHandler3)
	})
}

// 查询用户
func (u *DemoController) QueryHandler(c *context.Context) interface{} {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		logger.Error("%s", err.Error())
		return c.Error("参数错误")
	}
	demo := u.DemoService.GetDemoById(id)
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
	demo := u.DemoService.GetDemoById(id)
	return c.Render("test.html", demo)
}

type test1 struct {
	Aaa int
	Bbb string
}

// 查询用户
func (u *DemoController) QueryHandler3(c *context.Context) interface{} {
	id, err := strconv.Atoi(c.Query("id"))
	ttt := test1{}
	c.BindJSON(&ttt)
	logger.Info("body: %v", ttt)

	if err != nil {
		logger.Error("%s", err.Error())
		return c.Error("参数错误")
	}
	demo := u.DemoService.GetDemoById(id)
	return c.Success(demo)
}
