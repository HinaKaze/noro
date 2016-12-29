package models

import (
	"time"
)

type ChatRoom struct {
	Id          int
	Topic       string
	Creator     User
	MaxMember   uint16 // <= 0 unlimited
	CreateTime  time.Time
	CreateDay   int
	CreateMonth int
	CreateYear  int
}

type ChatMessage struct {
	Id   int
	Type int //0 join,1 leave,2 message
	User User
	Text string
	Time string
}
