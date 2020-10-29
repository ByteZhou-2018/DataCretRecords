package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type CertificateController struct {
	beego.Controller
}

func (c *CertificateController) Get() {
	//c.Ctx.WriteString("你好啊")
	//c.Data["Records"] = models.Records
	//c.Data["Phone"] = PhoneStr
	c.TplName = "certificate.html"

}
func (c *CertificateController) Post() {
	fileName := c.Ctx.Request.PostFormValue("fileName")
	fileSize := c.Ctx.Request.PostFormValue("fileSize")
	certTime := c.Ctx.Request.PostFormValue("certTime")
	fileTitle := c.Ctx.Request.PostFormValue("fileTitle")
	fileCert := c.Ctx.Request.PostFormValue("fileCert")
	phone := c.Ctx.Request.PostFormValue("phone")
	c.Data["Phone"] = phone
	c.Data["FileName"] = fileName
	c.Data["FileSize"] = fileSize
	c.Data["CertTime"] = certTime
	c.Data["FileTitle"] = fileTitle
	c.Data["FileCert"] = fileCert
	//c.TplName = "certificate.html"
	fmt.Println(phone)
	fmt.Println(fileName)
	fmt.Println(fileSize,"KB")
	//c.Ctx.WriteString("欢迎进入展示证书页面")
		c.TplName = "certificate.html"
}
