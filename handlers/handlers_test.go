package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Demitroi/query-builders/handlers"
	"github.com/Demitroi/query-builders/models/goqu"
	"github.com/kataras/iris/v12"
)

func TestGetPersons(t *testing.T) {
	// Create database mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// Create and assign the query builder
	handlers.QueryBuilder = goqu.New(db)
	// Create iris app
	app := iris.New()
	// Create party an register routes in root path
	party := app.Party("/")
	handlers.RegisterPersons(party)
	// Create a request
	req, err := http.NewRequest(http.MethodGet, "/persons", nil)
	if err != nil {
		t.Fatal(err)
	}
	// important
	err = app.Build()
	if err != nil {
		t.Fatal(err)
	}
	// Columns that must return
	columns := []string{"id", "name", "city", "birth_date", "weight", "height"}
	// Parse birthdate
	birthDate, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	// Must query
	mock.ExpectQuery("SELECT `id`, `name`, `city`, `birth_date`, `weight`, `height`").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow("1", "Ash Ketchum", "Pallet Town", birthDate, float32(91), float32(1.81)))
	// Create recorder and serve
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
}
