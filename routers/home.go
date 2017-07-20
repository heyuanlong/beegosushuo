package routers

import (
	"beegosushuo/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
    
    beego.Router("/home/login", &home.LoginOutController{})
	beego.Router("/home/logout", &home.LoginOutController{},"*:LoginOut")
	beego.Router("/home/reg", &home.RegController{})

	beego.Router("/:id([0-9]+)", &home.MainController{})
	beego.Router("/", &home.MainController{})
    beego.Router("/article/?:id", &home.MainController{},"*:OneArticle")

}
