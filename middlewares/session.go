package middlewares

import (
	"../services"
	"github.com/kataras/iris/context"
)

func StartSession(ctx context.Context) {
	session := services.Sessions.Start(ctx)
	ctx.Values().Set("session", session)

	ctx.Next()
}
