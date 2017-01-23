package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
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
	if user := models.GetUserByName(username); user != nil {
		c.Redirect("/login", 302)
	} else {
		newUser := models.CreateUser(username, password, 0)
		models.SaveUser(newUser)
		c.Redirect("/login", 302)
	}
}
