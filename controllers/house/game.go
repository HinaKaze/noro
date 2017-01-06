package house

import (
	"github.com/astaxie/beego"
)

type GameController struct {
	beego.Controller
}

func (c *GameController) Get() {
	c.TplName = "house/game.html"
}
