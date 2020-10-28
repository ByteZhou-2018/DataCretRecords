package main

import (
	"DataCertPhone/blockchain"
	"DataCertPhone/db_mysql"
	_ "DataCertPhone/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	bc := blockchain.NewBolckChain()
	fmt.Printf("最新区块的哈希值：%x\n",bc.LastHash)//00033049740c82d05c9cc7e63884b538cf6e88ddd9295ffe368fa810535cd572
	//bc.SaveData([]byte("用户A的数据"))
	//bc.SaveData([]byte("用户B的数据"))
	blocksMap ,err:= bc.Each()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("区块链的长度为", len(blocksMap))
	for i,v := range blocksMap {
		fmt.Println(i,v.Height)
	}

	return


	db_mysql.OpenDB()
	defer db_mysql.Db.Close()
	beego.SetStaticPath("../static/css","./static/css")
	beego.SetStaticPath("../static/img","./static/img")
	beego.SetStaticPath("../static/js","./static/js")
	beego.Run()
}

//block0,err := blockchain.CreatGenesisBlock()
//if err != nil {
//	fmt.Println(err.Error())
//}
//block1,err := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte(""))
//blockchain.BlockChain = append(blockchain.BlockChain, *block1)
////sysByte ,err:= utils.Int64ToByte(023)
////if err != nil {
////	fmt.Println(err.Error())
////}
////sysBytes := []byte("1234561136逗比哈皮dadw1234111")
////sysByte,err = utils.SHA256HashByte(sysByte)
////if err != nil {
////	fmt.Println(err.Error())
////}
//////blockchain.Minner(sysBytes)
////sysBytes,err = utils.SHA256HashByte(sysBytes)
////if err != nil {
////	fmt.Println(err.Error())
////}
////fmt.Println(sysByte)
////fmt.Println(sysBytes)
////fmt.Println(hex.EncodeToString(sysByte))
////fmt.Println(hex.EncodeToString(sysBytes))
////
////fmt.Println(blockchain.CompareBytes(sysBytes,sysByte))// s < e时
//sysBytes := blockchain.GetSySBytes(2)//1603253010587794800   1603253010693586200
//blockchain.Minner(sysBytes,nil)


//block0 := blockchain.CreatGenesisBlock()
//
//
//
//fmt.Println(block0)
//fmt.Println(hex.EncodeToString(block0.Hash))