package models

type Person struct {
	BusinessEntityID      uint   `gar:"BusinessEntityID" json:"business_entity_id"`
	PersonType            string `gar:"PersonType" json:"person_type"`
	NameStyle             string `gar:"NameStyle" json: "name_style"`
	Title                 string `gar:"Title" json: "title"`
	FirstName             string `gar:"FirstName" json: "first_name"`
	MiddleName            string `gar:"MiddleName" json: "middle_name"`
	LastName              string `gar:"LastName" json: "last_name"`
	Suffix                string `gar:"Suffix" json: "suffix"`
	EmailPromotion        uint   `gar:"EmailPromotion" json: "email_promotion"`
	AdditionalContactInfo string `gar:"AdditionalContactInfo" json: "additional_contact_info"`
}

