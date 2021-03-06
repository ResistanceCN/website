package controller

import (
	"../db"
	"../util"
	"github.com/Automattic/go-gravatar"
	"github.com/kataras/iris"
)

func UserPage(ctx iris.Context) {
	user := util.GetUser(ctx)

	avatar := gravatar.NewGravatarFromEmail(user.Email)
	avatar.Size = 240

	var articles []db.Article
	db.Conn().Model(&user).Order("created_at desc, id desc").Related(&articles)

	articlesByYears := make(map[int][]db.Article)

	for _, article := range articles {
		articlesByYear, ok := articlesByYears[article.CreatedAt.Year()]

		if !ok {
			articlesByYear = make([]db.Article, 0)
		}

		articlesByYear = append(articlesByYear, article)
		articlesByYears[article.CreatedAt.Year()] = articlesByYear
	}

	ctx.ViewData("user_avatar_url", avatar.GetURL())
	ctx.ViewData("user", user)
	ctx.ViewData("articles_by_year", articlesByYears)
	ctx.View("user/user.html")
}

func UserEditArticle(ctx iris.Context) {
	user := util.GetUser(ctx)

	id, _ := ctx.Params().GetInt("id")
	var article db.Article
	db.Conn().First(&article, id)

	if article.ID == 0 || article.UserID != user.ID {
		ctx.NotFound()
		return
	}

	ctx.ViewData("article", article)
	ctx.View("user/article.html")
}
