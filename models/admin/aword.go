package admin

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AWord struct {
	Id       int
	Source   string
	Content  string
	Created  time.Time 	`orm:"auto_now_add;type(datetime)"`
	Status   int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(AWord))
}