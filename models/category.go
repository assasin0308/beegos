package models

import "github.com/astaxie/beego/orm"

type Category struct {
	Id int `orm:"pk;auto"`
	Pid int
	Title string `orm:"size(30);unique"`
	intro string
	Icon string
	Cnt int // 统计下分类下图书
	Sort int //排序
	Status bool // 状态 true:显示
}

func (m *Category) TableName() string {
	return TNCategory()
}
// 获取所有分类
func (m *Category) GetCates(pid,status int) (cates []Category,err error) {
	querys := orm.NewOrm().QueryTable(TNCategory())
	if pid > -1 {
		querys = querys.Filter("pid",pid)
	}
	if 0 == status || 1 == status {
		querys = querys.Filter("status",status)
	}
	_,err = querys.OrderBy("-status","sort","title").All(&cates)
	return
}

func (m *Category) Find(cid int) (cate Category) {
	cate.Id = cid
	orm.NewOrm().Read(&cate)
	return cate
}