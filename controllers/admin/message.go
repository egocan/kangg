package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
	"kangg/utils"
)

type MessageController struct {
	BaseController
}

func (c *MessageController) List() {


	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量
	status, _ := c.GetInt("status", 0)

	c.Data["Status"] = status

	o := orm.NewOrm()
	message := new(admin.Message)

	var messages []*admin.Message
	qs := o.QueryTable(message)
	//qs = qs.Filter("status", status)

	if status != 0 {
		qs = qs.Filter("status", status)
	}else{
		qs = qs.Filter("status__lt", 3)
	}

	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&messages)
	if err != nil {
		c.Abort("404")
	}


	/*c.Data["json"]= &messages
	c.ServeJSON()
	c.StopRun()*/


	// 统计
	count, err := qs.Count()
	if err != nil {
		c.Abort("404")
	}

	c.Data["Data"] = &messages
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["StatusText"] = admin.Status


	c.TplName = "admin/message-list.html"
}

func (c *MessageController) Put() {
	id, err := c.GetInt("id", 0)

	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	message := new(admin.Message)
	var messages []*admin.Message
	qs := o.QueryTable(message)
	err = qs.Filter("id", id).One(&messages)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = messages[0]

	c.TplName = "admin/message-edit.html"
}

func (c *MessageController) Update() {

	id, _ := c.GetInt("id", 0)
	reply := c.GetString("reply")


	/*c.Data["json"] = c.Input()
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	o := orm.NewOrm()

	message := admin.Message{Id: id}
	if o.Read(&message) == nil {
		message.Reply = reply

		valid := validation.Validation{}
		valid.Required(message.Reply, "Reply")

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

		if _, err := o.Update(&message); err == nil {
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

func (c *MessageController) Delete() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})

	o := orm.NewOrm()
	message := admin.Message{Id: id}

	if o.Read(&message) == nil {
		message.Status = 2

		if _, err := o.Update(&message); err == nil {
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

func (c *MessageController) Check() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})


	o := orm.NewOrm()
	message := admin.Message{Id: id}
	if o.Read(&message) == nil {

		status := message.Status
		if status == 2 {
			message.Status = 1
		}else{
			message.Status = 2
		}

		if _, err := o.Update(&message); err == nil {
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