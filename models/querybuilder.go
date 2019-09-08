package models

// QueryBuilder represents the methods that must be implemented
type QueryBuilder interface {
	PersonMethods
}

// SelectedQueryBuilder is the QueryBuilder selected
var SelectedQueryBuilder QueryBuilder

// SelectQueryBuilder allows select QueryBuilder
func SelectQueryBuilder(querybuilder QueryBuilder) {
	SelectedQueryBuilder = querybuilder
}
