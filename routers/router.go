package routers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/controllers"
	"github.com/hinakaze/noro/controllers/chat"
	"github.com/hinakaze/noro/controllers/house"
	"github.com/hinakaze/noro/filter"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})

	//chat
	beego.InsertFilter("/chat/*", beego.BeforeRouter, filter.FilterLogin)
	beego.Router("/chat/lobby", &chat.ChatLobbyController{})
	beego.Router("/chat/room", &chat.ChatRoomController{})
	beego.Router("/chat/ws", &chat.WebSocketController{}, "get:Join")

	//house
	beego.Router("/house/game", &house.GameController{})
}
