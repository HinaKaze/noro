package user

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"

	"fmt"
	"runtime/debug"
	"time"

	"github.com/gorilla/websocket"
)

type RoomController struct {
	beego.Controller
}

func (c *RoomController) Get() {
	userId, err := c.GetInt("id")
	if err != nil {
		panic(err.Error())
	}
	user, ok := models.GetUser(userId)
	if !ok {
		panic(fmt.Sprintf("Try to enter a invalid user room [%d]", userId))
	}
	userRoom, ok := models.GetUserRoomDetail(userId)
	if !ok {
		userRoom = models.CreateUserRoomDetail(user)
		models.SaveUserRoomDetail(userRoom)
	}

	myself, ok := c.GetSession("user").(*models.User)
	if !ok {
		c.Redirect("/login", 302)
		return
	}

	c.SetSession("userRoomId", userId)
	c.Data["Room"] = userRoom.ToT()
	c.Data["Myself"] = myself.ToT(false)
	c.TplName = "user/room.html"
}

func (w *RoomController) WS() {
	defer func() {
		if x := recover(); x != nil {
			beego.BeeLogger.Warning("User room WebSocket disconnected [%+v],%s", x, debug.Stack())
		}
	}()

	userRoomId, ok := w.GetSession("userRoomId").(int)
	if !ok {
		//w.Redirect("/lobby", 302)
		panic("user room id is invalid")
	}
	userRoomDetail, ok := models.GetUserRoomDetail(userRoomId)
	if userRoomDetail == nil || !ok {
		//w.Redirect("/lobby", 302)
		panic("user room is invalid")
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
		userRoomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 2, User: user, Text: "", Time: time.Now()})
		userRoomDetail.RemoveMate(user.Id)
		err := ws.Close()
		if err != nil {
			beego.BeeLogger.Error(err.Error())
			return
		}
	}()
	userRoomDetail.AddMate(user, ws)
	userRoomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 1, User: user, Text: "", Time: time.Now()})
	type wsData struct {
		Type int
		Text string
	}
	for {
		var wsdata wsData
		beego.BeeLogger.Info("Get user room ws data [%+v]", wsdata)
		err := ws.ReadJSON(&wsdata)
		if err != nil {
			beego.BeeLogger.Warning(err.Error())
		}
		beego.BeeLogger.Info("Get user room ws data [%+v]", wsdata)
		userRoomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: wsdata.Type, User: user, Text: wsdata.Text, Time: time.Now()})
	}
	w.Ctx.WriteString("Finish")
}
