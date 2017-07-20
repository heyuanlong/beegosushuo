package home

import (
	"github.com/astaxie/beego"
	models "beegosushuo/models"
	utils "beegosushuo/utils"
)
type RegController struct {
	beego.Controller
}

func (c *RegController) Get() {
	c.TplName = "home/reg.tpl"
}

func (this *RegController) Post() {
	account := this.GetString("account")
	nickname := this.GetString("nickname")
	password := this.GetString("password")

	if account == "" || password == "" {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	}
	nickname = "无名_" + nickname

	_, err := models.GetUserByAccout(account)
	_, err_ := models.GetUserByNickname(nickname)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "user account is exist", "refer": "/"}
	} else if err_ == nil{
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "user nickname is exist", "refer": "/"}
	}else {
		passwd := utils.Md5(utils.Md5(password))
		err:=models.AddUser(account,nickname,passwd)
		if err != nil{
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "add user fail", "refer": "/"}
		}else{
			this.SetSession("account", account)
			this.Data["json"] = map[string]interface{}{"result": true, "msg": "add user success", "refer": "/"}
		}
	}
	this.ServeJSON()
}
