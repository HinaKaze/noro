package models

import (
	"time"
)

type ChatRoom struct {
	Id          int
	Name        string
	Creator     User
	CreateTime  time.Time
	MaxMember   uint16
	CreateDay   int
	CreateMonth int
	CreateYear  int
}

type ChatRoomDetail struct {
	ChatRoom
	Mates []User
}

type ChatMessage struct {
	Id   int
	Type uint8 //0 join,1 leave,2 message
	User User
	Text string
	Time string
}
