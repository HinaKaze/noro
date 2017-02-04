package controllers

import (
	"github.com/astaxie/beego"
	muser "github.com/hinakaze/noro/models/user"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Redirect("/login", 302)
	}
	if user := muser.GetUserByName(username); user != nil {
		c.Redirect("/login", 302)
	} else {
		muser.CreateUser(username, password, 0)
		c.Redirect("/login", 302)
	}
}
