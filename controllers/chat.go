package controllers

import (
	"runtime/debug"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/hinakaze/noro/models"
)

type ChatRoomsController struct {
	beego.Controller
}

func (c *ChatRoomsController) Get() {
	c.Data["List"] = models.GetRooms()
	c.Data["Title"] = "作戦本部"
	c.TplName = "chat_rooms.tpl"
}

type ChatRoomController struct {
	beego.Controller
}

func (c *ChatRoomController) Get() {
	fakeChatRoomDetail := models.ChatRoomMgr.GetRoomDetail(1)
	c.Data["json"] = &fakeChatRoomDetail
	c.ServeJSON()
}

type ChatMessagesController struct {
	beego.Controller
}

func (c *ChatMessagesController) Get() {
	c.TplName = "chat_messages.tpl"
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
	user := models.GetUser(1)
	ws, err := websocket.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		beego.BeeLogger.Error(err.Error())
		return
	}
	defer func() {
		roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 1, User: user, Text: "", Time: time.Now().String()})
		err := ws.Close()
		if err != nil {
			beego.BeeLogger.Error(err.Error())
			return
		}
	}()
	ws.WriteMessage(websocket.TextMessage, []byte("aa"))
	roomDetail.AddMate(models.GetUser(1), ws)
	roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 0, User: user, Text: "", Time: time.Now().String()})
	for {
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			panic(err.Error())
		}
		roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 2, User: user, Text: string(bytes), Time: time.Now().String()})
	}
}
