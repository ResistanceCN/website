package controller

import (
	"strconv"

	"../db"
	"github.com/kataras/iris/context"
	"github.com/shurcooL/github_flavored_markdown"
)

func ShowArticle(ctx context.Context) {
	articleID, _ := strconv.Atoi(ctx.Params().Get("id"))

	article := db.Article{
		ID: articleID,
	}

	if db.Conn().Where("published = true").First(&article).RecordNotFound() {
		ctx.NotFound()
		// TODO
		return
	}

	db.Conn().Model(&article).Related(&article.User)

	content := string(github_flavored_markdown.Markdown([]byte(article.Content)))

	ctx.ViewData("article", article)
	ctx.ViewData("content", content)
	ctx.View("article.html")
}
