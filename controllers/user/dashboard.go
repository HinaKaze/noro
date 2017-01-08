package user

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {
	userp, ok := c.GetSession("user").(*models.User)
	if !ok {
		c.Ctx.WriteString("")
		return
	}
	c.Data["User"] = userp.ToT(true)
	c.TplName = "user/dashboard.html"
}
