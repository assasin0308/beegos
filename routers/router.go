package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Router("/",&controllers.HomeController{},"get:Index")
}
