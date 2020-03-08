package dbx

import (
	"database/sql"

	"github.com/Demitroi/query-builders/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// AddPerson adds new person
func (qb *QueryBuilder) AddPerson(person models.Person) (lastID string, err error) {
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

// GetPerson gets a person by id
func (qb *QueryBuilder) GetPerson(id string) (found bool, person models.Person, err error) {
	selectFields := []string{"id", "name", "city", "birth_date", "weight", "height"}
	where := dbx.HashExp{"id": id}
	err = qb.database.Select(selectFields...).From("persons").Where(where).One(&person)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, person, nil
		}
		return false, person, errors.Wrap(err, "Failed to perform query")
	}
	return true, person, nil
}

// ListPersons lists persons using a filter
func (qb *QueryBuilder) ListPersons(filter models.FilterPerson) (persons []models.Person, err error) {
	selectFields := []string{"id", "name", "city", "birth_date", "weight", "height"}
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

// UpdatePerson updates a person by its id
func (qb *QueryBuilder) UpdatePerson(id string, person models.Person) (err error) {
	m := person.ToMap()
	delete(m, "id") // Don't update the id
	where := dbx.HashExp{"id": id}
	_, err = qb.database.Update("persons", m, where).Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to exec query")
	}
	return
}

// DeletePerson deletes a person
func (qb *QueryBuilder) DeletePerson(id string) (err error) {
	where := dbx.HashExp{"id": id}
	_, err = qb.database.Delete("persons", where).Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to exec query")
	}
	return
}
