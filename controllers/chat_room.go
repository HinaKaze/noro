package controllers

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/hinakaze/noro/models"
)

type ChatRoomController struct {
	beego.Controller
}

func (c *ChatRoomController) Get() {
	roomId, err := c.GetInt("room_id")
	if err != nil {
		panic(err.Error())
	}
	c.SetSession("room_id", roomId)
	roomDetail := models.ChatRoomMgr.GetRoomDetail(roomId)
	if roomDetail == nil {
		panic(fmt.Sprintf("Room [%d] invalid", roomId))
	}
	c.Data["RoomDetail"] = *roomDetail
	c.TplName = "chat_room.html"
}

type WebSocketController struct {
	beego.Controller
}

func (w *WebSocketController) Join() {
	defer func() {
		if x := recover(); x != nil {
			beego.BeeLogger.Warning("WebSocket disconnected [%+v],%s", x, debug.Stack())
		}
	}()
	roomDetail := models.ChatRoomMgr.GetRoomDetail(1)
	userp, ok := w.GetSession("user").(*models.User)
	user := *userp
	if !ok {
		w.Redirect("/login", 302)
		return
	}
	ws, err := websocket.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		beego.BeeLogger.Error(err.Error())
		return
	}
	defer func() {
		roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 1, User: user, Text: "", Time: time.Now().String()})
		roomDetail.RemoveMate(user.Id)
		err := ws.Close()
		if err != nil {
			beego.BeeLogger.Error(err.Error())
			return
		}
	}()
	roomDetail.AddMate(user, ws)
	roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 0, User: user, Text: "", Time: time.Now().String()})
	for {
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			panic(err.Error())
		}
		roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 2, User: user, Text: string(bytes), Time: time.Now().String()})
	}
}
