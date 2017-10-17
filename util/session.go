package util

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

func GetSession(ctx iris.Context) *sessions.Session {
	return ctx.Values().Get("session").(*sessions.Session)
}
