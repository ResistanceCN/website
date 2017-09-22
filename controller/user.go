package controller

import (
	"../util"
	"../db"
	"github.com/Automattic/go-gravatar"
	"github.com/kataras/iris/context"
)

func UserPage(ctx context.Context) {
	user := util.GetUser(ctx)

	avatar := gravatar.NewGravatarFromEmail(user.Email)
	avatar.Size = 240

	var articles []db.Article
	db.Conn().Model(&user).Related(&articles)

	ctx.ViewData("user_avatar_url", avatar.GetURL())
	ctx.ViewData("user", user)
	ctx.ViewData("articles", articles)
	ctx.View("user.html")
}
