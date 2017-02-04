package user

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/gorilla/websocket"
)

/*def*/
type User struct {
	Id            int64
	Name          string        `orm:"unique"`
	Friends       []*Friendship `orm:"reverse(many)"`
	HP            int
	MP            int
	San           int
	NPoint        int
	Gender        int // 1 male 0 female
	Password      string
	CanLogin      bool
	LoginSequence string //登陆序列，仅当用户口令修改后更新
	LoginToken    string //登陆token，新的登陆session会更新
	Show          *Show  `orm:"reverse(one)"`
}

func (u *User) AddFriend(friend *User) {
	newFriend := Friendship{User1: u, User2: friend}
	u.Friends = append(u.Friends, &newFriend)
}

func (u *User) GetFriend(targetId int64) *Friendship {
	for _, f := range u.Friends {
		if f.User2.Id == targetId {
			return f
		}
	}
	return nil
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

func (u *User) ToT(flag bool) (t TUser) {
	t.Id = u.Id
	t.Name = u.Name
	t.Gender = u.Gender
	if flag {
		t.Friends = make([]TFriendship, 0)
		for _, f := range u.Friends {
			t.Friends = append(t.Friends, f.ToT())
		}
		t.Show = u.Show.ToT()
	}
	return
}

type RUser struct {
	User
	WS *websocket.Conn
}

type TUser struct {
	Id      int64
	Name    string
	Gender  int
	Friends []TFriendship
	Show    TShow
}

/*logic*/
func CreateUser(name string, password string, gender int) (user *User) {
	user.Name = name
	user.Password = password
	user.CanLogin = true
	user.Gender = gender
	user.GenerateNewLoginSeq()
	user.GenerateNewLoginToken()
	user = SaveUser(user)
	user.Show = &Show{User: user, Body: 1, Hair: 1, Emotion: 1, Clothes: 1, Trousers: 1, Shoes: 1}
	user.Show = SaveShow(user.Show)
	return
}

/*db*/
func SaveUser(user *User) *User {
	var err error
	user.Id, err = orm.NewOrm().Insert(user)
	if err != nil {
		panic(err.Error())
	}
	return user
}

func UpdateUser(user *User) {
	_, err := orm.NewOrm().Update(user)
	if err != nil {
		panic(err.Error())
	}
}

func GetUser(id int64) (user *User) {
	user = new(User)
	user.Id = id
	err := orm.NewOrm().Read(user)
	if err != nil {
		panic(err.Error())
	}
	_, err = orm.NewOrm().LoadRelated(user, "Show")
	if err != nil {
		panic(err.Error())
	}
	_, err = orm.NewOrm().LoadRelated(user, "Friends")
	if err != nil {
		panic(err.Error())
	}
	return
}

func GetUserByName(name string) (user *User) {
	user = new(User)
	err := orm.NewOrm().Raw(`select * from "user" where name=?`, name).QueryRow(user)
	if err != nil {
		panic(err.Error())
	}
	return GetUser(user.Id)
}
