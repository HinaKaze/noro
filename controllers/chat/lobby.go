package chat

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
)

type ChatLobbyController struct {
	beego.Controller
}

func (c *ChatLobbyController) Get() {
	tRooms := make([]models.TChatRoom, 0)
	rooms := models.GetRooms()
	for _, r := range rooms {
		tRooms = append(tRooms, r.ToT())
	}
	c.Data["RoomList"] = tRooms
	c.TplName = "chat/lobby.html"
}
