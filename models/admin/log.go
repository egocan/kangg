package admin

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Log struct {
	Id       	int
	Ip    		string
	City   		string
	UserAgent   string    	`orm:"size(500)"`
	Create  	time.Time 	`orm:"auto_now_add;type(datetime)"`
	Page 		string
	Uri 		string		`orm:"size(500)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Log))
}
