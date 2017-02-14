package test

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/hinakaze/iniparser"
	"github.com/hinakaze/noro/models/user"
	_ "github.com/lib/pq"
)

type TUser struct {
	Id      int
	Name    string
	Profile *TProfile `orm:"rel(one)"`      // OneToOne relation
	Post    []*TPost  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type TProfile struct {
	Id   int
	Age  int16
	User *TUser `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type TPost struct {
	Id    int
	Title string
	User  *TUser  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*TTag `orm:"rel(m2m)"`
}

type TTag struct {
	Id    int
	Name  string
	Posts []*TPost `orm:"reverse(many)"`
}

func TestMain(m *testing.M) {
	fmt.Println("Start beego orm test")
	iniparser.DefaultParse("../conf/user.ini")
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
	orm.RegisterModel(new(TUser), new(TProfile), new(TPost), new(TTag))
	orm.RegisterModel(new(user.Show), new(user.User), new(user.Friendship))
	orm.RunSyncdb("default", false, true)
	m.Run()
}

//func TestInsert(t *testing.T) {
//	o := orm.NewOrm()
//	p := TProfile{Age: 35}
//	u := TUser{Name: "noro", Profile: &p}
//	p1 := TPost{Title: "p1", User: &u}
//	p2 := TPost{Title: "p2", User: &u}

//	pId, _ := o.Insert(&p)
//	t.Log(pId, p.Id)
//	uId, _ := o.Insert(&u)
//	t.Log(uId, u.Id)
//	o.Insert(&p1)
//	o.Insert(&p2)
//}

//func TestRead(t *testing.T) {
//	o := orm.NewOrm()
//	u := TUser{Id: 4}
//	o.Read(&u)

//	o.LoadRelated(&u, "Post")
//	o.Read(u.Profile)
//	t.Log(u)
//	t.Log(u.Post[0].Title)
//	t.Log(u.Profile)
//}

func TestInjectUserShow(t *testing.T) {
	//	o := orm.NewOrm()
	fakeShow := &user.Show{User: &user.User{Id: 3}, Body: 1, Hair: 1, Emotion: 1, Clothes: 1, Trousers: 1, Shoes: 1}
	user.SaveShow(fakeShow)
}

//func TestRead2(t *testing.T) {
//	o := orm.NewOrm()
//	p := TProfile{}
//	o.QueryTable("t_profile").Filter("User__Id", 4).RelatedSel().One(&p)
//	t.Log(p)
//	t.Log(p.User)
//}
