package dbx

import (
	"database/sql"

	"github.com/Demitroi/query-builders/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type queryBuilder struct {
	// DB is the connection pool to database
	DB       *sql.DB
	database *dbx.DB
}

// New creates query builder
func New(db *sql.DB) models.QueryBuilder {
	database := dbx.NewFromDB(db, "mysql")
	return &queryBuilder{db, database}
}
