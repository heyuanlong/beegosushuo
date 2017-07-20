package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       		int
	Account    		string
	Nickname  		string
	Password     	string
	Time    		int
}

func init() {
	orm.RegisterModel(new(User))
}
func GetUserById(id int) (User, error) {
	var err error
	var art User

	o := orm.NewOrm()
	o.Using("default")
	art = User{Id: id}
	err = o.Read(&art, "id")

	return art, err
}
func GetUserByAccout(account string) (User, error) {
	var err error
	var art User

	o := orm.NewOrm()
	o.Using("default")
	art = User{Account: account}
	err = o.Read(&art, "account")

	return art, err
}
func GetUserByNickname(nickname string) (User, error) {
	var err error
	var art User

	o := orm.NewOrm()
	o.Using("default")
	art = User{Nickname: nickname}
	err = o.Read(&art, "nickname")

	return art, err
}

func AddUser(account string,nickname string,password string) (error) {
	o := orm.NewOrm()
	o.Using("default")
	ctime := time.Now().Unix()

	sql := "insert into user(account, nickname, password ,time ) values(?, ?, ?, ?)"
	_, err := o.Raw(sql, account, nickname, password, ctime).Exec()
	if nil != err {
		return err
	} else {
		return nil
	}
} 