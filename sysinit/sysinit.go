package sysinit

import (
	"github.com/astaxie/beego"
	"path/filepath"
)

func sysinit () {
	uploads := filepath.Join("./","uploads")
	beego.BConfig.WebConfig.StaticDir["/uploads"] = uploads

	// 注册前端使用函数
	registerFunctions()
}

func registerFunctions() {

}
