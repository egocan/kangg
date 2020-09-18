package admin

import (
	"github.com/astaxie/beego"
	"kangg/models/admin"
	"time"
)

type BaseController struct {
	beego.Controller
	UserName  string
	LastLogin time.Time
	User      admin.User
}

func (ctl *BaseController) Prepare() {
	ctl.StartSession()
	//ctl.Data["PageStartTime"] = time.Now()
	user := ctl.GetSession("User")
	if user != nil {
		ctl.User = user.(admin.User)
		ctl.Data["LoginUser"] = user
		//ctl.Data["LastLogin"] = ctl.GetSession("LastLogin")
	}else{
		//ctl.TplName = "admin/login.html"
		//render.Redirect{401,'/',"1"}
		ctl.Redirect("/admin/login",302)
	}

}
