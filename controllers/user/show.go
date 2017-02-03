package user

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
)

type ShowController struct {
	beego.Controller
}

func (this *ShowController) Get() {
	myself, ok := this.GetSession("user").(*models.User)
	if !ok {
		this.Redirect("/login", 302)
		return
	}
	this.Data["Show"] = myself.ToT(false).Show
	this.TplName = "user/show.html"
}
