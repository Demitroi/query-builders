package gendry

import (
	"database/sql"

	"github.com/Demitroi/query-builders/models"
)

type queryBuilder struct {
	// DB is the connection pool to database
	DB *sql.DB
}

// New creates query builder
func New(db *sql.DB) models.QueryBuilder {
	return &queryBuilder{db}
}
