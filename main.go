package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/hinakaze/noro/models"
	_ "github.com/hinakaze/noro/routers"
)

func main() {
	//beego.SetStaticPath("static/img", "./static/img")
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	var FilterLogin = func(ctx *context.Context) {
		if user, ok := IsLogin(ctx); ok {
			ctx.Output.Session("user", user)
		} else {
			ctx.Redirect(302, "/login")
		}
	}
	beego.InsertFilter("/chat_lobby", beego.BeforeRouter, FilterLogin)
	beego.Run()
}

func IsLogin(ctx *context.Context) (*models.User, bool) {
	if username, ok := ctx.GetSecureCookie("noro_", "_n"); ok {
		if user, ok := models.GetUserByName(username); ok {
			loginToken, ok := ctx.GetSecureCookie("noro_", "_t")
			if !ok {
				return nil, false
			}
			loginSeq, ok := ctx.GetSecureCookie("noro_", "_s")
			if !ok {
				return nil, false
			}
			if user.CheckLoginSeq(loginSeq) && user.CheckLoginToken(loginToken) {
				return user, true
			} else {
				return nil, false
			}
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}
