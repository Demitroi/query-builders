package goqu

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/doug-martin/goqu/v8"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	// import the dialect
	_ "github.com/doug-martin/goqu/v8/dialect/mysql"
)

func (qb *queryBuilder) AddPerson(person models.Person) (lastID string, err error) {
	// look up the dialect
    dialect := goqu.Dialect("mysql")
	m := person.ToMap()
	query, args, err := dialect.Insert("personas").Rows(m).ToSQL()
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
