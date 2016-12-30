package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Redirect("/login", 302)
	}
	if user, ok := models.GetUserByName(username); ok {
		if user.CheckPasswork(password) {
			user.GenerateNewLoginToken()
			c.SetSecureCookie("noro_", "_n", user.Name)
			c.SetSecureCookie("noro_", "_s", user.LoginSequence)
			c.SetSecureCookie("noro_", "_t", user.LoginToken)
			c.Redirect("/chat/lobby", 302)
		} else {
			c.Redirect("/login", 302)
		}
	} else {
		c.Redirect("/login", 302)
	}
}
