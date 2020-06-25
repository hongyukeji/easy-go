package controllers

import (
	"github.com/kataras/iris/v12"
)

type IndexController struct{}

func (c *IndexController) Get(ctx iris.Context) {
	ctx.ViewData("Title", "Index Page")
	ctx.View("index.html")
}
