package main

import (
	"./controllers"
	"./middlewares"
	"./services"
	"github.com/kataras/iris"
)

func handle(app *iris.Application) {
	app.StaticWeb("/assets", "./assets")

	app.Use(middlewares.StartSession)

	app.Get("/", controllers.Home)
	app.Get("/article/{id:int min(1)}")
	app.Get("/agent/{id:int min(1)}")

	auth := app.Party("/", middlewares.RedirectIfAuthenticated)
	{
		auth.Get("/login", controllers.LoginPage)
		auth.Post("/login")
		auth.Get("/register")
		auth.Post("/register")
	}

	user := app.Party("/user", middlewares.RequireAuthentication)
	{
		user.Get("/")
		user.Put("/")

		user.Get("/settings")

		user.Get("/article")
		user.Post("/article")
		user.Get("/article/{id:int min(1)}") // Edit Page
		user.Put("/article/{id:int min(1)}") // Unpublished Only
		user.Delete("/article/{id:int min(1)}") // Unpublished Only
	}
}

func main() {
	app := iris.Default()

	view := iris.Django("./templates", ".html").Reload(services.Config.Debug)
	app.RegisterView(view)

	handle(app)

	app.Run(iris.Addr(services.Config.Listen))
}
