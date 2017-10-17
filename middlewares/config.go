package middlewares

import (
	"../service"
	"github.com/kataras/iris"
)

func PassConfigToFrontEnd(ctx iris.Context) {
	ctx.ViewData("google_client_id", service.Config.Google.ClientID)

	ctx.Next()
}
