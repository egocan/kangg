package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
	"kangg/utils"
)

type AWordController struct {
	BaseController
}

func (c *AWordController) List() {
	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量

	o := orm.NewOrm()
	aword := new(admin.AWord)

	var awords []*admin.AWord
	qs := o.QueryTable(aword)
	qs = qs.Filter("status", 1)

	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&awords)

	if err != nil {
		c.Abort("404")
	}

	// 统计
	count, err := qs.Count()
	if err != nil {
		c.Abort("404")
	}

	c.Data["AWord"] = &awords
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)

	c.TplName = "admin/aword.html"
}

func (c *AWordController) Add() {

	o := orm.NewOrm()
	aword := new(admin.AWord)
	var awords []*admin.AWord
	qs := o.QueryTable(aword)
	qs = qs.Filter("status", 1)
	qs.All(&awords)

	c.Data["Category"] = awords
	c.TplName = "admin/aword.html"
}

func (c *AWordController) Save() {

	source := c.GetString("source")
	content := c.GetString("content")


	o := orm.NewOrm()
	aword := admin.AWord{
		Source:   source,
		Content:  content,
		Status: 1,
	}

	/*c.Data["json"] = &article
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(aword.Source, "Source")
	//valid.Required(category.Ename, "Ename")
	valid.Required(aword.Content, "Content")
	//valid.Required(category.Pid, "Pid")

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

	if id, err := o.Insert(&aword); err == nil {
		response["msg"] = "新增成功！"
		response["code"] = 200
		response["id"] = id
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()

}

func (c *AWordController) Update() {
	id, _ := c.GetInt("id", 0)
	source := c.GetString("source")
	content := c.GetString("content")


	response := make(map[string]interface{})

	o := orm.NewOrm()

	aword := admin.AWord{Id: id}
	if o.Read(&aword) == nil {
		aword.Source = source
		aword.Content = content


		valid := validation.Validation{}
		valid.Required(aword.Source, "Source")
		valid.Required(aword.Content, "Content")
		//valid.Required(category.Pid, "Pid")

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



		if _, err := o.Update(&aword); err == nil {
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

func (c *AWordController) Delete() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})

	o := orm.NewOrm()
	word := admin.AWord{Id: id}

	if o.Read(&word) == nil {
		word.Status = 2

		if _, err := o.Update(&word); err == nil {
			response["msg"] = "删除成功！"
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "删除失败！"
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "删除失败！"
		response["code"] = 500
		response["err"] = "ID 不能为空！"
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *AWordController) Put() {
	id, err := c.GetInt("id", 0)

	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	word := new(admin.AWord)
	var words []*admin.AWord
	qs := o.QueryTable(word)
	err = qs.Filter("id", id).One(&words)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = words[0]

	/*c.Data["json"]= &articles
	c.ServeJSON()
	c.StopRun()*/

	c.TplName = "admin/aword-edit.html"

}