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
	var newDetail = ChatRoomDetail{ChatRoom: room}
	newDetail.Init()
	c.RoomMap[room.Id] = &newDetail
	return
}

func (c *ChatRoomManager) GetRoomDetail(roomId int) (detail *ChatRoomDetail) {
	c.RLock()
	defer c.RUnlock()
	return c.RoomMap[roomId]
}

type ChatRoomDetail struct {
	ChatRoom
	Mates       []*UserDetail
	HistoryMsgs []ChatMessage
	sync.RWMutex
}

func (c *ChatRoomDetail) Init() {
	c.Mates = make([]*UserDetail, 0)
	c.HistoryMsgs = make([]ChatMessage, 0)
}

func (c *ChatRoomDetail) AddMate(u User, ws *websocket.Conn) bool {
	c.Lock()
	defer c.Unlock()
	//	if c.MaxMember <= uint16(len(c.Mates)) {
	//		return false
	//	}
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
	if len(c.HistoryMsgs) >= 15 {
		c.HistoryMsgs = append(c.HistoryMsgs[1:], m)
	} else {
		c.HistoryMsgs = append(c.HistoryMsgs, m)
	}
	index := 1
	for i := range c.HistoryMsgs {
		c.HistoryMsgs[i].Id = index
		index++
	}
	for _, mate := range c.Mates {
		mate.ws.WriteJSON(m)
	}
}
