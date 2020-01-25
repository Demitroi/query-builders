package gendry

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func (qb *queryBuilder) AddPerson(person models.Person) (lastID string, err error) {
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

func (qb *queryBuilder) GetPerson(id string) (found bool, person models.Person, err error) {
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
		return false, person, errors.Wrap(err, "Failed to scan")
	}
	return true, person, nil
}

func (qb *queryBuilder) ListPersons(filter models.FilterPerson) (persons []models.Person, err error) {
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

func (qb *queryBuilder) UpdatePerson(id string, person models.Person) (err error) {
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

func (qb *queryBuilder) DeletePerson(id string) (err error) {
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
