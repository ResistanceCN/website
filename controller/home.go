package controller

import (
	"../db"
	"github.com/kataras/iris"
)

func Home(ctx iris.Context) {
	var articles []db.Article

	db.Conn().
		Where("status = 2").
		Order("published_at desc, id desc").
		Limit(12).
		Offset(0).
		Find(&articles)

	ctx.ViewData("articles", articles)

	ctx.View("home.html")
}
