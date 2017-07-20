package main

import (
	_ "beegosushuo/routers"
	_ "beegosushuo/initial"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

