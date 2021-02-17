package models

import "github.com/astaxie/beego/orm"

func init() {
	// 向orm注册model
	orm.RegisterModel(
		new(Category),
		new(Book),
		new(BookCategory),
		)
}

func TNCategory() string {
	return "md_category"
}
func TNBook() string {
	return "md_books"
}
func TNBookCategroy() string{
	return "md_book_category"
}