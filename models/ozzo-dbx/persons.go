package dbx

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func (qb *queryBuilder) AddPerson(person models.Person) (lastID string, err error) {
	m := person.ToMap()
	res, err := qb.database.Insert("persons", m).Execute()
	if err != nil {
		return "", errors.Wrap(err, "Failed to exec query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", errors.Wrap(err, "Failed to get lastid")
	}
	return cast.ToString(id), nil
}

func (*queryBuilder) GetPerson(id string) (person models.Person, err error) {
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
