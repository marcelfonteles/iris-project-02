package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"io/ioutil"
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
	b.Handle("GET","/findby/{field}/{value}","FindBy")
	b.Handle("GET","/findall/{field}/{value}","FindAll")
	b.Handle("POST","/create","CreatePerson")
	b.Handle("DELETE","/{businessEntityID}/delete","DeletePerson")
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
	param := ctx.Params().Get("field")
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

func (p *peopleController) FindAll(ctx iris.Context) {
	field := ctx.Params().Get("field")
	value := ctx.Params().Get("value")
	data := []interface{}{field, value}
	people, err := p.Person.FindAll(data...)
	if err != nil {
		error := map[string]interface{}{"error":"Could not find any record with combination informed.",
		 						         "err":err}
		ctx.JSON(error)
		return
	}
	ctx.JSON(people)
}

func (p *peopleController) CreatePerson(ctx iris.Context) {
	body := ctx.Request().Body
	binData, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	rawData := string(binData)
	fmt.Println(rawData)

	person, err := p.Person.NewSingle(rawData)
	if err != nil {
		fmt.Println("Erro 1:", err)
	}
	if err = person.Save(); err != nil {
		fmt.Println("Erro 2:", err)
	}
}

func (p *peopleController) DeletePerson(ctx iris.Context) {
	id := ctx.Params().Get("businessEntityID")
	fmt.Println("id:",id)
	data := []interface{}{"BusinessEntityID", id}
	person, err := p.Person.Find(data...)
	if err != nil {
		fmt.Println("erro:", err)
	}
	if err = person.Delete(); err != nil {
		fmt.Println("erro 2:", err)
	}
}