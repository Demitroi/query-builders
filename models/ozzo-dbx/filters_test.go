package dbx_test

import (
	"testing"

	"github.com/Demitroi/query-builders/models"
	dbx "github.com/Demitroi/query-builders/models/ozzo-dbx"
)

type Filter struct {
	Eq  *string       `field:"eq"  operator:"="`
	Gt  *string       `field:"gt"  operator:">"`
	Lt  *string       `field:"lt"  operator:"<"`
	Lte *string       `field:"lte" operator:"<="`
	Gte *string       `field:"gte" operator:">="`
	Ne  *string       `field:"ne"  operator:"!="`
	Dif *string       `field:"dif" operator:"<>"`
	In  *string       `field:"in"  operator:"in"`
	NIn *string       `field:"nin" operator:"not in"`
	Li  *string       `field:"li"  operator:"like"`
	NLi *string       `field:"nli" operator:"not like"`
	Be  []interface{} `field:"be"  operator:"between"`
	NBe []interface{} `field:"nbe" operator:"not between"`
}

// ForEach iterates over the FilterPerson fields
func (fp *Filter) ForEach(fn models.ForEachFunc) error {
	return models.ForEachFilter(fp, fn)
}

type FilterUnsupported struct {
	Unsupported *string `field:"unsupported" operator:"unsupported"`
}

// ForEach iterates over the FilterPerson fields
func (fp *FilterUnsupported) ForEach(fn models.ForEachFunc) error {
	return models.ForEachFilter(fp, fn)
}

func TestGenerateWhere(t *testing.T) {
	qb := dbx.New(nil)
	var s string
	filter := &Filter{
		Eq:  &s,
		Gt:  &s,
		Lt:  &s,
		Lte: &s,
		Gte: &s,
		Ne:  &s,
		Dif: &s,
		In:  &s,
		NIn: &s,
		Li:  &s,
		NLi: &s,
		Be:  []interface{}{s, s},
		NBe: []interface{}{s, s},
	}
	where, err := qb.GenerateWhere(filter)
	if err != nil {
		t.Error(err)
		return
	}
	if where == nil {
		t.Error("Failed to generate where map")
		return
	}
	// Test with unsupported filter
	filterUnsupported := &FilterUnsupported{
		Unsupported: &s,
	}
	where, err = qb.GenerateWhere(filterUnsupported)
	if err == nil {
		t.Error("Error expected!")
		return
	}
	if where != nil {
		t.Error("Where map must be nil")
		return
	}
}

type FilterError struct {
	Be  *string `field:"be"  operator:"between"`
	NBe *string `field:"nbe" operator:"not between"`
}

// ForEach iterates over the FilterPerson fields
func (fp *FilterError) ForEach(fn models.ForEachFunc) error {
	return models.ForEachFilter(fp, fn)
}

func TestGenerateWhereErrors(t *testing.T) {
	qb := dbx.New(nil)
	var s string
	filter := &Filter{
		Be: []interface{}{s},
	}
	where, err := qb.GenerateWhere(filter)
	if err == nil {
		t.Error("Error must be nil")
		return
	}
	if where != nil {
		t.Error("Where map must not be nil")
		return
	}
	filter = &Filter{
		NBe: []interface{}{s},
	}
	where, err = qb.GenerateWhere(filter)
	if err == nil {
		t.Error("Error must be nil")
		return
	}
	if where != nil {
		t.Error("Where map must not be nil")
		return
	}
	filterError := &FilterError{
		Be: &s,
	}
	where, err = qb.GenerateWhere(filterError)
	if err == nil {
		t.Error("Error must be nil")
		return
	}
	if where != nil {
		t.Error("Where map must not be nil")
		return
	}
	filterError = &FilterError{
		NBe: &s,
	}
	where, err = qb.GenerateWhere(filterError)
	if err == nil {
		t.Error("Error must be nil")
		return
	}
	if where != nil {
		t.Error("Where map must not be nil")
		return
	}
}
