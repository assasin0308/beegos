package models

type BookCategory struct {
	Id int
	Bookid int
	CategoryId int
}

func (m *BookCategory) TableName() string {
	return TNBookCategroy()
}