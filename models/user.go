package models

import (
	"crypto/md5"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	Id            int
	Name          string
	Friends       []*Friendship
	HP            int
	MP            int
	San           int
	NPoint        int
	Password      string
	CanLogin      bool
	LoginSequence string //登陆序列，仅当用户口令修改后更新
	LoginToken    string //登陆token，新的登陆session会更新
}

func (u *User) AddFriend(targetId int) {
	newFriend := Friendship{UserId1: u.Id, UserId2: targetId}
	u.Friends = append(u.Friends, &newFriend)
}

func (u *User) GetFriend(targetId int) *Friendship {
	for _, f := range u.Friends {
		if f.UserId2 == targetId {
			return f
		}
	}
	return nil
}

func (u *User) CheckPasswork(password string) bool {
	return u.CanLogin && (password == u.Password)
}

func (u *User) CheckLoginSeq(seq string) bool {
	return seq == u.LoginSequence
}

func (u *User) CheckLoginToken(seq string) bool {
	return seq == u.LoginToken
}

func (u *User) GenerateNewLoginSeq() {
	m5 := md5.New()
	u.LoginSequence = fmt.Sprintf("%x", m5.Sum([]byte(string(time.Now().Unix())+u.Name)))
}

func (u *User) GenerateNewLoginToken() {
	m5 := md5.New()
	u.LoginToken = fmt.Sprintf("%x", m5.Sum([]byte(string(time.Now().Unix())+u.Name)))
}

func (u *User) ToT(friendFlag bool) (t TUser) {
	t.Id = u.Id
	t.Name = u.Name
	if friendFlag {
		t.Friends = make([]TFriendship, 0)
		for _, f := range u.Friends {
			t.Friends = append(t.Friends, f.ToT())
		}
	}
	return
}

type UserDetail struct {
	User
	ws *websocket.Conn
}

type TUser struct {
	Id      int
	Name    string
	Friends []TFriendship
}

type UserRoomDetail struct {
	Id          int
	Owner       *User
	Mates       []*UserDetail
	HistoryMsgs []ChatMessage
	sync.RWMutex
}

func (c *UserRoomDetail) Init() {
	c.Mates = make([]*UserDetail, 0)
	c.HistoryMsgs = make([]ChatMessage, 0)
}

func (c *UserRoomDetail) AddMate(u User, ws *websocket.Conn) bool {
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

func (c *UserRoomDetail) RemoveMate(uId int) {
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

func (c *UserRoomDetail) BroadcastMessage(m ChatMessage) {
	c.RLock()
	defer c.RUnlock()
	if m.Type == 3 {
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
		mate.ws.WriteJSON(tm)
	}
}

func (this *UserRoomDetail) ToT() (t TUserRoom) {
	t.Owner = this.Owner.ToT(false)
	t.HistoryMsgs = make([]TChatMessage, 0)
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
	HistoryMsgs []TChatMessage
	Mates       []TUser
}
