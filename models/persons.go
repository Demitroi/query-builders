package models

import (
	"time"

	"github.com/tiaotiao/mapstruct"
)

// Person represents a customer model
type Person struct {
	ID        *string    `json:"id"`
	Name      *string    `json:"name"`
	City      *string    `json:"city"`
	BirthDate *time.Time `json:"birthdate"`
	Weight    *float32   `json:"weight"`
	Height    *float32   `json:"height"`
}

// ToMap converts person struct to mpa
func (p *Person) ToMap() map[string]interface{} {
	return mapstruct.Struct2MapTag(p, "json")
}

// PersonMethods represents the person's methods that must be implemented
type PersonMethods interface {
	AddPerson(person Person) (lastID string, err error)
	GetPerson(id string) (person Person, err error)
	ListPersons(filter map[string]interface{}) (persons []Person, err error)
	UpdatePerson(id string, person Person) (err error)
	DeletePerson(id string) (err error)
}
