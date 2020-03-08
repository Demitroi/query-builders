package goqu

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// AddPerson adds new person
func (qb *QueryBuilder) AddPerson(person models.Person) (lastID string, err error) {
	m := person.ToMap()
	res, err := qb.database.Insert("persons").Rows(m).Executor().Exec()
	if err != nil {
		return "", errors.Wrap(err, "Failed to exec query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", errors.Wrap(err, "Failed to get last id")
	}
	return cast.ToString(id), nil
}

// GetPerson gets a person by id
func (qb *QueryBuilder) GetPerson(id string) (found bool, person models.Person, err error) {
	selectFields := []interface{}{"id", "name", "city", "birth_date", "weight", "height"}
	where := goqu.Ex{
		"id": id,
	}
	found, err = qb.database.From("persons").Select(selectFields...).Where(where).Prepared(true).ScanStruct(&person)
	return
}

// ListPersons lists persons using a filter
func (qb *QueryBuilder) ListPersons(filter models.FilterPerson) (persons []models.Person, err error) {
	selectFields := []interface{}{"id", "name", "city", "birth_date", "weight", "height"}
	where, err := qb.GenerateWhere(&filter)
	if err != nil {
		return persons, errors.Wrap(err, "Failed to generate where")
	}
	err = qb.database.From("persons").Select(selectFields...).Where(where).Prepared(true).ScanStructs(&persons)
	if err != nil {
		return persons, errors.Wrap(err, "Failed to perform query")
	}
	return
}

// UpdatePerson updates a person by its id
func (qb *QueryBuilder) UpdatePerson(id string, person models.Person) (err error) {
	update := person.ToMap()
	delete(update, "id") // Don't update the id
	where := goqu.Ex{
		"id": id,
	}
	_, err = qb.database.Update("persons").Set(update).Where(where).Prepared(true).Executor().Exec()
	if err != nil {
		return errors.Wrap(err, "Failed to exec query")
	}
	return
}

// DeletePerson deletes a person
func (qb *QueryBuilder) DeletePerson(id string) (err error) {
	where := goqu.Ex{
		"id": id,
	}
	_, err = qb.database.Delete("persons").Where(where).Prepared(true).Executor().Exec()
	if err != nil {
		return errors.Wrap(err, "Failed to exec query")
	}
	return
}
