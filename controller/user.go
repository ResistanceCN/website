package controller

import (
	"../db"
	"../util"
	"github.com/Automattic/go-gravatar"
	"github.com/kataras/iris/context"
)

func UserPage(ctx context.Context) {
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
	ctx.View("user.html")
}
