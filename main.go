package main

import (
	"DataCertPhone/blockchain"
	"DataCertPhone/db_mysql"
	_ "DataCertPhone/routers"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.OpenDB()
	defer db_mysql.Db.Close()


block0 := blockchain.CreatGenesisBlock()



fmt.Println(block0)
fmt.Println(hex.EncodeToString(block0.Hash))


//序列化  ：将数据从内存中的形式转换为一种可以持久化存储在硬盘上或者在网络上传输的形式。称之为序列化
//反序列		：将数据从文件中或网络中读取，然后转化到计算机内存中的过程
//只有进行序列化以后的对象才能进行传输
//序列化和反序列化有很多种方式：
//json: 序列化 ：json.Marshal 反序列化 ：json.UnMarshal
//xml: 序列化 ：xml.Marshal 反序列化 ：xml.UnMarshal

//blockJson,_ := json.Marshal(block0)
blockXml,_ := xml.Marshal(block0)
blockGob := block0.Serialize()
fmt.Println("序列化以后的block",string(blockGob))
fmt.Println("xml序列化以后的block",string(blockXml))
//blockAsn1,_ :=asn1.Marshal(block0)
//fmt.Println("ans1序列化后的block",string(blockAsn1))


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