package handlers

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

// RegisterPersons registers the persons' routes
func RegisterPersons(party iris.Party) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodDelete},
	})
	persons := party.Party("/persons", crs).AllowMethods(iris.MethodOptions)
	{
		persons.Get("", GetPersons)
		persons.Get("/{id:int}", GetPersonByID)
		persons.Post("", AddPerson)
		persons.Put("/{id:int}", UpdatePerson)
		persons.Delete("/{id:int}", DeletePerson)
	}
}
