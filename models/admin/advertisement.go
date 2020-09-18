package admin

import "github.com/astaxie/beego/orm"

type Ad struct {
	Id 		 int      `orm:"size(11);pk"`
	Name   	 string    `orm:"size(255)"`
	Title    string
	Html     string    `orm:"type(text)"`
}

func init()  {
	orm.RegisterModel(new(Ad))
}