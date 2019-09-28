package models
import (
	"github.com/infarmasistemas/go-abstract-record/active"
	"iris-project-02/datasource"
)

type Person struct {
	BusinessEntityID      *uint   `gar:"BusinessEntityID" json:"business_entity_id"`
	PersonType            *string `gar:"PersonType" json:"person_type"`
	NameStyle             *bool `gar:"NameStyle" json:"name_style"`
	Title                 *string `gar:"Title" json:"title"`
	FirstName             *string `gar:"FirstName" json:"first_name"`
	MiddleName            *string `gar:"MiddleName" json:"middle_name"`
	LastName              *string `gar:"LastName" json:"last_name"`
	Suffix                *string `gar:"Suffix" json:"suffix"`
	EmailPromotion        *uint   `gar:"EmailPromotion" json:"email_promotion"`
}

func (p *Person) prepare(objectArray interface{}, extraOptions ...interface{}) *active.AbstractRecord {
	return active.NewAbstractRecord(p, objectArray, datasource.ConnectDB(), extraOptions...)
}

func (p *Person) All(values ...interface{}) ([]Person, error) {
	people := make([]Person, 0)
	return people, p.prepare(&people).Inner().Where(values...)
}

func (p *Person) Find(values ...interface{}) (Person, error) {
	return *p, p.prepare(nil, true).Find(values...)
}

func (p *Person) FindAll(values ...interface{}) ([]Person, error) {
	people := make([]Person, 0)
	return people, p.prepare(&people, true).Where(values...)
}

func (p *Person) NewSingle(values ...interface{}) (Person, error) {
	if len(values) == 0 {
		return *p, nil
	}
	return *p, p.prepare(nil).New(values...)
}

func (p *Person) Save() error{
	return p.prepare(nil).Save()
}

func (p *Person) Delete() error {
	return p.prepare(nil).Delete()
}