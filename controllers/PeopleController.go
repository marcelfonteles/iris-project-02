package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-project-02/models"
)

func PeopleController(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path %s", ctx.Path())
		ctx.Next()
	})

	app.Handle(new(peopleController))
}

type peopleController struct {
	Person models.Person
}

