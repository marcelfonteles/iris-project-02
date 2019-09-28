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
	b.Handle("GET", "/find/{value}", "Find")
	b.Handle("GET","/findby/{parameter}/{value}","FindBy")
}

func (p *peopleController) Index(ctx iris.Context) {
	people, err := p.Person.All()
	if err != nil {
		fmt.Printf("Erro:", err)
	}
	ctx.JSON(people)
}

func (p *peopleController)  Find(ctx iris.Context) {
	value, err := ctx.Params().GetInt("value")
	if err != nil {
		error := map[string]string{"error": "Invalid Value. Need to be a integer non-negative (uint)"}
		ctx.JSON(error)
		return
	}
	keys := []interface{}{"BusinessEntityID", value}
	person, err := p.Person.Find(keys...)
	if err != nil {
		error := map[string]string{"error": "Could not find a record with this BusinessEntityID"}
		ctx.JSON(error)
		return
	}

	ctx.JSON(person)
}

func (p *peopleController) FindBy(ctx iris.Context) {
	param := ctx.Params().Get("parameter")
	value := ctx.Params().Get("value")
	fmt.Println("param:",param, " value:", value)
	data := []interface{}{param, value}
	person, err := p.Person.Find(data...)
	if err != err {
		error := map[string]string{"error": "Could not find a record with this BusinessEntityID"}
		ctx.JSON(error)
		return
	}
	ctx.JSON(person)
}
