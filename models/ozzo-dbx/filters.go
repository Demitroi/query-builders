package dbx

import (
	"fmt"

	"github.com/Demitroi/query-builders/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// GenerateWhere generates the where map
func (qb *QueryBuilder) GenerateWhere(filter models.Filter) (where dbx.Expression, err error) {
	var expressions []dbx.Expression
	err = filter.ForEach(func(field string, operator string, value interface{}) error {
		switch operator {
		case "=":
			exp := dbx.HashExp{field: value}
			expressions = append(expressions, exp)
		case ">":
			e := fmt.Sprintf("%s>{:%s}", field, field)
			exp := dbx.NewExp(e, dbx.Params{field: value})
			expressions = append(expressions, exp)
		case "<":
			e := fmt.Sprintf("%s<{:%s}", field, field)
			exp := dbx.NewExp(e, dbx.Params{field: value})
			expressions = append(expressions, exp)
		case "<=":
			e := fmt.Sprintf("%s<={:%s}", field, field)
			exp := dbx.NewExp(e, dbx.Params{field: value})
			expressions = append(expressions, exp)
		case ">=":
			e := fmt.Sprintf("%s>={:%s}", field, field)
			exp := dbx.NewExp(e, dbx.Params{field: value})
			expressions = append(expressions, exp)
		case "!=":
			e := fmt.Sprintf("%s!={:%s}", field, field)
			exp := dbx.NewExp(e, dbx.Params{field: value})
			expressions = append(expressions, exp)
		case "<>":
			e := fmt.Sprintf("%s<>{:%s}", field, field)
			exp := dbx.NewExp(e, dbx.Params{field: value})
			expressions = append(expressions, exp)
		case "in":
			exp := dbx.In(field, value)
			expressions = append(expressions, exp)
		case "not in":
			exp := dbx.NotIn(field, value)
			expressions = append(expressions, exp)
		case "like":
			exp := dbx.Like(field, cast.ToString(value))
			expressions = append(expressions, exp)
		case "not like":
			exp := dbx.Like(field, cast.ToString(value))
			expressions = append(expressions, exp)
		case "between":
			slice, err := cast.ToSliceE(value)
			if err != nil {
				return err
			}
			if len(slice) < 2 {
				return errors.Errorf("%s field value must be a slice at least two", field)
			}
			exp := dbx.Between(field, slice[0], slice[1])
			expressions = append(expressions, exp)
		case "not between":
			slice, err := cast.ToSliceE(value)
			if err != nil {
				return err
			}
			if len(slice) < 2 {
				return errors.Errorf("%s field value must be a slice at least two", field)
			}
			exp := dbx.NotBetween(field, slice[0], slice[1])
			expressions = append(expressions, exp)
		default:
			return errors.Errorf("operator %s not supported", operator)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dbx.And(expressions...), nil
}
