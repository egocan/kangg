package installer

import (
	"github.com/astaxie/beego"
	"os"
)

type InstallController struct {
	beego.Controller
}

func (c *InstallController) CheckEnv()  {

	c.Data["json"] = "不行,还没安装!"
	c.ServeJSON()
	c.StopRun()
}

func (c *InstallController) Install() {
	f, err := os.Create("installed")
	defer f.Close()

	//判断是否出错
	if err != nil {
		c.Data["json"] = "安装失败!ERROR:" + err.Error()
		c.ServeJSON()
		c.StopRun()
	}

	c.Redirect("/",302)
}