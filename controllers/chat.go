package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/hinakaze/noro/models"
)

var wss []*websocket.Conn = make([]*websocket.Conn, 0)

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
	fakeMessage1 := models.ChatMessage{Id: 1, User: fakeUser1, Text: "mission start", Time: time.Now()}
	fakeMessage2 := models.ChatMessage{Id: 2, User: fakeUser2, Text: "mission end", Time: time.Now()}
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
			log.Println(x)
		}
	}()
	ws, err := websocket.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		panic(err.Error())
	}
	defer ws.Close()
	wss = append(wss, ws)
	for {
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			panic(err.Error())
		}
		fakeUserName := fmt.Sprintf("noro%d", len(wss))
		fakeUser := models.User{Id: len(wss), Name: fakeUserName}
		fakeMessage := models.ChatMessage{Id: 1, Type: 2, User: fakeUser, Text: string(bytes), Time: time.Now()}
		for _, c := range wss {
			c.WriteJSON(fakeMessage)
		}
	}
}
