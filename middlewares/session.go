package middlewares

import (
	"../service"
	"github.com/kataras/iris/context"
)

func StartSession(ctx context.Context) {
	session := service.Sessions.Start(ctx)
	ctx.Values().Set("session", session)

	ctx.Next()
}
