package models

// Filter represents the methods that must be implemented by any filter
type Filter interface {
	ForEach(fn ForEachFunc) error
}
