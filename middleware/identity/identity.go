package identity

import (
	"time"

	"github.com/kataras/iris/v12"
)

// Configure creates a new identity middleware and registers that to the app.
func Configure(app *iris.Application) {
	app.UseGlobal(globalMiddleware)
}

func globalMiddleware(ctx iris.Context) {
	// response headers
	ctx.Header("App-Name", "b.AppName")
	ctx.Header("App-Owner", "b.AppOwner")
	ctx.Header("App-Since", time.Since(time.Now()).String())

	ctx.Header("Server", "Iris: https://iris-go.com")

	// view data if ctx.View or c.Tmpl = "$page.html" will be called next.
	ctx.ViewData("AppName", "b.AppName")
	ctx.ViewData("AppOwner", "b.AppOwner")
	ctx.Next()
}
