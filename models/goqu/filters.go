package goqu

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
)

// GenerateWhere generates the where map
func (qb *QueryBuilder) GenerateWhere(filter models.Filter) (where goqu.Ex, err error) {
	where = make(goqu.Ex)
	err = filter.ForEach(func(field string, operator string, value interface{}) error {
		switch operator {
		case "=":
			where[field] = goqu.Op{"eq": value}
		case ">":
			where[field] = goqu.Op{"gt": value}
		case "<":
			where[field] = goqu.Op{"lt": value}
		case "<=":
			where[field] = goqu.Op{"lte": value}
		case ">=":
			where[field] = goqu.Op{"gte": value}
		case "!=":
			where[field] = goqu.Op{"neq": value}
		case "<>":
			where[field] = goqu.Op{"neq": value}
		case "in":
			where[field] = goqu.Op{"in": value}
		case "not in":
			where[field] = goqu.Op{"notin": value}
		case "like":
			where[field] = goqu.Op{"like": value}
		case "not like":
			where[field] = goqu.Op{"not": value}
		case "between":
			where[field] = goqu.Op{"between": value}
		case "not between":
			where[field] = goqu.Op{"notbetween": value}
		default:
			return errors.Errorf("operator %s not supported", operator)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return where, nil
}
