package chat

import (
	"time"

	"github.com/hinakaze/noro/common"
	"github.com/hinakaze/noro/models/user"

	"github.com/astaxie/beego/orm"
)

type ChatRoom struct {
	Id         int64
	Topic      string
	Creator    *user.User `orm:"rel(fk)"`
	MaxMember  uint16     // <= 0 unlimited
	CreateTime time.Time
}

type TChatRoom struct {
	Id          int64
	Topic       string
	Creator     user.TUser
	MaxMember   uint16
	CreateTime  string
	CreateYear  int
	CreateMonth int
	CreateDay   int
}

func (c *ChatRoom) ToT() (t TChatRoom) {
	t.Id = c.Id
	t.Topic = c.Topic
	t.Creator = c.Creator.ToT(false)
	t.MaxMember = c.MaxMember
	t.CreateTime = common.ToFormatTime(c.CreateTime)
	t.CreateYear = c.CreateTime.Year()
	t.CreateMonth = int(c.CreateTime.Month())
	t.CreateDay = c.CreateTime.Day()
	return
}

/*db*/
func GetRooms() (chatRooms []ChatRoom) {
	_, err := orm.NewOrm().QueryTable("chat_room").RelatedSel().All(&chatRooms)
	if err != nil {
		panic(err.Error())
	}
	return
}

func CreateRoom(topic string, creator *user.User, maxMember int) (room *ChatRoom) {
	room.Topic = topic
	room.Creator = creator
	if maxMember > 100 {
		maxMember = 100
	}
	room.MaxMember = uint16(maxMember)
	room.CreateTime = time.Now()
	room = SaveRoom(room)
	return
}

func SaveRoom(r *ChatRoom) *ChatRoom {
	var err error
	r.Id, err = orm.NewOrm().Insert(r)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func GetRoom(id int64) (room *ChatRoom) {
	room = new(ChatRoom)
	room.Id = id
	err := orm.NewOrm().Read(room)
	if err != nil {
		panic(err.Error())
	}
	_, err = orm.NewOrm().LoadRelated(room, "Creator")
	if err != nil {
		panic(err.Error())
	}
	return
}
