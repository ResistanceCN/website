package controller

import (
	"net/http"
	"strconv"

	"../db"
	"../service"
	"../util"
	"github.com/kataras/iris/context"
)

func LoginPage(ctx context.Context) {
	ctx.ViewData("google_client_id", service.Config.GoogleClientID)
	ctx.View("login.html")
}

func Login(ctx context.Context) {
	idToken := ctx.PostValue("id_token")

	info, err := util.VerifyGoogleIDToken(idToken)

	if err != nil {
		ctx.StatusCode(http.StatusUnauthorized)
		return
	}

	user := db.User{}

	db.Conn().Where(db.User{
		GoogleID: info.SUB,
	}).Attrs(db.User{
		Email: info.Email,
	}).FirstOrInit(&user)

	if user.ID != 0 {
		session := util.GetSession(ctx)
		session.Set("user_id", user.ID)
	}

	ctx.JSON(user)
}

func Register(ctx context.Context) {
	idToken := ctx.PostValue("id_token")
	name := ctx.PostValue("name")

	if !util.ValidateName(name) {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(AjaxResult{
			Msg: "Invalid name",
		})
		return
	}

	faction, err := strconv.Atoi(ctx.PostValue("faction"))

	if err != nil || faction < db.FACTION_UNSPECIFIED || faction > db.FACTION_ENLIGHTENED {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(AjaxResult{
			Msg: "Invalid faction",
		})
		return
	}

	info, err := util.VerifyGoogleIDToken(idToken)

	if err != nil {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(AjaxResult{
			Msg: "Invalid Google account. Please refresh page and try again",
		})
		return
	}

	user := db.User{}

	db.Conn().Where(db.User{
		GoogleID: info.SUB,
	}).Attrs(db.User{
		Email: info.Email,
		Faction: int8(faction),
	}).FirstOrCreate(&user)

	session := util.GetSession(ctx)
	session.Set("user_id", user.ID)

	ctx.JSON(AjaxResult{
		Msg: "Login successful",
		Redirect: "/",
	})
}
