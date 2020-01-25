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
	ID             *string    `field:"id"         form:"id"            operator:"="`
	Name           *string    `field:"name"       form:"name"          operator:"="`
	City           *string    `field:"city"       form:"city"          operator:"="`
	BirthDate      *time.Time `field:"birth_date" form:"birth_dateeq"  operator:"="`
	BirthDateStart *time.Time `field:"birth_date" form:"birth_dategte" operator:">="`
	BirthDateEnd   *time.Time `field:"birth_date" form:"birth_datelte" operator:"<="`
	Weight         *float32   `field:"weight"     form:"weight"        operator:"="`
	WeightStart    *float32   `field:"weight"     form:"weight"        operator:">="`
	WeightEnd      *float32   `field:"weight"     form:"weight"        operator:"<="`
	Height         *float32   `field:"height"     form:"height"        operator:"="`
	HeightStart    *float32   `field:"height"     form:"height"        operator:">="`
	HeightEnd      *float32   `field:"height"     form:"height"        operator:"<="`
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
