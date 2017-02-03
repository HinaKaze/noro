package main

import (
	"flag"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hinakaze/iniparser"
	"github.com/hinakaze/noro/models"
	_ "github.com/hinakaze/noro/routers"
)

func main() {
	//	dbFlag := flag.Bool("db", false, "true/false :enable/disable db")
	flag.Parse()
	/*
		db init
	*/
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
	orm.DefaultTimeLoc = time.Local
	orm.RegisterModel(new(models.Friendship), new(models.User), new(models.ChatRoom), new(models.ChatMessage))
	orm.RegisterModel(new(models.Show))
	orm.RunSyncdb("default", false, true)

	models.UserRobot = models.GetUserByName("Noro")
	//	models.SaveShow(models.Show{User: &models.User{Id: 3}, Body: 1, Hair: 1, Emotion: 1, Clothes: 1, Trousers: 1, Shoes: 1})
	fakeUser := models.CreateUser("nigger", "nigger", 1)
	models.SaveUser(fakeUser)
	log.Println(models.GetUser(3).Show)

	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.Run()
}
