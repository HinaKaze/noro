package controllers

import (
	"fmt"
	"runtime/debug"
	"sync/atomic"
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

type ChatEnterRoomController struct {
	beego.Controller
}

func (c *ChatEnterRoomController) Get() {
	c.Data["HistoryMsgs"] = models.ChatRoomMgr.GetRoomDetail(1).HistoryMsgs
	c.Data["HistoryMsgLength"] = len(models.ChatRoomMgr.GetRoomDetail(1).HistoryMsgs)
	c.TplName = "chat_room.tpl"
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
	user.Id = getUserId()
	user.Name = fmt.Sprintf("user%d", user.Id)
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

var userId uint32 = 0

func getUserId() int {
	atomic.AddUint32(&userId, 1)
	return int(userId)
}
