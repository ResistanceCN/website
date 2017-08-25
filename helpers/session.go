package helpers

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

func GetSession(ctx context.Context) *sessions.Session {
	return ctx.Values().Get("session").(*sessions.Session)
}
