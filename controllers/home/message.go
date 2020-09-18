package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
	"kangg/utils"
	"html/template"
)

type MessageController struct {
	BaseController
}

// 详情
func (c *MessageController) Get() {


	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量


	o := orm.NewOrm()
	message := new(admin.Message)

	var messages []*admin.Message
	qs := o.QueryTable(message)
	qs = qs.Filter("status", 1)


	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&messages)
	if err != nil {
		c.Abort("404")
	}



	// 统计
	count, err := qs.Count()
	if err != nil {
		c.Abort("404")
	}

	c.Data["Data"] = &messages
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["StatusText"] = admin.Status


	c.Log("message")
	c.Data["index"] = "留言"
	c.TplName = "home/" + beego.AppConfig.String("view") + "/message.html"
}

func (c *MessageController) Save() {


	name := c.GetString("name")
	review := c.GetString("review")
	mail := c.GetString("mail","")

	o := orm.NewOrm()
	message := admin.Message{
		Name:    	template.HTMLEscapeString(name),
		Review:     template.HTMLEscapeString(review),
		Mail:    	template.HTMLEscapeString(mail),
		Status:   	2,
	}

	response := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(message.Name, "Name")
	valid.Required(message.Review, "Review")

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


	if id, err := o.Insert(&message); err == nil {
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
