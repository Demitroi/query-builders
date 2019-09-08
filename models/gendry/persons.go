package gendry

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
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

func (qb *queryBuilder) GetPerson(id string) (person models.Person, err error) {
	rows, err := qb.DB.Query(`SELECT id, name, city, birthdate, weight, height
		FROM personas WHERE id=?`, id)
	if err != nil {
		return person, errors.Wrap(err, "Falied to perform query")
	}
	err = scanner.ScanClose(rows, &person)
	if err != nil {
		return person, errors.Wrap(err, "Failed to scan")
	}
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
