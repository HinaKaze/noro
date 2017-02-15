package user

import (
	"time"

	"github.com/hinakaze/noro/common"
)

const (
	MessageNull uint8 = iota
	MessageJoin
	MessageLeft
	MessageMsg
	MessageOPRight
	MessageOPLeft
	MessageOPUp
	MessageOPStop
)

type RoomMessage struct {
	Id   int
	Type uint8 //0 join,1 leave,2 message
	User *User `orm:"rel(fk)"`
	Text string
	Time time.Time
	Room *RUserRoom `orm:"rel(fk)"`
}

type TRoomMessage struct {
	Id   int
	Type uint8
	User TUser
	Text string
	Time string
}

func (c *RoomMessage) ToT() (t TRoomMessage) {
	t.Id = c.Id
	t.Type = c.Type
	t.User = c.User.ToT(false)
	t.Text = c.Text
	t.Time = common.ToFormatTime(c.Time)
	return
}
