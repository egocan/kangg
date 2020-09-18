package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
)

type BannerController struct {
	BaseController
}

func (c *BannerController) List() {

	o := orm.NewOrm()
	banner := new(admin.Banner)

	var banners []*admin.Banner
	qs := o.QueryTable(banner)
	qs = qs.Filter("Article__Title__isnull", false)

	// 获取数据
	_, err := qs.OrderBy("id").RelatedSel().All(&banners)

	if err != nil {
		c.Abort("404")
	}

	c.Data["Banner"] = &banners

	c.TplName = "admin/banner.html"
}

func (c *BannerController) Update() {
	id, _ := c.GetInt("id", 0)
	name := c.GetString("name")
	image := c.GetString("image")
	aid,_ := c.GetInt("aid")


	response := make(map[string]interface{})

	o := orm.NewOrm()

	banner := admin.Banner{Id: id}
	if o.Read(&banner) == nil {
		banner.Name = name
		banner.Image = image


		valid := validation.Validation{}
		valid.Required(banner.Image, "Image")
		valid.Required(banner.Article, "Article")

		if valid.HasErrors() {
			// 如果有错误信息，证明验证没通过
			// 打印错误信息
			for _, err := range valid.Errors {
				//log.Println(err.Key, err.Message)
				response["msg"] = "新增失败！"
				response["code"] = 500
				response["err"] = err.Key + " " + err.Message
				c.Data["json"] = response
				c.ServeJSON()
				c.StopRun()
			}
		}

		banner.Article = &admin.Article{Id:aid,Title: ""}

		if _, err := o.Update(&banner); err == nil {
			response["msg"] = "修改成功！"
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "修改失败！"
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "修改失败！"
		response["code"] = 500
		response["err"] = "ID 不能为空！"
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}


func (c *BannerController) Put() {
	id, err := c.GetInt("id", 0)

	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	banner := new(admin.Banner)
	var banners []*admin.Banner
	qs := o.QueryTable(banner)
	err = qs.Filter("id", id).One(&banners)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = banners[0]

	// 文章数据
	o = orm.NewOrm()
	article := new(admin.Article)
	var articles []*admin.Article
	qs = o.QueryTable(article)
	qs = qs.Filter("status", 1)
	qs.All(&articles)

	c.Data["Article"] = articles

	c.TplName = "admin/banner-edit.html"

}