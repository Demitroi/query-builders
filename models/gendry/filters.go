package gendry

import (
	"fmt"

	"github.com/Demitroi/query-builders/models"
	"github.com/pkg/errors"
)

// GenerateWhere generates the where map
func (qb *QueryBuilder) GenerateWhere(filter models.Filter) (where map[string]interface{}, err error) {
	where = make(map[string]interface{})
	err = filter.ForEach(func(field string, operator string, value interface{}) error {
		switch operator {
		case "=":
			k := fmt.Sprintf("%s =", field)
			where[k] = value
		case ">":
			k := fmt.Sprintf("%s >", field)
			where[k] = value
		case "<":
			k := fmt.Sprintf("%s <", field)
			where[k] = value
		case "<=":
			k := fmt.Sprintf("%s <=", field)
			where[k] = value
		case ">=":
			k := fmt.Sprintf("%s >=", field)
			where[k] = value
		case "!=":
			k := fmt.Sprintf("%s !=", field)
			where[k] = value
		case "<>":
			k := fmt.Sprintf("%s <>", field)
			where[k] = value
		case "in":
			k := fmt.Sprintf("%s in", field)
			where[k] = value
		case "not in":
			k := fmt.Sprintf("%s not in", field)
			where[k] = value
		case "like":
			k := fmt.Sprintf("%s like", field)
			where[k] = value
		case "not like":
			k := fmt.Sprintf("%s not like", field)
			where[k] = value
		case "between":
			k := fmt.Sprintf("%s between", field)
			where[k] = value
		case "not between":
			k := fmt.Sprintf("%s not between", field)
			where[k] = value
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
