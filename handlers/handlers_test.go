package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Demitroi/query-builders/handlers"
	"github.com/Demitroi/query-builders/models/gendry"
	"github.com/Demitroi/query-builders/models/goqu"
	dbx "github.com/Demitroi/query-builders/models/ozzo-dbx"
	"github.com/kataras/iris/v12"
)

func TestGetPersons(t *testing.T) {
	// Create database mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// Create iris app
	app := iris.New()
	// Create party an register routes in root path
	party := app.Party("/")
	handlers.RegisterPersons(party)
	// build the iris app
	err = app.Build()
	if err != nil {
		t.Fatal(err)
	}
	// Columns that must return
	columns := []string{"id", "name", "city", "birth_date", "weight", "height"}
	// Parse birthdate
	birthDate, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fn := func() {
		// Request GET persons must query a select
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow("1", "Ash Ketchum", "Pallet Town", birthDate, float32(91), float32(1.81)))
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create a request
		req, err := http.NewRequest(http.MethodGet, "/persons", nil)
		if err != nil {
			t.Fatal(err)
		}
		// Do the request and catch the response code and body
		app.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("%T %v %s", handlers.QueryBuilder, w.Code, w.Body.String())
			return
		}
	}
	// Create and assign the query builders
	handlers.QueryBuilder = goqu.New(db)
	fn()
	handlers.QueryBuilder = gendry.New(db)
	fn()
	handlers.QueryBuilder = dbx.New(db)
	fn()
}

func TestGetPersonsBadRequest(t *testing.T) {
	// Create iris app
	app := iris.New()
	// Create party an register routes in root path
	party := app.Party("/")
	handlers.RegisterPersons(party)
	// build the iris app
	err := app.Build()
	if err != nil {
		t.Fatal(err)
	}
	// Create recorder and serve
	w := httptest.NewRecorder()
	// Create a request
	req, err := http.NewRequest(http.MethodGet, "/persons?unvalidquerystr", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Do the request and catch the response code and body
	app.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%T %v %s", handlers.QueryBuilder, w.Code, w.Body.String())
		return
	}
}

func TestGetPersonsInternalServerError(t *testing.T) {
	// Create database mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// Create iris app
	app := iris.New()
	// Create party an register routes in root path
	party := app.Party("/")
	handlers.RegisterPersons(party)
	// build the iris app
	err = app.Build()
	if err != nil {
		t.Fatal(err)
	}
	fn := func() {
		err = errors.New("This is an error")
		// Request GET persons must query a select
		mock.ExpectQuery("SELECT").
			WillReturnError(err)
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create a request
		req, err := http.NewRequest(http.MethodGet, "/persons", nil)
		if err != nil {
			t.Fatal(err)
		}
		// Do the request and catch the response code and body
		app.ServeHTTP(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("%T %v %s", handlers.QueryBuilder, w.Code, w.Body.String())
			return
		}
	}
	// Create and assign the query builders
	handlers.QueryBuilder = goqu.New(db)
	fn()
	handlers.QueryBuilder = gendry.New(db)
	fn()
	handlers.QueryBuilder = dbx.New(db)
	fn()
}
