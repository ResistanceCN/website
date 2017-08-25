package controllers

import (
	"github.com/kataras/iris/context"
)

func Home(ctx context.Context) {
	ctx.ViewData("hello", "Hello!")
	ctx.View("home.html")
}
