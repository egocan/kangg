package admin

import "github.com/astaxie/beego/orm"

type Banner struct {
	Id 		 int      `orm:"size(11);pk"`
	Name   	 string   `orm:"size(255)"`
	Image	 string
	Article  *Article `orm:"rel(one)"`
}

func init()  {
	orm.RegisterModel(new(Banner))
}