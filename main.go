package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hinakaze/noro/routers"
)

func main() {
	//beego.SetStaticPath("static/img", "./static/img")
	beego.Run()
}
