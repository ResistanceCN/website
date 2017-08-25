package services

import (
	"time"

	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
)

var Sessions *sessions.Sessions

func init() {
	Sessions = sessions.New(sessions.Config{
		Expires: time.Hour * 24 * 7,
	})

	db := redis.New(service.Config{
		Addr: Config.RedisAddr,
		Password: Config.RedisPassword,
		Database: Config.RedisDatabase,
		Prefix: Config.RedisPrefix,
	})

	Sessions.UseDatabase(db)
}
