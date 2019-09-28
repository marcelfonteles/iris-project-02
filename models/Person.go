package models
import (
	"github.com/infarmasistemas/go-abstract-record/active"
	"iris-project-02/datasource"
)

type Person struct {
	BusinessEntityID      *uint   `gar:"BusinessEntityID" json:"business_entity_id"`
	PersonType            *string `gar:"PersonType" json:"person_type"`
	NameStyle             *string `gar:"NameStyle" json: "name_style"`
	Title                 *string `gar:"Title" json: "title"`
	FirstName             *string `gar:"FirstName" json: "first_name"`
	MiddleName            *string `gar:"MiddleName" json: "middle_name"`
	LastName              *string `gar:"LastName" json: "last_name"`
	Suffix                *string `gar:"Suffix" json: "suffix"`
	EmailPromotion        *uint   `gar:"EmailPromotion" json: "email_promotion"`
	AdditionalContactInfo *string `gar:"AdditionalContactInfo" json: "additional_contact_info"`
}

func (p *Person) prepare(objectArray interface{}, extraOptions ...interface{}) *active.AbstractRecord {
	return active.NewAbstractRecord(p, objectArray, datasource.ConnectDB(), extraOptions...)
}

func (p *Person) All(values ...interface{}) ([]Person, error) {
	person := make([]Person, 0)
	return person, p.prepare(&person).Inner().Where(values...)
}

func (p *Person) Find(values ...interface{}) (Person, error) {
	return *p, p.prepare(nil, true).Find(values...)
}