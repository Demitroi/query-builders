package goqu_test

import (
	"errors"
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

func TestAddPersonError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("INSERT INTO `persons`").
		WillReturnError(errors.New("this is an error"))
	qb := goqu.New(db)
	if err != nil {
		t.Fatal(err)
	}
	var person = models.Person{}
	_, err = qb.AddPerson(person)
	if err == nil {
		t.Error("Must return an error")
	}
}

func TestGetPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	columns := []string{"id", "name", "city", "birth_date", "weight", "height"}
	birthDate, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	mock.ExpectQuery("SELECT `id`, `name`, `city`, `birth_date`, `weight`, `height` FROM `persons`").
		WithArgs("1", 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow("1", "Ash Ketchum", "Pallet Town", birthDate, float32(91), float32(1.81)))
	qb := goqu.New(db)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = qb.GetPerson("1")
	if err != nil {
		t.Error(err)
	}
}

func TestGetPersonError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `id`, `name`, `city`, `birth_date`, `weight`, `height` FROM `persons`").
		WillReturnError(errors.New("this is an error"))
	qb := goqu.New(db)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = qb.GetPerson("1")
	if err == nil {
		t.Error("Must return an error")
	}
}

func TestListPersons(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	columns := []string{"id", "name", "city", "birth_date", "weight", "height"}
	birthDate, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	mock.ExpectQuery("SELECT `id`, `name`, `city`, `birth_date`, `weight`, `height`").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow("1", "Ash Ketchum", "Pallet Town", birthDate, float32(91), float32(1.81)))
	qb := goqu.New(db)
	filterPersons := models.FilterPerson{
		ID: &[]string{"1"}[0],
	}
	persons, err := qb.ListPersons(filterPersons)
	if err != nil {
		t.Error(err)
	}
	if len(persons) == 0 {
		t.Error("Persons should not be empty")
	}
}

func TestListPersonsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT `id`, `name`, `city`, `birth_date`, `weight`, `height`").
		WillReturnError(errors.New("this is an error"))
	qb := goqu.New(db)
	filterPersons := models.FilterPerson{
		ID: &[]string{"1"}[0],
	}
	_, err = qb.ListPersons(filterPersons)
	if err == nil {
		t.Error("Must return an error")
	}
}

func TestUpdatePerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("UPDATE `persons`").
		WillReturnResult(sqlmock.NewResult(0, 1))
	qb := goqu.New(db)
	var person models.Person
	err = qb.UpdatePerson("1", person)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdatePersonError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("UPDATE `persons`").
		WillReturnError(errors.New("this is an error"))
	qb := goqu.New(db)
	var person models.Person
	err = qb.UpdatePerson("1", person)
	if err == nil {
		t.Error("Must return an error")
	}
}

func TestDeletePerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("DELETE FROM `persons`").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	qb := goqu.New(db)
	err = qb.DeletePerson("1")
	if err != nil {
		t.Error(err)
	}
}

func TestDeletePersonError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("DELETE FROM `persons`").
		WillReturnError(errors.New("this is an error"))
	qb := goqu.New(db)
	err = qb.DeletePerson("1")
	if err == nil {
		t.Error("Must return an error")
	}
}
