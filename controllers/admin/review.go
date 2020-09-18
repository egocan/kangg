package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
	"kangg/utils"
	"time"
)

type ReviewController struct {
	BaseController
}

type Item struct {
	Id       	int
	Name    	string
	Review   	string
	Reply     	string
	Mail   		string
	Created  	time.Time
	Updated  	time.Time
	Status   	int
	Title	 	string
	Aid	 		int
}

func (c *ReviewController) List() {

	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量
	status, _ := c.GetInt("status", 0)

	c.Data["Status"] = status

	response := make(map[string]interface{})

	o := orm.NewOrm()
	review := new(admin.Review)

	var reviews []*admin.Review
	qs := o.QueryTable(review)

	if status != 0 {
		qs = qs.Filter("status", status)
	}else{
		qs = qs.Filter("status__lt", 3)
	}
	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&reviews)
	if err != nil {
		response["msg"] = "Error."
		response["code"] = 500
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}


	var item []Item

	for  _,v := range reviews{

		article := new(admin.Article)
		var articles []*admin.Article
		qs := o.QueryTable(article)
		err = qs.Filter("id", v.ArticleId).One(&articles,"Title","Id")

		item = append(item,Item{
			v.Id,
			v.Name,
			v.Review,
			v.Reply,
			v.Mail,
			v.Created,
			v.Updated,
			v.Status,
			articles[0].Title,
			articles[0].Id,
		})

	}

	// 统计
	count, err := qs.Count()
	if err != nil {
		response["msg"] = "Error."
		response["code"] = 500
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}


	c.Data["Data"] = &item
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["StatusText"] = admin.Status
	c.TplName = "admin/review-list.html"
}


func (c *ReviewController) Put() {
	id, err := c.GetInt("id", 0)

	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	review := new(admin.Review)
	var messages []*admin.Review
	qs := o.QueryTable(review)
	err = qs.Filter("id", id).One(&messages)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = messages[0]

	c.TplName = "admin/review-edit.html"
}

func (c *ReviewController) Update() {

	id, _ := c.GetInt("id", 0)
	reply := c.GetString("reply")


	/*c.Data["json"] = c.Input()
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	o := orm.NewOrm()

	review := admin.Review{Id: id}
	if o.Read(&review) == nil {
		review.Reply = reply

		valid := validation.Validation{}
		valid.Required(review.Reply, "Reply")

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

		if _, err := o.Update(&review); err == nil {
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

func (c *ReviewController) Delete() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})

	o := orm.NewOrm()
	review := admin.Review{Id: id}

	if o.Read(&review) == nil {
		review.Status = 2

		if _, err := o.Update(&review); err == nil {
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

func (c *ReviewController) Check() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})


	o := orm.NewOrm()
	review := admin.Review{Id: id}
	if o.Read(&review) == nil {

		status := review.Status
		if status == 2 {
			review.Status = 1
		}else{
			review.Status = 2
		}

		if _, err := o.Update(&review); err == nil {
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