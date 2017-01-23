package models

import (
	"sync"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

var UserRobot *User

var userRoomId int32 = 0

var userRoomMap map[int64]*UserRoomDetail = make(map[int64]*UserRoomDetail)
var userRoomMapMutex *sync.RWMutex = new(sync.RWMutex)
var userRoomMutex sync.RWMutex

func CreateUserRoomDetail(owner *User) (room *UserRoomDetail) {
	room = new(UserRoomDetail)
	room.Id = owner.Id
	room.Owner = owner
	room.Mates = make([]*UserDetail, 0)
	room.HistoryMsgs = make([]ChatMessage, 0)
	return
}

func SaveUserRoomDetail(room *UserRoomDetail) {
	userRoomMutex.Lock()
	defer userRoomMutex.Unlock()
	if _, ok := userRoomMap[room.Id]; ok {
		beego.BeeLogger.Warning("User want to create user room,but id duplicated [%d]", room.Id)
		return
	}
	userRoomMap[room.Id] = room
	return
}
func GetUserRoomDetail(id int64) (room *UserRoomDetail, ok bool) {
	userRoomMutex.RLock()
	defer userRoomMutex.RUnlock()
	room, ok = userRoomMap[id]
	return
}
