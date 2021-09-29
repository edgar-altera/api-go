package models

type Person struct {
	Name string 
	Age int
	Degree string
	Status bool
}

type Address struct {
	Street string `bson:"street" json:"street"`
	Number int
	Status bool 
}

type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}