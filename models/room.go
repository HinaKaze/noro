package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

var ChatRoomMgr *ChatRoomManager

func init() {
	ChatRoomMgr = new(ChatRoomManager)
	ChatRoomMgr.RoomMap = make(map[int]*ChatRoomDetail)
	ChatRoomMgr.AddRoomDetail(GetRoom(1))
}

type ChatRoomManager struct {
	RoomMap map[int]*ChatRoomDetail
	sync.RWMutex
}

func (c *ChatRoomManager) AddRoomDetail(room ChatRoom) {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.RoomMap[room.Id]; ok {
		return
	}
	c.RoomMap[room.Id] = &ChatRoomDetail{ChatRoom: room}
	return
}

func (c *ChatRoomManager) GetRoomDetail(roomId int) (detail *ChatRoomDetail) {
	c.RLock()
	defer c.RUnlock()
	return c.RoomMap[roomId]
}

type ChatRoomDetail struct {
	ChatRoom
	Mates []*UserDetail
	sync.RWMutex
}

func (c *ChatRoomDetail) AddMate(u User, ws *websocket.Conn) bool {
	c.Lock()
	defer c.Unlock()
	if c.MaxMember <= uint16(len(c.Mates)) {
		return false
	}
	newUserDetail := &UserDetail{User: u, ws: ws}
	//	for _, ou := range c.Mates {
	//		if ou.Id == u.Id {
	//			ou = newUserDetail
	//			return true
	//		}
	//	}
	c.Mates = append(c.Mates, newUserDetail)
	return true
}

func (c *ChatRoomDetail) RemoveMate(uId int) {
	c.Lock()
	defer c.Unlock()
	for i := range c.Mates {
		if c.Mates[i].Id == uId {
			c.Mates = append(c.Mates[:i], c.Mates[i+1:]...)
			break
		}
	}
	return
}

func (c *ChatRoomDetail) BroadcastMessage(m ChatMessage) {
	c.RLock()
	defer c.RUnlock()
	for _, m := range c.Mates {
		m.ws.WriteJSON(m)
	}
}
