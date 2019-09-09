package goqu

import (
	"database/sql"

	"github.com/Demitroi/query-builders/models"
	"github.com/doug-martin/goqu/v8"

	// import the dialect
	_ "github.com/doug-martin/goqu/v8/dialect/mysql"
)

type queryBuilder struct {
	// DB is the connection pool to database
	DB      *sql.DB
	dialect goqu.DialectWrapper
}

// New creates query builder
func New(db *sql.DB) models.QueryBuilder {
	// look up the dialect
	dialect := goqu.Dialect("mysql")
	return &queryBuilder{db, dialect}
}
