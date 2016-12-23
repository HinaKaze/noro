package controllers

import (
	"fmt"
	//"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/hinakaze/noro/models"
)

var wss []*websocket.Conn = make([]*websocket.Conn, 0)

//var wssM sync.Mutex

type ChatRoomsController struct {
	beego.Controller
}

func (c *ChatRoomsController) Get() {
	fakeChatRooms := make([]models.ChatRoom, 0)
	fakeChatRoom1 := models.ChatRoom{Id: 1, Name: "noro作战本部1", CreateTime: time.Now(), CreateDay: time.Now().Day(), CreateMonth: int(time.Now().Month()), CreateYear: time.Now().Year(), Creator: models.User{Id: 1, Name: "HinaKaze"}}
	fakeChatRoom2 := models.ChatRoom{Id: 2, Name: "noro作战本部2", CreateTime: time.Now(), CreateDay: time.Now().Day(), CreateMonth: int(time.Now().Month()), CreateYear: time.Now().Year(), Creator: models.User{Id: 2, Name: "Smilok"}}
	fakeChatRooms = append(fakeChatRooms, fakeChatRoom1, fakeChatRoom2)
	c.Data["List"] = fakeChatRooms
	c.Data["Title"] = "作戦本部"
	c.TplName = "chat_rooms.tpl"
}

type ChatMessagesController struct {
	beego.Controller
}

func (c *ChatMessagesController) Get() {
	fakeMessages := make([]models.ChatMessage, 0)
	fakeUser1 := models.User{Id: 1, Name: "HinaKaze"}
	fakeUser2 := models.User{Id: 2, Name: "Smilok"}
	fakeMessage1 := models.ChatMessage{Id: 1, User: fakeUser1, Text: "mission start", Time: time.Now().String()}
	fakeMessage2 := models.ChatMessage{Id: 2, User: fakeUser2, Text: "mission end", Time: time.Now().String()}
	fakeMessages = append(fakeMessages, fakeMessage1, fakeMessage2)
	c.Data["List"] = fakeMessages
	c.TplName = "chat_messages.tpl"
}

type WebSocketController struct {
	beego.Controller
}

func (w *WebSocketController) Join() {
	defer func() {
		if x := recover(); x != nil {
			beego.BeeLogger.Warning("WebSocket disconnected [%+v]", x)
		}
	}()
	fakeUserName := fmt.Sprintf("noro%d", len(wss)+1)
	fakeUser := models.User{Id: len(wss) + 1, Name: fakeUserName}
	ws, err := websocket.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		beego.BeeLogger.Error(err.Error())
		return
	}
	defer func() {
		//		wssM.Lock()
		//		defer wssM.Unlock()
		for i, w := range wss {
			if w == ws {
				wss = append(wss[:i], wss[i+1:]...)
				break
			}
		}
		fakeLeaveMessage := models.ChatMessage{Id: 1, Type: 1, User: fakeUser, Text: "", Time: time.Now().String()}
		broadCastMessage(fakeLeaveMessage)
		err := ws.Close()
		if err != nil {
			beego.BeeLogger.Error(err.Error())
			return
		}
	}()
	wss = append(wss, ws)
	fakeJoinMessage := models.ChatMessage{Id: 1, Type: 0, User: fakeUser, Text: "", Time: time.Now().String()}
	broadCastMessage(fakeJoinMessage)
	for {
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			panic(err.Error())
		}
		fakeMessage := models.ChatMessage{Id: 1, Type: 2, User: fakeUser, Text: string(bytes), Time: time.Now().String()}
		broadCastMessage(fakeMessage)
	}
}

func broadCastMessage(m models.ChatMessage) {
	//wssM.Lock()
	//defer wssM.Unlock()
	beego.BeeLogger.Debug("WebSocket conn count [%d]", len(wss))
	for _, c := range wss {
		c.WriteJSON(m)
	}
}
