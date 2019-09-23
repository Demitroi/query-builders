package dbx

import (
	"github.com/Demitroi/query-builders/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
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

func (qb *queryBuilder) GetPerson(id string) (person models.Person, err error) {
	selectFields := []string{"id", "name", "city", "birthdate", "weight", "height"}
	conditions := dbx.HashExp{"id": id}
	err = qb.database.Select(selectFields...).From("persons").Where(conditions).One(&person)
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
