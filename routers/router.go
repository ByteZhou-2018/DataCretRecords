package routers

import (
	"DataCertPhone/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/register.html", &controllers.RegisterController{})

    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/login.html", &controllers.LoginController{})

    beego.Router("/home", &controllers.HomeController{})
    beego.Router("/home.html",&controllers.HomeController{})

    beego.Router("/records",&controllers.RecordsController{})
    //beego.Router("/home.html",&controllers.HomeController{})
}
