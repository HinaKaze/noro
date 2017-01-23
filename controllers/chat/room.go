package chat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/hinakaze/noro/models"
)

type ChatRoomController struct {
	beego.Controller
}

func (c *ChatRoomController) Get() {
	roomId, err := c.GetInt64("id")
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
			beego.BeeLogger.Warning("WebSocket disconnected [%+v]", x)
		}
	}()

	roomId, ok := w.GetSession("roomId").(int64)
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

	ws, err := websocket.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		beego.BeeLogger.Error(err.Error())
		return
	}
	defer func() {
		roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 1, User: userp, Text: "", Time: time.Now()})
		roomDetail.RemoveMate(userp.Id)
		err := ws.Close()
		if err != nil {
			beego.BeeLogger.Error(err.Error())
			return
		}
	}()
	roomDetail.AddMate(*userp, ws)
	roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 0, User: userp, Text: "", Time: time.Now()})
	for {
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			panic(err.Error())
		}
		roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 2, User: userp, Text: string(bytes), Time: time.Now()})
		//robot
		if roomDetail.Id == 1 {
			answer := GetRobotAnswer(string(bytes))
			roomDetail.BroadcastMessage(models.ChatMessage{Id: 1, Type: 2, User: models.UserRobot, Text: answer, Time: time.Now()})
		}
	}
	w.Ctx.WriteString("Finish")
}

var msg string = `{"key":"bd2fc9d82bab426681a40e6c1393b53d","info":"%s","loc":"noro","userid":"1"}`

//func init() {
//	var ok bool
//	userRobot, ok = models.GetUser(9988)
//	if !ok {
//		panic("User robot not found")
//	}
//}

func GetRobotAnswer(ask string) (answer string) {
	askMsg := fmt.Sprintf(msg, ask)
	resp, err := http.Post("http://www.tuling123.com/openapi/api", "application/json", strings.NewReader(askMsg))
	if err != nil {
		panic(err.Error())
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var Answer struct {
		Code int
		Text string
	}

	err = json.Unmarshal(respBody, &Answer)
	if err != nil {
		log.Println(err)
		return "我被玩坏了。。。"
	}
	return Answer.Text
}
