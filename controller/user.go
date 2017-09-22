package controller

import (
	"../util"
	"github.com/Automattic/go-gravatar"
	"github.com/kataras/iris/context"
)

func UserPage(ctx context.Context) {
	user := util.GetUser(ctx)

	avatar := gravatar.NewGravatarFromEmail(user.Email)
	avatar.Size = 240

	ctx.ViewData("user_avatar_url", avatar.GetURL())
	ctx.ViewData("user", user)
	ctx.View("user.html")
}
