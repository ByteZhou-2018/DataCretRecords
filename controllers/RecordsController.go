package controllers

import "github.com/astaxie/beego"

type RecordsController struct {
	beego.Controller
}

func (r *RecordsController)Get()  {
	r.TplName ="recordsList.html"
}