package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func ApplicationController(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})
	app.Handle(new(applicationController))
}

type applicationController struct {}

func (a *applicationController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Index")
}

func (a *applicationController) Index(ctx iris.Context) {
	data := map[int]string{
		0: "Hello World",
		1: "Welcome",
		2: "This API serves data from Adventure Works 2017, a MS SQL database",
	}
	ctx.JSON(data)
}