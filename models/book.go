package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type Book struct {
	BookId int `orm:"pk;auto" json:"book_id"`
	BookName string `orm:"size(500)" json:"book_name"`
	Identify string `orm:"size(100);unique" json:"identify"`
	OrderIndex int `orm:"default(0)" json:"order_index"`
	Description string `orm:"size(1000)" json:"description"`
}

func (m *Book) TableName() string {
	return TNBook()
}

func (m *Book) HomeData(pageIndex,pageSize int,cid int,fields ...string) (books []Book,totalCount int,err error) {
		if len(fields) == 0 {
			fields = append(fields,"book_id","book_name","identify","order_index")
		}
		fieldsStr := "b." + strings.Join(fields,",b.")
		sqlFormat := "SELECT %v FROM "+ TNBook() +" b LEFT JOIN "+ TNBookCategroy() +" c " +
			"ON b.book_id = c.book_id WHERE c.category_id = %v" + strconv.Itoa(cid)
		sql := fmt.Sprintf(sqlFormat,fieldsStr)
		sqlCount := fmt.Sprintf(sqlFormat,"count(*) cnt ",cid)

		o := orm.NewOrm()
		var params []orm.Params
		if _,err := o.Raw(sqlCount).Values(&params);err == nil {
			if len(params) > 0 {
				totalCount,_ = strconv.Atoi(params[0]["cnt"].(string))
			}
		}
		_,err = o.Raw(sql).QueryRow(&books)
		return
}