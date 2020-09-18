package admin

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const WOMAN = 0
const MAN = 1
const UNKNOWN = 2

var Sex = map[int]string{MAN: "男", WOMAN: "女", UNKNOWN: "未知"}

type WechatUser struct {
	Id 				int
	Openid 			string 		`orm:"size(32)"`
	Nickname 		string
	Sex 			int32		`orm:"default(2)"`
	City 			string 		`orm:"size(32)"`
	Country			string 		`orm:"size(32)"`
	Province		string 		`orm:"size(32)"`
	Language		string 		`orm:"size(32)"`
	Headimgurl 		string
	SubscribeTime 	int32
	Created 		time.Time 	`orm:"auto_now_add;type(datetime)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(WechatUser))
}
