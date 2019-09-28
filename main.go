package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-project-02/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Configuring controllers
	mvc.Configure(app.Party("/"), controllers.ApplicationController)
	mvc.Configure(app.Party("/api/people"), controllers.PeopleController)

	app.Run(iris.Addr(":3000"))
}

