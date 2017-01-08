package routers

import (
	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/controllers"
	"github.com/hinakaze/noro/controllers/chat"
	"github.com/hinakaze/noro/controllers/user"
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
	beego.Router("/chat/ws", &chat.ChatRoomController{}, "get:WS")

	//user
	beego.Router("/user/dashboard", &user.DashboardController{})
	beego.Router("/user/room", &user.RoomController{})
	beego.Router("/user/ws", &user.RoomController{}, "get:WS")
}
