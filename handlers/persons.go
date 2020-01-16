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
	})
	persons := party.Party("/persons", crs).AllowMethods(iris.MethodOptions)
	{
		persons.Get("", func(ctx iris.Context) {
			ctx.Text("Hello world")
		})
	}
}
