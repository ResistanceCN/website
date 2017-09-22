package main

import (
	"./controller"
	"./middlewares"
	"./service"
	"github.com/kataras/iris"
)

func handle(app *iris.Application) {
	app.StaticWeb("/assets", "./assets")

	app.Use(middlewares.PassConfigToFrontEnd)
	app.Use(middlewares.StartSession)
	app.Use(middlewares.Authentication)

	app.Get("/", controller.Home)
	app.Get("/article/{id:int min(1)}", controller.ShowArticle)
	app.Get("/agent/{id:int min(1)}")

	auth := app.Party("/", middlewares.RedirectIfAuthenticated)
	{
		auth.Get("/login", controller.LoginPage)
		auth.Post("/login", controller.Login)
		auth.Post("/register", controller.Register)
	}

	user := app.Party("/user", middlewares.RequireAuthentication)
	{
		user.Get("/", controller.UserPage)
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

	view := iris.Django("./templates", ".html").Reload(service.Config.Debug)
	app.RegisterView(view)

	handle(app)

	app.Run(iris.Addr(service.Config.Listen))
}
