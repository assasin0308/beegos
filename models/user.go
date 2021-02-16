package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
type Bee_user struct {
	Id int
	Name string
	Age int
	Sex string
}

func  init() {
	orm.RegisterDataBase("default","mysql", "root:a12345@tcp(127.0.0.1:3306)/beegos?charset=utf8&loc=Local", 30)
	orm.RegisterModel(new(Bee_user))
}

func PrintUserByOrm(){
	o := orm.NewOrm()
	user := Bee_user{Id:2}
	o.Read(&user)
	fmt.Println(user)
}