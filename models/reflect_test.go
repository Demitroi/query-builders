package models_test

import (
	"testing"
	"time"

	"github.com/Demitroi/query-builders/models"
)

func TestForEachFilter(t *testing.T) {
	type FilterPerson struct {
		ID        *string    `field:"id"        operator:"="`
		Name      *string    `field:"name"      operator:"="`
		City      *[]string  `field:"city"      operator:"in"`
		BirthDate *time.Time `field:"birthdate" operator:"="`
		Weight    *float32   `field:"weight"    operator:"="`
	}
	filterPerson := FilterPerson{
		ID: &[]string{"4"}[0],
	}
	err := models.ForEachFilter(filterPerson, func(field, operator string, value interface{}) error {
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
