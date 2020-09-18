package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
	"kangg/utils"
	"time"
)


type UserController struct {
	BaseController
}

func (c *UserController) List() {

	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量

	name := c.GetString("name")
	status, _ := c.GetInt("status", 0)

	c.Data["Status"] = status
	c.Data["Name"] = name


	o := orm.NewOrm()

	var users []*admin.User
	qs := o.QueryTable(new(admin.User))

	// 状态
	if status != 0 {
		qs = qs.Filter("status", status)
	}

	// 名称
	if name != "" {
		qs = qs.Filter("name__icontains", name)
	}

	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&users)

	if err != nil {
		c.Abort("404")
	}

	/*c.Data["json"]= &articles
	c.ServeJSON()
	c.StopRun()*/

	// 统计
	count, err := qs.Count()
	if err != nil {
		c.Abort("404")
	}

	c.Data["Data"] = &users
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["StatusText"] = admin.Status
	c.TplName = "admin/user-list.html"
}

func (c *UserController) Add() {

	c.TplName = "admin/user-add.html"
}

func (c *UserController) Put() {
	id, err := c.GetInt("id", 0)

	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	var users []*admin.User
	qs := o.QueryTable(new(admin.User))
	err = qs.Filter("id", id).One(&users)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = &users[0]

	c.TplName = "admin/user-edit.html"

}

func (c *UserController) Save() {

	name := c.GetString("name")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	email := c.GetString("email")


	o := orm.NewOrm()
	user := admin.User{
		Name:    		name,
		Password:      	utils.PasswordMD5(password,name),
		Email:     		email,
		Created:   		time.Now(),
		Status:   1,
	}


	response := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(user.Name, "Name")
	valid.Required(user.Password, "Password")
	valid.Required(user.Email, "Email")
	valid.Email(user.Email, "Email")

	if password != repassword {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = "密码不一致"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

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



	if id, err := o.Insert(&user); err == nil {
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

func (c *UserController) Update() {
	id, _ := c.GetInt("id", 0)
	name := c.GetString("name")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	email := c.GetString("email")

	/*c.Data["json"] = c.Input()
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	o := orm.NewOrm()

	user := admin.User{Id: id}
	if o.Read(&user) == nil {
		user.Name = name
		user.Email = email
		if password != "" {
			user.Password = utils.PasswordMD5(password,name)
		}

		valid := validation.Validation{}
		valid.Required(user.Name, "Name")
		valid.Required(user.Email, "Email")
		valid.Email(user.Email, "Email")

		if password != repassword {
			response["msg"] = "新增失败！"
			response["code"] = 500
			response["err"] = "密码不一致"
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}


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

		if _, err := o.Update(&user); err == nil {
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

func (c *UserController) Delete() {
	id, _ := c.GetInt("id", 0)
	status, _ := c.GetInt("status", 1)

	response := make(map[string]interface{})

	o := orm.NewOrm()
	user := admin.User{Id: id}
	if o.Read(&user) == nil {
		if status == 1{
			status = 2
		}else{
			status = 1
		}
		user.Status = status

		if _, err := o.Update(&user); err == nil {
			response["msg"] = "禁用成功！"
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "禁用失败！"
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "禁用失败！"
		response["code"] = 500
		response["err"] = "ID 不能为空！"
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

