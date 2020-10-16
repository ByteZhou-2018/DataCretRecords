package db_mysql

import (
	//"DataCertPhone/models"
	"database/sql"
	"github.com/astaxie/beego"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var Db *sql.DB
func OpenDB()  {
	if Db != nil{
		return
	}
	config := beego.AppConfig
	db_dirver_name := config.String("db_dirver_name")
	db_admin_name := config.String("db_admin_name")
	db_admin_pwd := config.String("db_admin_pwd")
	db_port := config.String("db_port")
	db_name := config.String("db_name")
	database,err := sql.Open(db_dirver_name,db_admin_name+":"+db_admin_pwd+"@tcp("+db_port+")/"+db_name+"?charset=utf8")


	if err != nil {
		fmt.Println(err.Error())
	}

	Db =database

}






































//db,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/data_phone?charset=utf8")
//fmt.Println(db_dirver_name,db_admin_name+":"+db_admin_pwd+"@tcp("+db_port+")/"+db_name+"?charset=utf8")

//appName :=config.String("appname")
//fmt.Println("程序名称",appName)
//httpPort,err := config.Int("httpport")
//if err != nil {
//	panic("程序错误")
//
//}
//fmt.Println("监听端口",httpPort)
//"mysql", "root:123456@tcp(127.0.0.1:3306)/weather?charset=utf8"
//fmt.Println(db)
//"mysql", "root:123456@tcp(127.0.0.1:3306)/weather?charset=utf8"
//			root:123456@tcp(127.0.0.1:3306)/weather?charset=utf8
//fmt.Println(db_admin_name+":"+db_admin_pwd+"@tcp("+db_port+")/"+db_name+"?charset=utf8")