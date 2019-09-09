package goqu_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Demitroi/query-builders/models"
	"github.com/Demitroi/query-builders/models/goqu"
)

func TestAddPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("INSERT INTO `persons`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	qb := goqu.New(db)
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
	_, err = qb.AddPerson(person)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	columns := []string{"id", "name", "city", "birthdate", "weight", "height"}
	birthDate, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	mock.ExpectQuery("SELECT `id`, `name`, `city`, `birthdate`, `weight`, `height` FROM `persons`").
		WithArgs("1", 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow("1", "Ash Ketchum", "Pallet Town", birthDate, float32(91), float32(1.81)))
	qb := goqu.New(db)
	if err != nil {
		t.Fatal(err)
	}
	_, err = qb.GetPerson("1")
	if err != nil {
		t.Error(err)
	}
}
