package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "muyang.work"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
