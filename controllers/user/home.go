package user

import (
	"log"

	"github.com/astaxie/beego"
	muser "github.com/hinakaze/noro/models/user"

	"fmt"
	"runtime/debug"
	"time"

	"github.com/gorilla/websocket"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	userId, err := c.GetInt64("id")
	if err != nil {
		log.Println(err.Error())
	}
	user := muser.GetUser(userId)
	if user == nil {
		log.Println(fmt.Sprintf("Try to enter a invalid user home [%d]", userId))
	}
	userRoom, ok := muser.GetRUserRoom(userId)
	if !ok {
		userRoom = muser.CreateRUserRoom(user)
		muser.SaveRUserRoom(userRoom)
	}

	myself, ok := c.GetSession("user").(*muser.User)
	if !ok {
		c.Redirect("/login", 302)
		return
	}

	c.SetSession("userRoomId", userId)
	c.Data["Room"] = userRoom.ToT()
	c.Data["Myself"] = myself.ToT(false)
	c.TplName = "user/home.html"
}

func (w *HomeController) WS() {
	defer func() {
		if x := recover(); x != nil {
			beego.BeeLogger.Warning("User room WebSocket disconnected [%+v],%s", x, debug.Stack())
		}
	}()

	userRoomId, ok := w.GetSession("userRoomId").(int64)
	if !ok {
		//w.Redirect("/lobby", 302)
		panic("user room id is invalid")
	}
	rUserRoom, ok := muser.GetRUserRoom(userRoomId)
	if rUserRoom == nil || !ok {
		//w.Redirect("/lobby", 302)
		panic("user room is invalid")
	}
	user, ok := w.GetSession("user").(*muser.User)
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
		rUserRoom.BroadcastMessage(muser.RoomMessage{Id: 1, Type: muser.MessageLeft, User: user, Text: "", Time: time.Now()})
		rUserRoom.RemoveMate(user.Id)
		err := ws.Close()
		if err != nil {
			beego.BeeLogger.Error(err.Error())
			return
		}
	}()
	rUserRoom.AddMate(user, ws)
	rUserRoom.BroadcastMessage(muser.RoomMessage{Id: 1, Type: muser.MessageJoin, User: user, Text: "", Time: time.Now()})
	type wsData struct {
		Type uint8
		Text string
	}
	for {
		var wsdata wsData
		err := ws.ReadJSON(&wsdata)
		if err != nil {
			beego.BeeLogger.Warning(err.Error())
		}
		beego.BeeLogger.Info("Get user room ws data [%+v]", wsdata)
		rUserRoom.BroadcastMessage(muser.RoomMessage{Id: 1, Type: wsdata.Type, User: user, Text: wsdata.Text, Time: time.Now()})
	}
	w.Ctx.WriteString("Finish")
}
