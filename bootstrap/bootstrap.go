package bootstrap

import (
	"github.com/hongyukeji/easy-go/routes"
	"github.com/hongyukeji/easy-go/utils"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	"github.com/hongyukeji/easy-go/middleware/identity"
	"log"
)

func NewApplication() *iris.Application {
	app := iris.New()

	//app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./views", ".html").Layout("shared/layout.html").Reload(true))

	app.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})

	const (
		// StaticAssets is the root directory for public assets like images, css, js.
		StaticAssets = "./public/"
		// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
		Favicon = "favicon.ico"
	)

	// static files
	app.Favicon(StaticAssets + Favicon)
	app.HandleDir(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	app.Use(recover.New())
	app.Use(logger.New())

	app.Use(iris.Gzip)
	app.Use(iris.NoCache)

	app.Configure(identity.Configure, routes.Configure)

	app.Configure(iris.WithConfiguration(iris.YAML("./env.yml")))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return app
}

func StartApplication(app *iris.Application) {
	/*addr, err := app.ConfigurationReadOnly().GetOther()["AppAddr"].(string)
	if err != true {
		addr = ":8080"
	}*/
	app.Listen(utils.GetEnv("APP_ADDR", ":8080"))
}
