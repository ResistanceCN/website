package util

import (
	"regexp"

	"../db"
	"github.com/kataras/iris"
)

func GetUser(ctx iris.Context) db.User {
	cached := ctx.Values().Get("user")

	if cached != nil {
		return cached.(db.User)
	}

	user := db.User{}

	session := GetSession(ctx)
	id, err := session.GetInt("user_id")

	if err == nil {
		db.Conn().First(&user, id)
	}

	if user.ID != 0 {
		ctx.Values().Set("user", user)
	}

	return user
}

var nameRegex = regexp.MustCompile(`^[0-9A-Za-z]{3,16}$`)

func ValidateName(name string) bool {
	return nameRegex.MatchString(name)
}
