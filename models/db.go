package models

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hinakaze/iniparser"
	_ "github.com/lib/pq"
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

func Init(dbFlag bool) {

	if dbFlag {
		iniparser.DefaultParse("./conf/user.ini")
		section, ok := iniparser.GetSection("DB")
		if !ok {
			panic("ini parse error")
		}
		driverName, ok := section.GetValue("driverName")
		if !ok {
			panic("[driverName] not found")
		}
		dataSource, ok := section.GetValue("dataSource")
		if !ok {
			panic("[dataSource] not found")
		}

		orm.Debug = true
		orm.RegisterDataBase("default", driverName, dataSource)
		//orm.RegisterModel(new(User))
		//orm.RunSyncdb("default", false, true)
	}
}

var roomId int32 = 0

var chatRoomMap map[int]*ChatRoom = make(map[int]*ChatRoom)
var chatRoomMapMutex *sync.RWMutex = new(sync.RWMutex)

func CreateRoom(topic string, creator User, maxMember int) (room ChatRoom) {
	room.Id = int(atomic.AddInt32(&roomId, 1))
	room.Topic = topic
	room.Creator = creator
	if maxMember > 100 {
		maxMember = 100
	}
	room.MaxMember = uint16(maxMember)
	room.CreateTime = time.Now()
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

func GetRoom(rId int) (room *ChatRoom, ok bool) {
	chatRoomMapMutex.RLock()
	defer chatRoomMapMutex.RUnlock()
	room, ok = chatRoomMap[rId]
	return
}

var userRoomId int32 = 0

var userRoomMap map[int]*UserRoomDetail = make(map[int]*UserRoomDetail)
var userRoomMapMutex *sync.RWMutex = new(sync.RWMutex)

func CreateUserRoomDetail(owner *User) (room *UserRoomDetail) {
	room = new(UserRoomDetail)
	room.Id = owner.Id
	room.Owner = owner
	room.Mates = make([]*UserDetail, 0)
	room.HistoryMsgs = make([]ChatMessage, 0)
	return
}

func SaveUserRoomDetail(room *UserRoomDetail) {
	chatRoomMapMutex.Lock()
	defer chatRoomMapMutex.Unlock()
	if _, ok := userRoomMap[room.Id]; ok {
		beego.BeeLogger.Warning("User want to create user room,but id duplicated [%d]", room.Id)
		return
	}
	userRoomMap[room.Id] = room
	return
}
func GetUserRoomDetail(id int) (room *UserRoomDetail, ok bool) {
	userRoomMapMutex.RLock()
	defer userRoomMapMutex.RUnlock()
	room, ok = userRoomMap[id]
	return
}

var userMap map[int]*User = make(map[int]*User)
var userMapByName map[string]*User = make(map[string]*User)
var userMutex *sync.RWMutex = new(sync.RWMutex)
var userId int32

func CreateUser(name string, password string) (user User) {
	user.Id = int(atomic.AddInt32(&userId, 1))
	user.Name = name
	user.Password = password
	user.CanLogin = true
	user.AddFriend(3)
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
