package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	_ "hello/sysinit"
)

func main() {
	beego.Run()
}

