package admin

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Status   int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

func Login(username,password string) (*User,bool) {
	o := orm.NewOrm()
	var (
		user User
		err  error
	)
	ok := false
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("status", 1).And("Name", username).Or("Email", username)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err = qs.One(&user); err == nil {
		if user.Password == password {
			ok = true
		}
	}
	return &user,ok
}