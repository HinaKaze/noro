package chat

import (
	"sync"

	"github.com/hinakaze/noro/models/user"

	"github.com/gorilla/websocket"
)

var ChatRoomMgr *ChatRoomManager

func init() {
	ChatRoomMgr = new(ChatRoomManager)
	ChatRoomMgr.Init()
	ChatRoomMgr.RoomMap = make(map[int64]*RChatRoom)
}

/*runtime*/
type ChatRoomManager struct {
	RoomMap map[int64]*RChatRoom
	rwmutex *sync.RWMutex
}

func (c *ChatRoomManager) Init() {
	c.rwmutex = new(sync.RWMutex)
}

func (c *ChatRoomManager) AddRoom(room ChatRoom) *RChatRoom {
	c.rwmutex.Lock()
	defer c.rwmutex.Unlock()
	var newRRoom = RChatRoom{ChatRoom: room}
	newRRoom.Init()
	c.RoomMap[room.Id] = &newRRoom
	return &newRRoom
}

func (c *ChatRoomManager) addRoom(room ChatRoom) *RChatRoom {
	var newRRoom = RChatRoom{ChatRoom: room}
	newRRoom.Init()
	c.RoomMap[room.Id] = &newRRoom
	return &newRRoom
}

func (c *ChatRoomManager) GetRoom(roomId int64) (detail *RChatRoom, ok bool) {
	c.rwmutex.RLock()
	defer c.rwmutex.RUnlock()
	if detail, ok = c.RoomMap[roomId]; ok {
		return
	} else {
		if room := GetRoom(roomId); room != nil {
			detail = c.addRoom(*room)
			return detail, true
		} else {
			return nil, false
		}
	}
}

type RChatRoom struct {
	ChatRoom
	Mates       []*user.RUser
	HistoryMsgs []ChatMessage
	sync.RWMutex
}

func (c *RChatRoom) Init() {
	c.Mates = make([]*user.RUser, 0)
	c.HistoryMsgs = make([]ChatMessage, 0)
}

func (c *RChatRoom) AddMate(u user.User, ws *websocket.Conn) bool {
	c.Lock()
	defer c.Unlock()
	//	if c.MaxMember <= uint16(len(c.Mates)) {
	//		return false
	//	}
	newRUser := &user.RUser{User: u, WS: ws}
	//	for _, ou := range c.Mates {
	//		if ou.Id == u.Id {
	//			ou = newUserDetail
	//			return true
	//		}
	//	}
	c.Mates = append(c.Mates, newRUser)
	return true
}

func (c *RChatRoom) RemoveMate(uId int64) {
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

func (c *RChatRoom) BroadcastMessage(m ChatMessage) {
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
		mate.WS.WriteJSON(tm)
	}
}

type TRChatRoom struct {
	TChatRoom
	Mates       []user.TUser
	HistoryMsgs []TChatMessage
	Myself      user.TUser
}

func (c *RChatRoom) ToT(myself user.User) (t TRChatRoom) {
	t.TChatRoom = c.ChatRoom.ToT()
	t.Mates = make([]user.TUser, 0)
	for _, u := range c.Mates {
		t.Mates = append(t.Mates, u.ToT(false))
	}
	t.HistoryMsgs = make([]TChatMessage, 0)
	for _, c := range c.HistoryMsgs {
		t.HistoryMsgs = append(t.HistoryMsgs, c.ToT())
	}
	t.Myself = myself.ToT(false)
	return
}
