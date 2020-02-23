package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestGetPersonByID(t *testing.T) {
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
		req, err := http.NewRequest(http.MethodGet, "/persons/1", nil)
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

func TestGetPersonByIDNotFound(t *testing.T) {
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
	fn := func() {
		// Request GET persons must query a select
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows(columns))
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create a request
		req, err := http.NewRequest(http.MethodGet, "/persons/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		// Do the request and catch the response code and body
		app.ServeHTTP(w, req)
		if w.Code != http.StatusNotFound {
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

func TestGetPersonByIDInternalServerError(t *testing.T) {
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
		req, err := http.NewRequest(http.MethodGet, "/persons/1", nil)
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

func TestAddPerson(t *testing.T) {
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
		// Request POST persons must exec an insert
		result := sqlmock.NewResult(1, 1)
		mock.ExpectExec("INSERT INTO").
			WillReturnResult(result)
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create body reader
		body := strings.NewReader(`{
			"name": "Ash Ketchum",
			"city": "Pallet Town",
			"birth_date": "2006-01-02T15:04:05Z",
			"weight": 91,
			"height": 81
		}`)
		// Create a request
		req, err := http.NewRequest(http.MethodPost, "/persons", body)
		if err != nil {
			t.Fatal(err)
		}
		// Do the request and catch the response code and body
		app.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
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

func TestAddPersonBadRequest(t *testing.T) {
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
	req, err := http.NewRequest(http.MethodPost, "/persons", nil)
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

func TestAddPersonInternalServerError(t *testing.T) {
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
		// Request POST persons must exec an insert
		mock.ExpectExec("INSERT INTO").
			WillReturnError(err)
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create body reader
		body := strings.NewReader(`{
			"name": "Ash Ketchum",
			"city": "Pallet Town",
			"birth_date": "2006-01-02T15:04:05Z",
			"weight": 91,
			"height": 81
		}`)
		// Create a request
		req, err := http.NewRequest(http.MethodPost, "/persons", body)
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

func TestUpdatePerson(t *testing.T) {
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
		// Request PUT persons must exec an insert
		result := sqlmock.NewResult(0, 1)
		mock.ExpectExec("UPDATE").
			WillReturnResult(result)
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create body reader
		body := strings.NewReader(`{
			"name": "Ash Ketchum",
			"city": "Pallet Town",
			"birth_date": "2006-01-02T15:04:05Z",
			"weight": 91,
			"height": 81
		}`)
		// Create a request
		req, err := http.NewRequest(http.MethodPut, "/persons/1", body)
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

func TestUpdatePersonBadRequest(t *testing.T) {
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
		// Request PUT persons must exec an insert
		result := sqlmock.NewResult(0, 1)
		mock.ExpectExec("UPDATE").
			WillReturnResult(result)
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create a request
		req, err := http.NewRequest(http.MethodPut, "/persons/1", nil)
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
	// Create and assign the query builders
	handlers.QueryBuilder = goqu.New(db)
	fn()
	handlers.QueryBuilder = gendry.New(db)
	fn()
	handlers.QueryBuilder = dbx.New(db)
	fn()
}

func TestUpdatePersonInternalServerError(t *testing.T) {
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
		// Request PUT persons must exec an insert
		mock.ExpectExec("UPDATE").
			WillReturnError(err)
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create body reader
		body := strings.NewReader(`{
			"name": "Ash Ketchum",
			"city": "Pallet Town",
			"birth_date": "2006-01-02T15:04:05Z",
			"weight": 91,
			"height": 81
		}`)
		// Create a request
		req, err := http.NewRequest(http.MethodPut, "/persons/1", body)
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

func TestDeletePerson(t *testing.T) {
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
		// Request DELETE persons must exec an delete from
		result := sqlmock.NewResult(0, 0)
		mock.ExpectExec("DELETE FROM").
			WillReturnResult(result)
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create a request
		req, err := http.NewRequest(http.MethodDelete, "/persons/1", nil)
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

func TestDeletePersonINternalServerError(t *testing.T) {
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
		// Request DELETE persons must exec an delete from
		mock.ExpectExec("DELETE FROM").
			WillReturnError(err)
		// Create recorder and serve
		w := httptest.NewRecorder()
		// Create a request
		req, err := http.NewRequest(http.MethodDelete, "/persons/1", nil)
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
