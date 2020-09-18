package admin

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Review struct {
	Id       	int
	Name    	string
	Review   	string 		`orm:"size(500)"`
	Reply     	string    	`orm:"size(500)"`
	Mail   	    string		`orm:"size(500)"`
	Created  	time.Time 	`orm:"auto_now_add;type(datetime)"`
	Updated  	time.Time 	`orm:"auto_now;type(datetime)"`
	Status   	int       	`orm:"default(1)"`
	ArticleId 	int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Review))
}
