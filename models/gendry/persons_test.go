package gendry_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Demitroi/query-builders/models"
	"github.com/Demitroi/query-builders/models/gendry"
)

func TestAddPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	qb := gendry.New(db)
	birthDate, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		t.Fatal(err)
	}
	var person = models.Person{
		Name:      &[]string{"Ash Ketchum"}[0],
		City:      &[]string{"Pallet Town"}[0],
		BirthDate: &birthDate,
		Weight:    &[]float32{91}[0],
		Height:    &[]float32{1.81}[0],
	}
	mock.ExpectExec("INSERT INTO personas").WillReturnResult(sqlmock.NewResult(1, 1))
	_, err = qb.AddPerson(person)
	if err != nil {
		t.Error(err)
	}
}
