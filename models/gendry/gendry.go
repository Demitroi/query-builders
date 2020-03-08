package gendry

import (
	"database/sql"

	"github.com/Demitroi/query-builders/models"
	"github.com/didi/gendry/scanner"
)

// QueryBuilder is the query builder to gendry
type QueryBuilder struct {
	// DB is the connection pool to database
	DB *sql.DB
}

// New creates query builder
func New(db *sql.DB) *QueryBuilder {
	scanner.SetTagName("db")
	return &QueryBuilder{db}
}

// Verify that implements the interface
var _ models.QueryBuilder = &QueryBuilder{}
