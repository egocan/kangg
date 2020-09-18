package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"os"
)

var Installer = func(ctx *context.Context) {
	_ , err := os.Stat(beego.AppPath + "/installed")
	if err != nil &&
		ctx.Request.RequestURI != "/installer" &&
		ctx.Request.RequestURI != "/installer/create" {
		ctx.Redirect(302,"/installer")
	}
}