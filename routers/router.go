package routers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/chat_rooms", &controllers.ChatRoomsController{})
	beego.Router("/chat_room", &controllers.ChatRoomController{})
	beego.Router("/chat_enter_room", &controllers.ChatEnterRoomController{})
	beego.Router("/ws", &controllers.WebSocketController{}, "get:Join")
}
