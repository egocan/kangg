package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
	"kangg/utils"
)

type CateController struct {
	BaseController
}

func (c *CateController) List() {

	o := orm.NewOrm()
	category := new(admin.Category)

	var categorys []*admin.Category
	qs := o.QueryTable(category)
	qs = qs.Filter("status", 1)

	qs.OrderBy("-sort").All(&categorys)

	tree := utils.CategoryTreeR(categorys, 0, 0)
	/*c.Data["json"]= tree
	c.ServeJSON()
	c.StopRun()*/

	c.Data["Category"] = tree
	c.TplName = "admin/cate.html"
}

func (c *CateController) Add() {

	o := orm.NewOrm()
	category := new(admin.Category)
	var categorys []*admin.Category
	qs := o.QueryTable(category)
	qs = qs.Filter("status", 1)
	qs.All(&categorys)

	c.Data["Category"] = categorys
	c.TplName = "admin/cate-add.html"
}

func (c *CateController) Save() {

	name := c.GetString("name")
	ename := c.GetString("ename")

	pid, _ := c.GetInt("pid", 0)
	sort, _ := c.GetInt("sort", 0)

	o := orm.NewOrm()
	category := admin.Category{
		Name:   name,
		Ename:  ename,
		Sort:   sort,
		Pid:    pid,
		Status: 1,
	}

	/*c.Data["json"] = &article
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(category.Name, "Name")
	//valid.Required(category.Ename, "Ename")
	valid.Required(category.Sort, "Sort")
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



	if id, err := o.Insert(&category); err == nil {
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

func (c *CateController) Update() {
	id, _ := c.GetInt("id", 0)
	name := c.GetString("name")
	ename := c.GetString("ename")
	sort, _ := c.GetInt("sort", 0)
	pid, _ := c.GetInt("pid", 0)

	response := make(map[string]interface{})

	o := orm.NewOrm()

	category := admin.Category{Id: id}
	if o.Read(&category) == nil {
		category.Name = name
		category.Ename = ename
		category.Pid = pid
		category.Sort = sort

		valid := validation.Validation{}
		valid.Required(category.Name, "Name")
		valid.Required(category.Sort, "Sort")
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



		if _, err := o.Update(&category); err == nil {
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

func (c *CateController) Delete() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})

	o := orm.NewOrm()
	category := admin.Category{Id: id}

	if o.Read(&category) == nil {
		category.Status = 2

		if _, err := o.Update(&category); err == nil {
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

func (c *CateController) Put() {
	id, err := c.GetInt("id", 0)

	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	category := new(admin.Category)
	var categorys []*admin.Category
	qs := o.QueryTable(category)
	err = qs.Filter("id", id).One(&categorys)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = categorys[0]

	var categorysTree []*admin.Category
	qs = o.QueryTable(category)
	qs = qs.Filter("status", 1)
	qs.All(&categorysTree)

	c.Data["Category"] = categorysTree
	/*c.Data["json"]= &articles
	c.ServeJSON()
	c.StopRun()*/

	c.TplName = "admin/cate-edit.html"

}
