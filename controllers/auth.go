package controllers

import (
	"../services"
	"github.com/kataras/iris/context"
)

func LoginPage(ctx context.Context) {
	ctx.ViewData("google_client_id", services.Config.GoogleClientID)
	ctx.View("auth/login.html")
}

func Login(ctx context.Context) {
	//
}

func RegisterPage(ctx context.Context) {
	//
}

func Register(ctx context.Context) {
	//
}
