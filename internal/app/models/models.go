package models

import (
	"encoding/xml"
	"time"
)
type Categories struct{
	Id int16
	Nom string
}

type Livres struct{
	Iden string
	Description string
	Date time.Time
	Lieu string
	Editions string
	Exemplaire int16
	Images string
	Categoried Categories
}

type Consultes struct{
Date time.Time
Livred Livres
}

type Theses struct{
	Id string
	Name string
	Lien string
	Images string
}

type Telecharge struct{
	Dates time.Time
	Theses Theses
}

//course models

type Person struct{
	XMLName xml.Name `xml :"person"`
	FirstName string `xml :"firstName,attr"`
	LastName string `xml :"lastName, attr"`
}

