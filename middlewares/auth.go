package middlewares

import (
	"../util"
	"github.com/kataras/iris"
)

func Authentication(ctx iris.Context) {
	user := util.GetUser(ctx)

	ctx.Values().Set("user", user)
	ctx.ViewData("user", user)

	ctx.Next()
}

func RedirectIfAuthenticated(ctx iris.Context) {
	user := util.GetUser(ctx)

	if user.ID != 0 {
		ctx.Redirect("/")
		return
	}

	ctx.Next()
}

func RequireAuthentication(ctx iris.Context) {
	user := util.GetUser(ctx)

	if user.ID == 0 {
		ctx.Redirect("/login")
		return
	}

	ctx.Next()
}
