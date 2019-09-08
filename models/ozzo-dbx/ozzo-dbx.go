package dbx

import (
	"github.com/Demitroi/query-builders/models"
)

type queryBuilder struct{}

// New creates query builder
func New() models.QueryBuilder {
	return new(queryBuilder)
}
