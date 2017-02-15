package user

import (
	"sync"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var UserRobot *User

var userRoomId int32 = 0

var userRoomMap map[int64]*RUserRoom = make(map[int64]*RUserRoom)
var userRoomMapMutex *sync.RWMutex = new(sync.RWMutex)
var userRoomMutex sync.RWMutex

func CreateRUserRoom(owner *User) (room *RUserRoom) {
	room = new(RUserRoom)
	room.Id = owner.Id
	room.Owner = owner
	room.Mates = make([]*RUser, 0)
	//	room.HistoryMsgs = make([]ChatMessage, 0)
	return
}

func SaveRUserRoom(room *RUserRoom) {
	userRoomMutex.Lock()
	defer userRoomMutex.Unlock()
	if _, ok := userRoomMap[room.Id]; ok {
		beego.BeeLogger.Warning("User want to create user room,but id duplicated [%d]", room.Id)
		return
	}
	userRoomMap[room.Id] = room
	return
}
func GetRUserRoom(id int64) (room *RUserRoom, ok bool) {
	userRoomMutex.RLock()
	room, ok = userRoomMap[id]
	userRoomMutex.RUnlock()
	if ok {
		return room, true
	} else {
		if u := GetUser(id); u == nil {
			return nil, false
		} else {
			room = CreateRUserRoom(u)
			if room == nil {
				return nil, false
			} else {
				SaveRUserRoom(room)
				return room, true
			}
		}
	}
}

type RUserRoom struct {
	Id          int64
	Owner       *User
	Mates       []*RUser
	HistoryMsgs []RoomMessage
	sync.RWMutex
}

func (c *RUserRoom) Init() {
	c.Mates = make([]*RUser, 0)
	c.HistoryMsgs = make([]RoomMessage, 0)
}

func (c *RUserRoom) AddMate(u *User, ws *websocket.Conn) bool {
	c.Lock()
	defer c.Unlock()
	//	if c.MaxMember <= uint16(len(c.Mates)) {
	//		return false
	//	}
	newRUser := &RUser{User: *u, WS: ws}
	//	for _, ou := range c.Mates {
	//		if ou.Id == u.Id {
	//			ou = newUserDetail
	//			return true
	//		}
	//	}
	c.Mates = append(c.Mates, newRUser)
	return true
}

func (c *RUserRoom) RemoveMate(uId int64) {
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

func (c *RUserRoom) BroadcastMessage(m RoomMessage) {
	c.RLock()
	defer c.RUnlock()
	if m.Type == MessageMsg {
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
	}
	tm := m.ToT()
	for _, mate := range c.Mates {
		if m.User.Id == mate.User.Id {
			continue
		}
		mate.WS.WriteJSON(tm)
	}
}

func (this *RUserRoom) ToT() (t TUserRoom) {
	t.Owner = this.Owner.ToT(false)
	t.HistoryMsgs = make([]TRoomMessage, 0)
	for _, msg := range this.HistoryMsgs {
		t.HistoryMsgs = append(t.HistoryMsgs, msg.ToT())
	}
	t.Mates = make([]TUser, 0)
	for _, mate := range this.Mates {
		t.Mates = append(t.Mates, mate.ToT(false))
	}
	return
}

type TUserRoom struct {
	Owner       TUser
	HistoryMsgs []TRoomMessage
	Mates       []TUser
}
