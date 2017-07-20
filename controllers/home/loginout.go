package home

import (
	"github.com/astaxie/beego"
	. "beegosushuo/models"
	utils "beegosushuo/utils"
)
type LoginOutController struct {
	beego.Controller
}

func (c *LoginOutController) Get() {
	c.TplName = "home/login.tpl"
}

func (this *LoginOutController) Post() {
	account := this.GetString("account")
	password := this.GetString("password")

	if account == "" || password == "" {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	}

	user, err := GetUserByAccout(account)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "user does not exist", "refer": "/"}
	} else {
		passwd := utils.Md5(utils.Md5(password))

		if passwd == user.Password {
			this.SetSession("account", account)
			this.Data["json"] = map[string]interface{}{"result": true, "msg": "user[" + user.Account + "] login success ", "refer": "/admin/admin"}
		} else {
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "login failed ", "refer": "/"}
		}
	}
	this.ServeJSON()
}
func (this *LoginOutController) LoginOut() {
	this.DelSession("account")
	this.Redirect("/", 302)
}

