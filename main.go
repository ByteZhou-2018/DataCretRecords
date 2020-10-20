package main

import (
	"DataCertPhone/blockchain"
	"DataCertPhone/db_mysql"
	_ "DataCertPhone/routers"
	"DataCertPhone/utils"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.OpenDB()
	defer db_mysql.Db.Close()

	block0,err := blockchain.CreatGenesisBlock()
	if err != nil {
		fmt.Println(err.Error())
	}
	block1,err := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte(""))
	blockchain.BlockChain = append(blockchain.BlockChain, *block1)
	sysByte ,err:= utils.Int64ToByte(023)
	if err != nil {
		fmt.Println(err.Error())
	}
	sysBytes := []byte("1234561136逗比哈皮dadw1234111")
	sysByte,err = utils.SHA256HashByte(sysByte)
	if err != nil {
		fmt.Println(err.Error())
	}
	//blockchain.Minner(sysBytes)
	sysBytes,err = utils.SHA256HashByte(sysBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(sysByte)
	fmt.Println(sysBytes)
	fmt.Println(hex.EncodeToString(sysByte))
	fmt.Println(hex.EncodeToString(sysBytes))

	fmt.Println(blockchain.CompareBytes(sysBytes,sysByte))// s < e时
	blockchain.GetSySBytes()








	beego.SetStaticPath("../static/css","./static/css")
	beego.SetStaticPath("../static/img","./static/img")
	beego.SetStaticPath("../static/js","./static/js")
	beego.Run()
}

