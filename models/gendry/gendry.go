package gendry

import (
	"database/sql"

	"github.com/Demitroi/query-builders/models"
	"github.com/didi/gendry/scanner"
)

type queryBuilder struct {
	// DB is the connection pool to database
	DB *sql.DB
}

// New creates query builder
func New(db *sql.DB) models.QueryBuilder {
	scanner.SetTagName("db")
	return &queryBuilder{db}
}
