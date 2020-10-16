package controllers

import (
	"DataCertPhone/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

/**
 * 该控制器结构体用于处理文件上传的功能
 */
type HomeController struct {
	beego.Controller
}

/**
 * 该post方法用于处理用户在客户端提交的认证文件
 */
func (h *HomeController) Get() {
	h.TplName = "home.html"
}
func (h *HomeController) Post() { //该post方法用于处理用户在客户端提交的文件
	title := h.Ctx.Request.PostFormValue("title") //用户输入的标题
	fmt.Println("电子数据标签：", title)
	file, header, err := h.GetFile("file")
	if err != nil { //解析客户端提交的文件出现错误
		h.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}

	defer file.Close() // invalid memorey or nil nil pointer dereference// 无效的   内存  或 空 指针 错误
	isJpg := strings.HasSuffix(header.Filename, ".jpg")
	isPng := strings.HasSuffix(header.Filename, ".png")
	if !isJpg && !isPng { //文件类型不支持
		h.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
		return
	}

	config := beego.AppConfig
	fileSize, err := config.Int64("file_size") //文件的大小 200kb
	if header.Size/1024 > fileSize {
		h.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
		return
	}

	savePath := "static/upload" + "/" + header.Filename
	savefile, err := os.OpenFile(savePath, os.O_CREATE|os.O_RDWR, 777)
	if err != nil {
		h.Ctx.WriteString("创建文件失败")
		return
	}
	//使用 io包提供的方法保存文件
	_, err = io.Copy(savefile, file) //length
	if err != nil {
		h.Ctx.WriteString("电子数据认证，请重新尝试！")
		return
	}
	// hash256 加密  fielCert
	hashinstance := sha256.New()
	Filebytes, _ := ioutil.ReadAll(file)
	hashinstance.Write(Filebytes)
	bytes := hashinstance.Sum(nil)

	//////////////////给结构体一个个赋值 /////////////////////////////////////
	//UserId, err := uploadRecordThis.QueryUserId(models.User_this.Phone)
	userId,err := models.QueryUserId(models.User_login.Phone)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("拿userId遇到错误")
		h.Ctx.WriteString("拿useid出错，请重新尝试！")
		return
	}
	//var uploadRecordThis models.UploadRecord{
	//
	//}
	uploadFileInfo := models.UploadRecord{
		//Id:        userId,
		UserId:    userId,
		FileName: header.Filename,
		FileSize:  header.Size,
		FileCert:  hex.EncodeToString(bytes),
		FileTitle:  title,
		CertTime:  time.Now().String(),
	}

	//////////////////存入数据库 /////////////////////////////
	_, err = uploadFileInfo.AddFiles()
	if err != nil {
		fmt.Printf("文件信息插入数据库出错！！！\n")
		fmt.Println(err.Error())
		h.Ctx.WriteString("电子数据认证信息保存失败，请重新尝试！")
		return
	}

	////////////////////// 从数据库中读取该用户存储的所有文件记录/////////////////////////////////
	records,err := models.QueryRecordsByUserId(userId)
	if err != nil {
		fmt.Println(err.Error())
		h.Ctx.WriteString("电子数据认证信息获取失败,请稍后重试!")
		return
	}
	h.Data["Records"] = records
	h.TplName = "recordsList.html"
		//h.Ctx.WriteString("电子数据认证成功")

}




























//func (h *HomeController) Post() {
//	//1、解析用户上传的数据及文件内容
//	//用户上传的自定义的标题
//	title := h.Ctx.Request.PostFormValue("title") //用户输入的标题
//
//	//用户上传的文件
//	file, header, err := h.GetFile("file")
//	if err != nil { //解析客户端提交的文件出现错误
//		h.Ctx.WriteString("抱歉，文件解析失败，请重试！")
//		return
//	}
//	defer file.Close()
//	fmt.Println("自定义的标题：", title)
//	//获得到了上传的文件
//	fmt.Println("上传的文件名称:", header.Filename)
//	//eg：支持jpg,png类型，不支持jpeg，gif类型
//	//文件名： 文件名 + "." + 扩展名
//	fileNameSlice := strings.Split(header.Filename, ".")
//	fileType := fileNameSlice[1]
//	fmt.Println(fileNameSlice)
//	fmt.Println(":", fileType)
//	isJpg := strings.HasSuffix(header.Filename, ".jpg")
//	isPng := strings.HasSuffix(header.Filename, ".png")
//	if !isJpg && !isPng {
//		//文件类型不支持
//		h.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
//		return
//	}
//
//	//if fileType != " jpg" || fileType != "png" {
//	//	//文件类型不支持
//	//	u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
//	//	return
//	//}
//
//	//文件的大小 200kb
//	config := beego.AppConfig
//	fileSize, err := config.Int64("file_size")
//
//	if header.Size/1024 > fileSize {
//		h.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
//		return
//	}
//
//	fmt.Println("上传的文件的大小:", header.Size) //字节大小
//
//	fmt.Println("上传的文件:", file)
//	//保存上传文件到文件存储系统  	h.SaveToFile()
//	//formFile ：文件
//	//tofile ：要保存的文件
//	//h.SaveToFile()
//	saveDir := "static/upload"
//	//①先尝试打开文件夹
//	_, err = os.Open(saveDir)
//	//os.OpenFile("文件名",os.O_CREATE|os.O_RDWR,os.ModePerm)
//	if err != nil {
//		////打开文件遇到错误
//		//fmt.Println(err.Error())
//		//h.Ctx.WriteString("打开文件夹目录失败")
//		//return
//		//② 创建文件夹
//		err = os.Mkdir(saveDir, os.ModePerm) //prem:permission 权限
//		if err != nil {
//			fmt.Println(err.Error())
//			h.Ctx.WriteString("抱歉，文件认证遇到错误，请重试！")
//			return
//		}
//
//	}
//	//fmt.Println("打开文件夹", f.Name())
//
//	//权限的组成： a + b + c  最高权限： 777 : rwx rwx rwx
//	//a :文件所有者对文件的操作权限  读 写 执行   r w x
//	//b :文件所有者所在组的用户的操作权限  读 写 执行   r w x
//	//c ：其他用户的操作权限   读 写 执行   r w x
//
//	//文件名 : 文件路径 + 文件名 + "." +文件扩展名
//	saveName := saveDir + "/" + header.Filename
//	fmt.Println("要保存的文件名:",saveName)
//	fmt.Println(saveName)
//	err = h.SaveToFile("file",saveName)
//	if err != nil {
//		fmt.Println(err.Error())
//		h.Ctx.WriteString("失败了")
//		return
//	}
//	h.Ctx.WriteString("已获取到上传文件。")
//}
