// models/person.go

package models

type Person struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Age uint `json:"age"`
	City string `json:"city"`
}