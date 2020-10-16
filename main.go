package main

import (
	"DataCertPhone/db_mysql"
	_ "DataCertPhone/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.OpenDB()
	defer db_mysql.Db.Close()
	beego.SetStaticPath("../static/css","./static/css")
	beego.SetStaticPath("../static/img","./static/img")
	beego.SetStaticPath("../static/js","./static/js")
	beego.Run()
}

