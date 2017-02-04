package user

import (
	"github.com/astaxie/beego"
	muser "github.com/hinakaze/noro/models/user"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {
	userp, ok := c.GetSession("user").(*muser.User)
	if !ok {
		c.Ctx.WriteString("")
		return
	}
	c.Data["User"] = userp.ToT(true)
	c.TplName = "user/dashboard.html"
}
