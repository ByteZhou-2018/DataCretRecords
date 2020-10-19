package main

import (
	"DataCertPhone/blockchain"
	"DataCertPhone/db_mysql"
	_ "DataCertPhone/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.OpenDB()
	defer db_mysql.Db.Close()

	block0 := blockchain.CreatGenesisBlock()
	blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte(""))


	beego.SetStaticPath("../static/css","./static/css")
	beego.SetStaticPath("../static/img","./static/img")
	beego.SetStaticPath("../static/js","./static/js")
	beego.Run()
}

