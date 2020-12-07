package http

import (
	"github.com/kataras/iris/v12"
	"blog-api/api"
)

func LoadRouter(app *iris.Application) {
	article(app)
}

func article(app *iris.Application) {
	app.Get("/posts/getTitle", api.HelloWorld)
	app.Get("/posts/getId", api.HelloWorld)
	app.Post("/posts/add", api.HelloWorld)
	app.Put("/posts/put", api.HelloWorld)
	app.Get("/posts/getList", api.HelloWorld)
	app.Get("/posts/getShowList", api.HelloWorld)
}
