package controller

import (
	"../db"
	"github.com/kataras/iris/context"
)

func Home(ctx context.Context) {
	var articles []db.Article

	db.Conn().
		Where("published = true").
		Limit(12).
		Offset(0).
		Find(&articles)

	ctx.ViewData("articles", articles)

	ctx.View("home.html")
}
