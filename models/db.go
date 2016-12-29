package models

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/astaxie/beego"
)

func init() {
	fakeUser0 := CreateUser("God", "")
	fakeUser1 := CreateUser("HinaKaze", "HinaKaze")
	fakeUser2 := CreateUser("Smilok", "Smilok")
	SaveUser(fakeUser1)
	SaveUser(fakeUser2)
	fakeChatRoom := CreateRoom("Noro", fakeUser0, 0)
	SaveRoom(fakeChatRoom)
}

var roomId int32 = 0

var chatRoomMap map[int]*ChatRoom = make(map[int]*ChatRoom)
var chatRoomMapMutex sync.RWMutex

func CreateRoom(topic string, creator User, maxMember int) (room ChatRoom) {
	room.Id = int(atomic.AddInt32(&roomId, 1))
	room.Topic = topic
	room.Creator = creator
	if maxMember > 100 {
		maxMember = 100
	}
	room.MaxMember = uint16(maxMember)
	room.CreateTime = time.Now()
	room.CreateDay = room.CreateTime.Day()
	room.CreateMonth = int(room.CreateTime.Month())
	room.CreateYear = room.CreateTime.Year()
	return
}

func SaveRoom(room ChatRoom) {
	chatRoomMapMutex.Lock()
	defer chatRoomMapMutex.Unlock()
	if _, ok := chatRoomMap[room.Id]; ok {
		beego.BeeLogger.Warning("User want to create room,but id duplicated [%d]", room.Id)
		return
	}
	chatRoomMap[room.Id] = &room
	return
}

func GetRooms() (chatRooms []ChatRoom) {
	chatRoomMapMutex.RLock()
	defer chatRoomMapMutex.RUnlock()
	for _, r := range chatRoomMap {
		chatRooms = append(chatRooms, *r)
	}
	return
}

func GetRoom(rId int) ChatRoom {
	chatRoomMapMutex.RLock()
	defer chatRoomMapMutex.RUnlock()
	fmt.Println(chatRoomMap == nil)
	return *chatRoomMap[rId]
}

var userMap map[int]*User = make(map[int]*User)
var userMapByName map[string]*User = make(map[string]*User)
var userMutex sync.RWMutex
var userId int32

func CreateUser(name string, password string) (user User) {
	user.Id = int(atomic.AddInt32(&userId, 1))
	user.Name = name
	user.Password = password
	user.CanLogin = true
	return
}

func SaveUser(user User) {
	userMutex.Lock()
	defer userMutex.Unlock()
	if _, ok := userMap[user.Id]; ok {
		beego.BeeLogger.Warning("User want to create user,but id duplicated [%d]", user.Id)
		return
	}
	userMap[user.Id] = &user
	userMapByName[user.Name] = &user
	return
}

func GetUser(uId int) (*User, bool) {
	userMutex.RLock()
	defer userMutex.RUnlock()
	up, ok := userMap[uId]
	return up, ok
}

func GetUserByName(name string) (*User, bool) {
	userMutex.RLock()
	defer userMutex.RUnlock()
	up, ok := userMapByName[name]
	return up, ok
}
