package handlers

import (
	"github.com/kataras/iris"
)

// RegisterPersons registers the persons' routes
func RegisterPersons(party iris.Party) {
	party.Get("", func(ctx iris.Context) {
		ctx.Text("Hello world")
	})
}
