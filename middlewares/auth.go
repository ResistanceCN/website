package middlewares

import (
	"../helpers"
	"github.com/kataras/iris/context"
)

func RedirectIfAuthenticated(ctx context.Context) {
	user := helpers.GetUser(ctx)

	if user.ID != 0 {
		ctx.Redirect("/")
		return
	}

	ctx.Next()
}

func RequireAuthentication(ctx context.Context) {
	user := helpers.GetUser(ctx)

	if user.ID == 0 {
		ctx.Redirect("/login")
		return
	}

	ctx.Values().Set("user", user)

	ctx.Next()
}
