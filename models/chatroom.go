package models

import (
	"time"
)

type ChatRoom struct {
	Id         int
	Name       string
	Creator    User
	CreateTime time.Time
	MaxMember  uint16
}
