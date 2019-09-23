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
	where := dbx.HashExp{"id": id}
	err = qb.database.Select(selectFields...).From("persons").Where(where).One(&person)
	return
}

func (qb *queryBuilder) ListPersons(filter models.FilterPerson) (persons []models.Person, err error) {
	selectFields := []string{"id", "name", "city", "birthdate", "weight", "height"}
	where, err := qb.GenerateWhere(&filter)
	if err != nil {
		return persons, errors.Wrap(err, "Failed to generate where")
	}
	err = qb.database.Select(selectFields...).From("persons").Where(where).All(&persons)
	if err != nil {
		return persons, errors.Wrap(err, "Failed to perform query")
	}
	return
}

func (qb *queryBuilder) UpdatePerson(id string, person models.Person) (err error) {
	m := person.ToMap()
	delete(m, "id") // Don't update the id
	where := dbx.HashExp{"id": id}
	_, err = qb.database.Update("persons", m, where).Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to exec query")
	}
	return
}

func (*queryBuilder) DeletePerson(id string) (err error) {
	return
}
