package controllers

import (
	"DataCertPhone/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ListController struct {
	beego.Controller
}

func ( l *ListController) Get(){

	phone := l.GetString("phone")
	userId,err := models.QueryUserId(phone)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("拿userId遇到错误")
		l.Ctx.WriteString("获取用户电子认证数据失败，请重新尝试！")
		return
	}
		records,err := models.QueryRecordsByUserId(userId)
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("电子数据认证信息获取失败,请稍后重试!")
		return
	}
	l.Data["Phone"] = phone
	l.Data["Records"] = records
	l.TplName = "recordsList.html"

	//phone1 := l.Ctx.Input.Header("phone")//获取请求头中属性为:phone的值
	//phone2 := l.Ctx.Request.Form.Get("phone")//form表单为get提交方法时有效.
	//phone3 := l.Ctx.Input.Query("phone")//输入信息中查询 key 为phone的值
	//phone4 :=l.GetString("phone")//查询输入
	//fmt.Printf("input.header:%v\n,request.form.get:%v\n,input.query:%v\n,input.getstring:%v\n",phone1,phone2,phone3,phone4)
	//l.Ctx.Output


}