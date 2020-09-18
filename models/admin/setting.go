package admin

import (
	"github.com/astaxie/beego/orm"
)

type Setting struct {
	Name   	 string    		`orm:"size(255);pk"`
	Value    string      	`orm:"type(text)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Setting))
}