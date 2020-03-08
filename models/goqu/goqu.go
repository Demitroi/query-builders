package goqu

import (
	"database/sql"

	"github.com/Demitroi/query-builders/models"
	"github.com/doug-martin/goqu/v9"

	// import the dialect
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// QueryBuilder is the query builder to goqu
type QueryBuilder struct {
	// DB is the connection pool to database
	DB       *sql.DB
	database *goqu.Database
}

// New creates query builder
func New(db *sql.DB) *QueryBuilder {
	database := goqu.New("mysql", db)
	return &QueryBuilder{db, database}
}

// Verify that implements the interface
var _ models.QueryBuilder = &QueryBuilder{}
