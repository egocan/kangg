package admin

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id     int
	Name   string
	Ename  string
	Pid    int
	Sort   int
	Status int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Category))
}
