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

type ChatMessage struct {
	Id   int
	User User
	Text string
	Time time.Time
}
