package routers

import (
	"DataCertPhone/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //默认路由
    beego.Router("/", &controllers.MainController{})
    //注册
    beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/register.html", &controllers.RegisterController{})
//登录
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/login.html", &controllers.LoginController{})
//文件上传
    beego.Router("/home", &controllers.HomeController{})
    beego.Router("/home.html",&controllers.HomeController{})
//文件展示 页面 Get直接登录
    beego.Router("/list",&controllers.ListController{})//list?phonee=18379139021
    //查看证书
    beego.Router("/certificate",&controllers.CertificateController{})
    beego.Router("/ShowCertificate",&controllers.CertificateController{})
}
