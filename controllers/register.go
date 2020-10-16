package controllers

import (
	"DataCertPhone/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
func (r *RegisterController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	r.TplName = "register.html"
}
func (r *RegisterController)Post(){
	var user models.User
	err := r.ParseForm(&user)
	if err  != nil{
		fmt.Println(err.Error())
		panic("r.ParseForm err :")
	}
	fmt.Println(user)

	_ ,err = user.AddUser()
	if err != nil {
		r.Ctx.WriteString("注册用户信息失败！请重试")
	}
	r.TplName = "login.html"
}