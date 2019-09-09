package goqu

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/doug-martin/goqu/v8"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func (qb *queryBuilder) AddPerson(person models.Person) (lastID string, err error) {
	m := person.ToMap()
	query, args, err := qb.database.Insert("persons").Rows(m).ToSQL()
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
	selectFields := []interface{}{"id", "name", "city", "birthdate", "weight", "height"}
	where := goqu.Ex{
		"id": id,
	}
	_, err = qb.database.From("persons").Select(selectFields...).Where(where).Prepared(true).ScanStruct(&person)
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
