package main

import (
	"flag"

	"github.com/astaxie/beego"
	"github.com/hinakaze/noro/models"
	_ "github.com/hinakaze/noro/routers"
)

func main() {
	dbFlag := flag.Bool("db", false, "true/false :enable/disable db")
	flag.Parse()
	models.Init(*dbFlag)

	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.Run()
}
