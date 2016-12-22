package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
)

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
