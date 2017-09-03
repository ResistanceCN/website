package middlewares

import (
	"../util"
	"github.com/kataras/iris/context"
)

func Authentication(ctx context.Context) {
	user := util.GetUser(ctx)

	ctx.Values().Set("user", user)
	ctx.ViewData("user", user)

	ctx.Next()
}

func RedirectIfAuthenticated(ctx context.Context) {
	user := util.GetUser(ctx)

	if user.ID != 0 {
		ctx.Redirect("/")
		return
	}

	ctx.Next()
}

func RequireAuthentication(ctx context.Context) {
	user := util.GetUser(ctx)

	if user.ID == 0 {
		ctx.Redirect("/login")
		return
	}

	ctx.Next()
}
