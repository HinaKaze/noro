package models

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	Id            int
	Name          string
	Password      string
	CanLogin      bool
	LoginSequence string //登陆序列，仅当用户口令修改后更新
	LoginToken    string //登陆token，新的登陆session会更新
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

type UserDetail struct {
	User
	ws *websocket.Conn
}
