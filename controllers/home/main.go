package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"kangg/models/admin"
	"net/http"
	"path/filepath"
	"strings"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Log("home")

	// 推荐
	o := orm.NewOrm()
	var list []*admin.Article
	o.QueryTable(new(admin.Article)).Filter("status", 1).Filter("recommend", 1).Filter("User__Name__isnull", false).Filter("Category__Name__isnull", false).OrderBy("-id").RelatedSel().All(&list, "id", "title")
	c.Data["Recommend"] = list

	c.Data["home"] = "首页"

	c.TplName = "home/" + beego.AppConfig.String("view") + "/home.html"
}

func(c *MainController) RootStatic(){
	//root_static_ext根目录静态文件后缀
	//允许根目录获取的文件后缀，多个后缀用英文逗号
	ext := strings.ToLower(beego.AppConfig.String("rootStaticExt"))
	if len(ext) == 0 {
		ext = ".txt,.xml,.html"
	}
	file := c.Ctx.Input.Param(":splat")

	if !strings.Contains(file,"/"){
		fileSlice := strings.Split(file,".")
		l := len(fileSlice)
		if l >0 {
			fileExt := strings.ToLower(fileSlice[l-1])
			if strings.Contains(ext,fileExt) {
				path := filepath.Join("static",file)
				http.ServeFile(c.Ctx.ResponseWriter,c.Ctx.Request,path)
			}
		}
	}
	c.Abort("404")

}