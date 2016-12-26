package models

import (
	"github.com/gorilla/websocket"
)

type User struct {
	Id   int
	Name string
}

type UserDetail struct {
	User
	ws *websocket.Conn
}
