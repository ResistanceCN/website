package helpers

import (
	"../models"
	"../services"
	"github.com/kataras/iris/context"
)

func GetUser(ctx context.Context) models.User {
	cached := ctx.Values().Get("user")

	if cached != nil {
		return cached.(models.User)
	}

	user := models.User{}

	session := GetSession(ctx)
	id, err := session.GetInt("user")

	if err == nil {
		services.DB.First(&user, id)
	}

	return user
}
