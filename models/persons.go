package models

import (
	"time"

	"github.com/tiaotiao/mapstruct"
)

// Person represents a customer model
type Person struct {
	ID        *string    `json:"id"         db:"id"`
	Name      *string    `json:"name"       db:"name"`
	City      *string    `json:"city"       db:"city"`
	BirthDate *time.Time `json:"birth_date" db:"birth_date"`
	Weight    *float32   `json:"weight"     db:"weight"`
	Height    *float32   `json:"height"     db:"height"`
}

// ToMap converts person struct to mpa
func (p *Person) ToMap() map[string]interface{} {
	return mapstruct.Struct2MapTag(p, "json")
}

// FilterPerson filters the results of select statement
type FilterPerson struct {
	ID             *string    `field:"id"        operator:"="`
	Name           *string    `field:"name"      operator:"="`
	City           *[]string  `field:"city"      operator:"in"`
	BirthDate      *time.Time `field:"birthdate" operator:"="`
	BirthDateStart *time.Time `field:"birthdate" operator:">="`
	BirthDateEnd   *time.Time `field:"birthdate" operator:"<="`
	Weight         *float32   `field:"weight"    operator:"="`
	WeightStart    *float32   `field:"weight"    operator:">="`
	WeightEnd      *float32   `field:"weight"    operator:"<="`
	Height         *float32   `field:"height"    operator:"="`
	HeightStart    *float32   `field:"height"    operator:">="`
	HeightEnd      *float32   `field:"height"    operator:"<="`
}

// ForEach iterates over the FilterPerson fields
func (fp *FilterPerson) ForEach(fn ForEachFunc) error {
	return ForEachFilter(fp, fn)
}

// PersonMethods represents the person's methods that must be implemented
type PersonMethods interface {
	AddPerson(person Person) (lastID string, err error)
	GetPerson(id string) (person Person, err error)
	ListPersons(filter FilterPerson) (persons []Person, err error)
	UpdatePerson(id string, person Person) (err error)
	DeletePerson(id string) (err error)
}
