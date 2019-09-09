package goqu

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func (qb *queryBuilder) AddPerson(person models.Person) (lastID string, err error) {
	m := person.ToMap()
	query, args, err := qb.dialect.Insert("persons").Rows(m).ToSQL()
	if err != nil {
		return "", errors.Wrap(err, "Failed to build query")
	}
	res, err := qb.DB.Exec(query, args...)
	if err != nil {
		return "", errors.Wrap(err, "Failed to exec query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", errors.Wrap(err, "Failed to get last id")
	}
	return cast.ToString(id), nil
}

func (qb *queryBuilder) GetPerson(id string) (person models.Person, err error) {
	selectFields := []string{"id", "name", "city", "birthdate", "weight", "height"}
	qb.dialect.From("").Select(selectFields)
	return
}

func (*queryBuilder) ListPersons(filter models.FilterPerson) (persons []models.Person, err error) {
	return
}

func (*queryBuilder) UpdatePerson(id string, person models.Person) (err error) {
	return
}

func (*queryBuilder) DeletePerson(id string) (err error) {
	return
}
