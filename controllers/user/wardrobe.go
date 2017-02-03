package user

import (
	"github.com/astaxie/beego"
)

type WardrobeController struct {
	beego.Controller
}

func (this *WardrobeController) Get() {
	this.TplName = "user/wardrobe.html"
}
