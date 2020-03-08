package dbx

import (
	"database/sql"

	"github.com/Demitroi/query-builders/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// QueryBuilder is the query builder to dbx
type QueryBuilder struct {
	// DB is the connection pool to database
	DB       *sql.DB
	database *dbx.DB
}

// New creates query builder
func New(db *sql.DB) *QueryBuilder {
	database := dbx.NewFromDB(db, "mysql")
	return &QueryBuilder{db, database}
}

// Verify that implements the interface
var _ models.QueryBuilder = &QueryBuilder{}
