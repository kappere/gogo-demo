package boot

import (
	"gogo-demo/app/controller"

	"wataru.com/gogo"
	"wataru.com/gogo/frame/context"
	"wataru.com/gogo/frame/router"
)

func InitRouter() {
	rt := gogo.HttpServer().Router()
	context.RegistInitializer(func() {
		rt.Group("/demo", func(group *router.RouterGroup) {
			group.Get("/query", controller.DemoControllerBean, controller.DemoControllerBean.QueryHandler)
			group.Get("/queryPage", controller.DemoControllerBean, controller.DemoControllerBean.QueryHandler2)
			group.Post("/postQuery", controller.DemoControllerBean, controller.DemoControllerBean.QueryHandler3)
		})
	})
}
