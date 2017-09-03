package middlewares

import (
	"../service"
	"github.com/kataras/iris/context"
)

func PassConfigToFrontEnd(ctx context.Context) {
	ctx.ViewData("google_client_id", service.Config.GoogleClientID)

	ctx.Next()
}
