package main

import (
	"DataCertPhone/blockchain"
	"DataCertPhone/db_mysql"
	_ "DataCertPhone/routers"
	"github.com/astaxie/beego"
)

func main() {

//先准备一条区块链
	 blockchain.NewBolckChain()


		db_mysql.OpenDB()
		defer db_mysql.Db.Close()
		beego.SetStaticPath("../static/css", "./static/css")
		beego.SetStaticPath("../static/img", "./static/img")
		beego.SetStaticPath("../static/js", "./static/js")
		beego.Run()
	}
/*
//fmt.Println(123)
	bc := blockchain.NewBolckChain()// 0 1 224 .。。
	//block1,err := bc.SaveData([]byte("你好，世界"))
	//if err != nil {
	//	 fmt.Println("新添区块错误：",err.Error())
	//	//fmt.Println("哈哈")
	//	return
	//}
	//fmt.Println(block1.Hash)//
	//block,err := bc.QueryBlockByHeight(2)
	//if err  != nil{
	//	 fmt.Println(err.Error())
	//}
	//fmt.Println(block.Height)
	blocks,err := bc.QueryAllBolcks()
	if err != nil {

		fmt.Println(err.Error())

	}
	fmt.Println("区块链的长度为：",len(blocks))
	for _,v := range blocks {
		fmt.Printf("%d,\t,%x\n",v.Height,v.Hash)
	}
	return









*/
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
//bc := blockchain.NewBolckChain()
//blocks,err := bc.QueryAllBolcks()
//if err != nil {
//	fmt.Println(err.Error())
//	return
//}
//for i,v := range blocks{
//	fmt.Printf("序号：%d,区块高度:%d,区块哈希：%x",i,v.Height,v.Hash)
//}
//block,err := bc.QueryBlockByHeight(3)
//if err != nil {
//	fmt.Println(err.Error())
//}
//fmt.Println(block)
//fmt.Printf("最新区块的哈希值：%x\n", bc.LastHash) //00033049740c82d05c9cc7e63884b538cf6e88ddd9295ffe368fa810535cd572
//block1, err := bc.SaveData([]byte("用户A的数据"))
//if err != nil {
//	fmt.Println(err.Error())
//	return
//
//	fmt.Printf("%+v\n", block1.Height)
//
//	fmt.Printf("%+v\n", block1.Hash)
//	fmt.Printf("%+v\n", block1.PrevHash)
//	//bc.SaveData([]byte("用户B的数据"))
//	blocksMap, err := bc.Each()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println("区块链的长度为", len(blocksMap))
//	for i, v := range blocksMap {
//		fmt.Println(i, v.Height)
//	}
//
//	return
