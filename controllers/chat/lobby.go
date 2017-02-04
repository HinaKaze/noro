package chat

import (
	"github.com/astaxie/beego"
	mchat "github.com/hinakaze/noro/models/chat"
)

type ChatLobbyController struct {
	beego.Controller
}

func (c *ChatLobbyController) Get() {
	tRooms := make([]mchat.TChatRoom, 0)
	rooms := mchat.GetRooms()
	for _, r := range rooms {
		tRooms = append(tRooms, r.ToT())
	}
	c.Data["RoomList"] = tRooms
	c.TplName = "chat/lobby.html"
}
