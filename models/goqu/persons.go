package goqu

import (
	"github.com/Demitroi/query-builders/models"
)

func (*queryBuilder) AddPerson(person models.Person) (lastID string, err error) {
	return
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
