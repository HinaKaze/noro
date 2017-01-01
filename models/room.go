package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

var ChatRoomMgr *ChatRoomManager

func init() {
	ChatRoomMgr = new(ChatRoomManager)
	ChatRoomMgr.Init()
	ChatRoomMgr.RoomMap = make(map[int]*ChatRoomDetail)
	if room, ok := GetRoom(1); ok {
		ChatRoomMgr.AddRoomDetail(*room)
	}
}

type ChatRoomManager struct {
	RoomMap map[int]*ChatRoomDetail
	rwmutex *sync.RWMutex
}

func (c *ChatRoomManager) Init() {
	c.rwmutex = new(sync.RWMutex)
}

func (c *ChatRoomManager) AddRoomDetail(room ChatRoom) *ChatRoomDetail {
	c.rwmutex.Lock()
	defer c.rwmutex.Unlock()
	var newDetail = ChatRoomDetail{ChatRoom: room}
	newDetail.Init()
	c.RoomMap[room.Id] = &newDetail
	return &newDetail
}

func (c *ChatRoomManager) addRoomDetail(room ChatRoom) *ChatRoomDetail {
	var newDetail = ChatRoomDetail{ChatRoom: room}
	newDetail.Init()
	c.RoomMap[room.Id] = &newDetail
	return &newDetail
}

func (c *ChatRoomManager) GetRoomDetail(roomId int) (detail *ChatRoomDetail, ok bool) {
	c.rwmutex.RLock()
	defer c.rwmutex.RUnlock()
	if detail, ok = c.RoomMap[roomId]; ok {
		return
	} else {
		if room, ok := GetRoom(roomId); ok {
			detail = c.addRoomDetail(*room)
			return detail, true
		} else {
			return nil, false
		}
	}
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
	tm := m.ToT()
	for _, mate := range c.Mates {
		mate.ws.WriteJSON(tm)
	}
}

type TChatRoomDetail struct {
	TChatRoom
	Mates       []TUser
	HistoryMsgs []TChatMessage
	Myself      TUser
}

func (c *ChatRoomDetail) ToT(myself User) (t TChatRoomDetail) {
	t.TChatRoom = c.ChatRoom.ToT()
	t.Mates = make([]TUser, 0)
	for _, u := range c.Mates {
		t.Mates = append(t.Mates, u.ToT())
	}
	t.HistoryMsgs = make([]TChatMessage, 0)
	for _, c := range c.HistoryMsgs {
		t.HistoryMsgs = append(t.HistoryMsgs, c.ToT())
	}
	t.Myself = myself.ToT()
	return
}
