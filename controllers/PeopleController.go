package controllers

import (
	"fmt"
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

func (p *peopleController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Index")
	b.Handle("GET", "/find", "Find")
}

func (p *peopleController) Index(ctx iris.Context) {
	people, err := p.Person.All()
	if err != nil {
		fmt.Printf("Erro:", err)
	}
	ctx.JSON(people)
}

func (p *peopleController)  Find(ctx iris.Context) {
	keys := []interface{}{"BusinessEntityID", 269}
	person, err := p.Person.Find(keys...)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	ctx.JSON(person)
}

