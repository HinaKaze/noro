package chat

import (
	"encoding/json"
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
	roomId, err := c.GetInt("id")
	if err != nil {
		panic(err.Error())
	}
	c.SetSession("roomId", roomId)

	userp, ok := c.GetSession("user").(*models.User)
	if !ok {
		c.Redirect("/login", 302)
		return
	}
	roomDetail, ok := models.ChatRoomMgr.GetRoomDetail(roomId)
	if roomDetail == nil || !ok {
		panic(fmt.Sprintf("Room [%d] invalid", roomId))
	}
	c.Data["RoomDetail"] = roomDetail.ToT(*userp)
	c.TplName = "chat/room.html"
}

func (c *ChatRoomController) Post() {
	topic := c.GetString("topic")
	maxmember, err := c.GetInt("maxmember")
	if err != nil {
		panic(err.Error())
	}
	if topic == "" {
		c.Redirect("/login", 302)
	}
	if maxmember <= 0 {
		c.Redirect("/login", 302)
	}
	userp, ok := c.GetSession("user").(*models.User)
	if !ok {
		c.Redirect("/login", 302)
		return
	}
	user := *userp
	chatRoom := models.CreateRoom(topic, user, maxmember)
	models.SaveRoom(chatRoom)

	bytes, err := json.Marshal(chatRoom.ToT())
	if err != nil {
		panic(err.Error())
	}
	c.Data["json"] = string(bytes)
	c.ServeJSON()
}

func (w *ChatRoomController) WS() {
	defer func() {
		if x := recover(); x != nil {
			beego.BeeLogger.Warning("WebSocket disconnected [%+v],%s", x, debug.Stack())
		}
	}()

	roomId, ok := w.GetSession("roomId").(int)
	if !ok {
		w.Redirect("/lobby", 302)
	}
	roomDetail, ok := models.ChatRoomMgr.GetRoomDetail(roomId)
	if roomDetail == nil || !ok {
		w.Redirect("/lobby", 302)
	}
	userp, ok := w.GetSession("user").(*models.User)
	if !ok {
		w.Redirect("/login", 302)
		return
	}
	user := *userp

	ws, err := websocket.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		beego.BeeLogger.Error(err.Error())
		return
	}
	defer func() {
		roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 1, User: user, Text: "", Time: time.Now()})
		roomDetail.RemoveMate(user.Id)
		err := ws.Close()
		if err != nil {
			beego.BeeLogger.Error(err.Error())
			return
		}
	}()
	roomDetail.AddMate(user, ws)
	roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 0, User: user, Text: "", Time: time.Now()})
	for {
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			panic(err.Error())
		}
		roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 2, User: user, Text: string(bytes), Time: time.Now()})
	}
	w.Ctx.WriteString("Finish")
}
