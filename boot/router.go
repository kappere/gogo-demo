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
			group.Get("/query", controller.DemoController_, controller.DemoController_.QueryHandler)
			group.Get("/queryPage", controller.DemoController_, controller.DemoController_.QueryHandler2)
		})
	})
}
