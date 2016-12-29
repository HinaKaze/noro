package routers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/chat_lobby", &controllers.ChatLobbyController{})
	beego.Router("/chat_room", &controllers.ChatRoomController{})
	beego.Router("/ws", &controllers.WebSocketController{}, "get:Join")
}
