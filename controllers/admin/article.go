package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
	"kangg/utils"
	"time"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) List() {

	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量

	start := c.GetString("start")
	end := c.GetString("end")
	status, _ := c.GetInt("status", 0)
	title := c.GetString("title")

	c.Data["Start"] = start
	c.Data["End"] = end
	c.Data["Status"] = status
	c.Data["Title"] = title

	o := orm.NewOrm()
	article := new(admin.Article)

	var articles []*admin.Article
	qs := o.QueryTable(article)
	qs = qs.Filter("User__Name__isnull", false)
	qs = qs.Filter("Category__Name__isnull", false)

	// 状态
	if status != 0 {
		qs = qs.Filter("status", status)
	}else{
		qs = qs.Filter("status__lt", 3)
	}

	// 开始时间
	if start != "" {
		qs = qs.Filter("created__gte", start)
	}

	// 结束时间
	if end != "" {
		qs = qs.Filter("created__lte", end)
	}

	// 标题
	if title != "" {
		qs = qs.Filter("title__icontains", title)
	}

	// 过滤删除掉的
	/*cond := orm.NewCondition()
	cond1 := cond.AndNot("status", 3)
	qs = qs.SetCond(cond1)*/

	// 获取数据
	qs.OrderBy("-id","-pv").RelatedSel().Limit(limit).Offset(offset).All(&articles)

	// 统计
	count, _ := qs.Count()

	c.Data["Data"] = &articles
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["StatusText"] = admin.Status
	c.Data["RecommendText"] = admin.Recommend

	c.TplName = "admin/article-list.html"
}

func (c *ArticleController) Add() {

	o := orm.NewOrm()
	category := new(admin.Category)

	var categorys []*admin.Category
	qs := o.QueryTable(category)
	qs = qs.Filter("status", 1)
	qs.All(&categorys)

	c.Data["Category"] = categorys
	c.TplName = "admin/article-add.html"
}

func (c *ArticleController) Put() {
	id, err := c.GetInt("id", 0)

	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	article := new(admin.Article)
	var articles []*admin.Article
	qs := o.QueryTable(article)
	err = qs.Filter("id", id).One(&articles)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = &articles[0]

	// 分类数据
	o = orm.NewOrm()
	category := new(admin.Category)
	var categorys []*admin.Category
	qs = o.QueryTable(category)
	qs = qs.Filter("status", 1)
	qs.All(&categorys)

	c.Data["Category"] = categorys
	/*c.Data["json"]= &articles
	c.ServeJSON()
	c.StopRun()*/

	c.TplName = "admin/article-edit.html"

}

func (c *ArticleController) Save() {

	title := c.GetString("title")
	tag := c.GetString("tag")
	cate, _ := c.GetInt("cate", 0)
	remark := c.GetString("remark")
	desc := c.GetString("desc_content")
	html := c.GetString("desc_html")
	cover := c.GetString("cover","null")

	o := orm.NewOrm()
	article := admin.Article{
		Title:    title,
		Tag:      tag,
		Desc:     desc,
		Html:	  html,
		Remark:   remark,
		Cover:    cover,
		Status:   1,
		User:     &admin.User{1, "", "", "", time.Now(), 0},
		Category: &admin.Category{cate, "", "", 0, 0,0},
	}


	/*c.Data["json"] = &article
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(article.Title, "Title")
	valid.Required(article.Html, "Html")
	valid.Required(article.Tag, "Tag")
	valid.Required(article.Desc, "Desc")
	valid.Required(article.Remark, "Remark")

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



	if id, err := o.Insert(&article); err == nil {
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

func (c *ArticleController) Update() {
	id, _ := c.GetInt("id", 0)
	title := c.GetString("title")
	tag := c.GetString("tag")
	cate, _ := c.GetInt("cate", 0)
	remark := c.GetString("remark")
	desc := c.GetString("desc_content")
	html := c.GetString("desc_html")
	cover := c.GetString("cover","null")
	/*c.Data["json"] = c.Input()
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	o := orm.NewOrm()

	article := admin.Article{Id: id}
	if o.Read(&article) == nil {
		article.Title = title
		article.Tag = tag
		article.Desc = desc
		article.Html = html
		article.Remark = remark
		article.Cover = cover

		valid := validation.Validation{}
		valid.Required(article.Title, "Title")
		valid.Required(article.Tag, "Tag")
		valid.Required(article.Desc, "Desc")
		valid.Required(article.Html, "Html")
		valid.Required(article.Remark, "Remark")

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

		article.Category = &admin.Category{cate, "", "", 0, 0,0}

		if _, err := o.Update(&article); err == nil {
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

func (c *ArticleController) Delete() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})

	o := orm.NewOrm()
	article := admin.Article{Id: id}
	if o.Read(&article) == nil {
		article.Status = 3

		if _, err := o.Update(&article); err == nil {
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


func (c *ArticleController) Top() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})


	o := orm.NewOrm()
	article := admin.Article{Id: id}
	if o.Read(&article) == nil {

		recommend := article.Recommend
		if recommend == 0 {
			article.Recommend = 1
		}else{
			article.Recommend = 0
		}

		if _, err := o.Update(&article); err == nil {
			response["msg"] = "操作成功！"
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "操作失败！"
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