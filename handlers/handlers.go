package handlers

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/kataras/iris/v12"
)

// QueryBuilder is used across the handlers
var QueryBuilder models.QueryBuilder

// GetPersons lists persons
func GetPersons(ctx iris.Context) {
	var fp = models.FilterPerson{}
	persons, err := QueryBuilder.ListPersons(fp)
	if err != nil {
		ctx.JSON(iris.Map{"error": err.Error()})
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.JSON(persons)
	ctx.StatusCode(iris.StatusOK)
}

// GetPersonByID gets a person by its id
func GetPersonByID(ctx iris.Context) {

}

// AddPerson adds a new person
func AddPerson(ctx iris.Context) {

}

// UpdatePerson uptades a person
func UpdatePerson(ctx iris.Context) {

}

// DeletePerson deletes a person
func DeletePerson(ctx iris.Context) {

}
