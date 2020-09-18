package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
)

type AdController struct {
	BaseController
}

func (c *AdController) List() {

	o := orm.NewOrm()
	ad := new(admin.Ad)

	var ads []*admin.Ad
	qs := o.QueryTable(ad)
	//qs = qs.Filter("name__isnull", false)

	// 获取数据
	_, err := qs.OrderBy("id").All(&ads)

	if err != nil {
		c.Abort("404")
	}

	c.Data["Ad"] = &ads

	c.TplName = "admin/ad.html"
}

func (c *AdController) Update() {
	id, _ := c.GetInt("id", 0)
	name := c.GetString("name")
	title := c.GetString("title")
	html := c.GetString("desc_html")



	response := make(map[string]interface{})

	o := orm.NewOrm()

	ad := admin.Ad{Id:id}
	if o.Read(&ad) == nil {
		ad.Name = name
		ad.Title = title
		ad.Html = html



		valid := validation.Validation{}
		//valid.Required(ad.Image, "Image")
		valid.Required(ad.Name, "Name")

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


		if _, err := o.Update(&ad); err == nil {
			response["msg"] = "修改成功！"
			response["code"] = 200
			response["name"] = name
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


func (c *AdController) Put() {
	id, err := c.GetInt("id", 0)
	pre := c.GetString("preview")
	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	ad := new(admin.Ad)
	var ads []*admin.Ad
	qs := o.QueryTable(ad)
	err = qs.Filter("id", id).One(&ads)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = ads[0]

	if len(pre) !=0 {
		c.TplName = "admin/ad-preview.html"
	}else{
		c.TplName = "admin/ad-edit.html"
	}


}