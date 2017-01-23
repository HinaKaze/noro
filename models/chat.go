package models

import (
	"time"

	"github.com/hinakaze/noro/common"
)

type ChatRoom struct {
	Id         int64
	Topic      string
	Creator    *User  `orm:"rel(fk)"`
	MaxMember  uint16 // <= 0 unlimited
	CreateTime time.Time
}

type TChatRoom struct {
	Id          int64
	Topic       string
	Creator     TUser
	MaxMember   uint16
	CreateTime  string
	CreateYear  int
	CreateMonth int
	CreateDay   int
}

func (c *ChatRoom) ToT() (t TChatRoom) {
	t.Id = c.Id
	t.Topic = c.Topic
	t.Creator = c.Creator.ToT(false)
	t.MaxMember = c.MaxMember
	t.CreateTime = common.ToFormatTime(c.CreateTime)
	t.CreateYear = c.CreateTime.Year()
	t.CreateMonth = int(c.CreateTime.Month())
	t.CreateDay = c.CreateTime.Day()
	return
}

type ChatMessage struct {
	Id   int
	Type int   //0 join,1 leave,2 message
	User *User `orm:"rel(fk)"`
	Text string
	Time time.Time
}

type TChatMessage struct {
	Id   int
	Type int
	User TUser
	Text string
	Time string
}

func (c *ChatMessage) ToT() (t TChatMessage) {
	t.Id = c.Id
	t.Type = c.Type
	t.User = c.User.ToT(false)
	t.Text = c.Text
	t.Time = common.ToFormatTime(c.Time)
	return
}
