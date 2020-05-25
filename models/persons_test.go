package models_test

import (
	"testing"
	"time"

	"github.com/Demitroi/query-builders/models"
)

func TestToMap(t *testing.T) {
	birthDate, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		t.Fatal(err)
	}
	var person = models.Person{
		Name:      &[]string{"Ash Ketchum"}[0],
		City:      &[]string{"Pallet Town"}[0],
		BirthDate: &birthDate,
		Weight:    &[]float32{91}[0],
	}
	m := person.ToMap()
	if m == nil {
		t.Error("Map must not be nil")
	}
}

func TestForEach(t *testing.T) {
	var id string = "1"
	filter := models.FilterPerson{
		ID: &id,
	}
	filter.ForEach(func(field, operator string, value interface{}) error {
		if field != "id" {
			t.Error("Field must be id")
		}
		return nil
	})
}
