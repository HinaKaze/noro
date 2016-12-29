package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
)

type ChatLobbyController struct {
	beego.Controller
}

func (c *ChatLobbyController) Get() {
	c.Data["List"] = models.GetRooms()
	c.TplName = "chat_lobby.html"
}
