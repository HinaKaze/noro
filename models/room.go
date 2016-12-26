package models

import(
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)
type ChatRoomDetail struct{
	ChatRoom
	Mates []User
	wss []*websocket.Conn
	sync.RWMutex
}
func(c *ChatRoomDetail)AddMate(u User){
	RWMutex.Lock()
	defer RWMutex.Unlock()
	if c.MaxMember <= len(c.Mates){
		return
	}
	for _,ou := range c.Mates{
		if ou.Id == u.Id{
			return
		}
	}
	c.Mates = append(c.Mates,u)
}
func(c *ChatRoomDetail)RemoveMate(){
	RWMutex.lock()
	defer REMutex.unlock()
	for i := 0;i < len(c.:g)
}

