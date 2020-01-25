package handlers

import (
	"github.com/Demitroi/query-builders/models"
	"github.com/kataras/iris/v12"
)

// QueryBuilder is used across the handlers
var QueryBuilder models.QueryBuilder

// GetPersons lists persons
func GetPersons(ctx iris.Context) {
	var fp models.FilterPerson
	if err := ctx.ReadForm(&fp); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	persons, err := QueryBuilder.ListPersons(fp)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(persons)
}

// GetPersonByID gets a person by its id
func GetPersonByID(ctx iris.Context) {
	id := ctx.Params().Get("id")
	found, person, err := QueryBuilder.GetPerson(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	if !found {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "not founf"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(person)
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
