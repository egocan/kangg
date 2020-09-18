package admin

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const ONLINE = 1
const UNSALE = 2
const DELETE = 3

var Status = map[int]string{ONLINE: "在线", UNSALE: "下架", DELETE: "删除"}
var Recommend = map[int]string{0: "否", 1: "是"}


type Article struct {
	Id       int
	Title    string
	Tag      string
	Remark   string
	Desc     string    `orm:"type(text)"`
	Cover    string    `orm:"default(null)" description:"MasterImage"`
	Html     string    `orm:"type(text)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
	Status   	int       `orm:"default(1)"`
	Pv       	int       `orm:"default(0)"`
	Review   	int	   `orm:"default(0)"`
	Recommend 	int		`orm:"default(0)"`
	Like		int  	`orm:"default(0)"`
	User     *User     `orm:"rel(fk)"`
	Category *Category `orm:"rel(one)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Article))
}
