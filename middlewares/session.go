package middlewares

import (
	"../service"
	"github.com/kataras/iris"
)

func StartSession(ctx iris.Context) {
	session := service.Sessions.Start(ctx)
	ctx.Values().Set("session", session)

	ctx.Next()
}
