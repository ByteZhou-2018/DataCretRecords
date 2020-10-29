package controllers

import (
	"DataCertPhone/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController)Get()  {
	l.TplName ="login.html"
}
func (l *LoginController)Post()  {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Data["Error"] = err.Error()
		l.TplName = "error.html"
	}
	fmt.Println(user.Phone,user.Password)
	user1,err := user.LoginUser()
	if err != nil {
		l.Data["Error"] = err.Error()
		l.TplName = "error.html"
		return
	}
	l.Data["Phone"] = user1.Phone
	l.TplName = "home.html"
}






















