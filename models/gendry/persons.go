package gendry

import (
	"github.com/Demitroi/query-builders/models"
	builder "github.com/didi/gendry/builder"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func (qb *queryBuilder) AddPerson(person models.Person) (lastID string, err error) {
	var data []map[string]interface{}
	m := person.ToMap()
	data = append(data, m)
	query, vals, err := builder.BuildInsert("personas", data)
	if err != nil {
		return "", errors.Wrap(err, "Failed to build query")
	}
	res, err := qb.DB.Exec(query, vals...)
	if err != nil {
		return "", errors.Wrap(err, "Failed to execute query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", errors.Wrap(err, "Failed to get last id")
	}
	return cast.ToString(id), nil
}

func (*queryBuilder) GetPerson(id string) (person models.Person, err error) {
	return
}

func (*queryBuilder) ListPersons(filter map[string]interface{}) (persons []models.Person, err error) {
	return
}

func (*queryBuilder) UpdatePerson(id string, person models.Person) (err error) {
	return
}

func (*queryBuilder) DeletePerson(id string) (err error) {
	return
}
