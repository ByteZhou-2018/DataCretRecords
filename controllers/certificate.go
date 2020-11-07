package controllers

import (
	"DataCertPhone/blockchain"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertificateController struct {
	beego.Controller
}

func (c *CertificateController) Get() {
	cert_id := c.Ctx.Request.FormValue("cert_id")
	fileSize :=c.GetString("fileSize")
	fileName:=c.GetString("fileName")
	fileTitle :=c.GetString("fileTitle")
	cert_time :=c.GetString("certTime")
	phone :=c.GetString("phone")
	fmt.Println(fileSize ,fileName , fileTitle)
	block, err := blockchain.CHAIN.QueryBlockByCertId(cert_id)
	if err != nil {
		c.Ctx.WriteString("抱歉，查询链上数据遇到错误，请重试！")
		return
	}
	if block == nil {//遍历整条区块链，但未查询到数据
		c.Ctx.WriteString("抱歉，未查询到链上数据，请重试！")
		return
	}
	fmt.Println("查询到的区块的高度",block.Height)
	//c.Ctx.WriteString("你好啊")
	//c.Data["Records"] = models.Records
	//c.Data["Phone"] = PhoneStr
	c.Data["CertId"] = strings.ToUpper(string(block.Data))
	c.Data["FileSize"] =fileSize
	c.Data["FileName"] =fileName
	c.Data["FileTitle"] =fileName
	c.Data["CertTime"] =cert_time
	c.Data["Phone"] =phone

	c.TplName = "certificate.html"

}
//func (c *CertificateController) Post() {
//	fileName := c.Ctx.Request.PostFormValue("fileName")
//	fileSize := c.Ctx.Request.PostFormValue("fileSize")
//	certTime := c.Ctx.Request.PostFormValue("certTime")
//	fileTitle := c.Ctx.Request.PostFormValue("fileTitle")
//	fileCert := c.Ctx.Request.PostFormValue("fileCert")
//	phone := c.Ctx.Request.PostFormValue("phone")
//	block := c.Ctx.Request.PostFormValue("block")
//	blockHash := c.Ctx.Request.PostFormValue("blockHash")
//	blockData := c.Ctx.Request.PostFormValue("blockData")
//	c.Data["Phone"] = phone
//	c.Data["Block"] = block
//	c.Data["BlockHash"] = blockHash
//	c.Data["BlockData"] = blockData
//	c.Data["FileName"] = fileName
//	c.Data["FileSize"] = fileSize
//	c.Data["CertTime"] = certTime
//	c.Data["FileTitle"] = fileTitle
//	c.Data["FileCert"] = fileCert
//	//c.TplName = "certificate.html"
//	fmt.Println("电话为：", phone)
//	fmt.Println(fileName)
//	fmt.Println(fileSize, "KB")
//	//c.Ctx.WriteString("欢迎进入展示证书页面")
//	c.TplName = "certificate.html"
//}
