package gendry

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// AddPerson adds new person
func (qb *QueryBuilder) AddPerson(person models.Person) (lastID string, err error) {
	var insert []map[string]interface{}
	m := person.ToMap()
	insert = append(insert, m)
	query, args, err := builder.BuildInsert("persons", insert)
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

// GetPerson gets a person by id
func (qb *QueryBuilder) GetPerson(id string) (found bool, person models.Person, err error) {
	selectFields := []string{"id", "name", "city", "birth_date", "weight", "height"}
	where := map[string]interface{}{
		"id =": id,
	}
	query, args, err := builder.BuildSelect("persons", where, selectFields)
	if err != nil {
		return false, person, errors.Wrap(err, "Failed to build query")
	}
	rows, err := qb.DB.Query(query, args...)
	if err != nil {
		return false, person, errors.Wrap(err, "Falied to perform query")
	}
	err = scanner.ScanClose(rows, &person)
	if err != nil {
		if err == scanner.ErrEmptyResult {
			return false, person, nil
		}
		return false, person, errors.Wrap(err, "Failed to scan")
	}
	return true, person, nil
}

// ListPersons lists persons using a filter
func (qb *QueryBuilder) ListPersons(filter models.FilterPerson) (persons []models.Person, err error) {
	where, err := qb.GenerateWhere(&filter)
	if err != nil {
		return persons, errors.Wrap(err, "Failed to generate where")
	}
	selectFields := []string{"id", "name", "city", "birth_date", "weight", "height"}
	query, args, err := builder.BuildSelect("persons", where, selectFields)
	if err != nil {
		return persons, errors.Wrap(err, "Failed to build query")
	}
	rows, err := qb.DB.Query(query, args...)
	if err != nil {
		return persons, errors.Wrap(err, "Falied to perform query")
	}
	err = scanner.ScanClose(rows, &persons)
	if err != nil {
		return persons, errors.Wrap(err, "Failed to scan")
	}
	return persons, nil
}

// UpdatePerson updates a person by its id
func (qb *QueryBuilder) UpdatePerson(id string, person models.Person) (err error) {
	update := person.ToMap()
	delete(update, "id") // Don't update the id
	where := map[string]interface{}{
		"id =": id,
	}
	query, args, err := builder.BuildUpdate("persons", where, update)
	if err != nil {
		return errors.Wrap(err, "Failed to build query")
	}
	_, err = qb.DB.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "Failed to exec query")
	}
	return nil
}

// DeletePerson deletes a person
func (qb *QueryBuilder) DeletePerson(id string) (err error) {
	where := map[string]interface{}{
		"id =": id,
	}
	query, args, err := builder.BuildDelete("persons", where)
	if err != nil {
		return errors.Wrap(err, "Failed to build query")
	}
	_, err = qb.DB.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "Failed to exec query")
	}
	return
}
