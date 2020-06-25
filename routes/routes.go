package routes

import (
	"github.com/hongyukeji/easy-go/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(app *iris.Application) {

	mvc.New(app.Party("/")).Handle(new(controllers.IndexController))

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	app.Get("/test", GetTestHandler)
}

func GetTestHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "test",
	})
}
