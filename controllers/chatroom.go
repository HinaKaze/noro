package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
)

type ChatRoomListController struct {
	beego.Controller
}

func (c *ChatRoomListController) Get() {
	fakeChatRooms := make([]models.ChatRoom, 0)
	fakeChatRooms = append(fakeChatRooms, models.ChatRoom{Id: 1, Name: "noro作战本部", CreateTime: time.Now(), Creator: models.User{Id: 1, Name: "HinaKaze"}})
	c.Data["List"] = fakeChatRooms
	c.Data["Title"] = "作戦本部"
	c.TplName = "chat_rooms.tpl"
}
