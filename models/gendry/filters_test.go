package gendry_test

import (
	"testing"

	"github.com/Demitroi/query-builders/models"
	"github.com/Demitroi/query-builders/models/gendry"
)

type Filter struct {
	Eq  *string `field:"eq"  operator:"="`
	Gt  *string `field:"gt"  operator:">"`
	Lt  *string `field:"lt"  operator:"<"`
	Lte *string `field:"lte" operator:"<="`
	Gte *string `field:"gte" operator:">="`
	Ne  *string `field:"ne"  operator:"!="`
	Dif *string `field:"dif" operator:"<>"`
	In  *string `field:"in"  operator:"in"`
	NIn *string `field:"nin" operator:"not in"`
	Li  *string `field:"li"  operator:"like"`
	NLi *string `field:"nli" operator:"not like"`
	Be  *string `field:"be"  operator:"between"`
	NBe *string `field:"nbe" operator:"not between"`
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
	qb := gendry.New(nil)
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
		Be:  &s,
		NBe: &s,
	}
	where, err := qb.GenerateWhere(filter)
	if err != nil {
		t.Error(err)
		return
	}
	if len(where) != 13 {
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
