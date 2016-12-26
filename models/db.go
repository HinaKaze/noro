package models

import (
	"sync"
	"time"

	"github.com/astaxie/beego"
)

func init() {
	fakeUser1 := User{Id: 1, Name: "HinaKaze"}
	fakeUser2 := User{Id: 2, Name: "Smilok"}
	CreateUser(fakeUser1)
	CreateUser(fakeUser2)
	fakeChatRoom1 := ChatRoom{Id: 1, Name: "noro作战本部1", CreateTime: time.Now(), CreateDay: time.Now().Day(), CreateMonth: int(time.Now().Month()), CreateYear: time.Now().Year(), Creator: fakeUser1}
	fakeChatRoom2 := ChatRoom{Id: 2, Name: "noro作战本部2", CreateTime: time.Now(), CreateDay: time.Now().Day(), CreateMonth: int(time.Now().Month()), CreateYear: time.Now().Year(), Creator: fakeUser2}
	CreateRoom(fakeChatRoom1)
	CreateRoom(fakeChatRoom2)
}

var chatRoomMap map[int]*ChatRoom = make(map[int]*ChatRoom)
var chatRoomMapMutex sync.RWMutex

func CreateRoom(room ChatRoom) {
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
	return *chatRoomMap[rId]
}

var userMap map[int]*User = make(map[int]*User)
var userMutex sync.RWMutex

func CreateUser(user User) {
	userMutex.Lock()
	defer userMutex.Unlock()
	if _, ok := userMap[user.Id]; ok {
		beego.BeeLogger.Warning("User want to create user,but id duplicated [%d]", user.Id)
		return
	}
	userMap[user.Id] = &user
	return
}

func GetUser(uId int) User {
	userMutex.RLock()
	defer userMutex.RUnlock()
	return *userMap[uId]
}
