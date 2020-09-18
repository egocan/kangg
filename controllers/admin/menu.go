package admin

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"kangg/models/admin"
	"kangg/utils"
	"strconv"
	"strings"
)

// MenuController operations for Menu
type MenuController struct {
	BaseController
}

// URLMapping ...
func (c *MenuController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Add", c.Add)
}


// @router /menu/add [get]
func (c *MenuController) Add() {

	o := orm.NewOrm()

	var menus []*admin.Menu
	qs := o.QueryTable(new(admin.Menu))
	qs.All(&menus)

	c.Data["Menu"] = menus
	c.TplName = "admin/menu-add.html"
}
// Post ...
// @Title Post
// @Description create Menu
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 201 {int} models.Menu
// @Failure 403 body is empty
// @router /menu/save [post]
func (c *MenuController) Post() {

	response := make(map[string]interface{})

	sort , _ := c.GetInt("sort")
	pid , _ := c.GetInt("pid")

	if _, err := admin.AddMenu(&admin.Menu{
		Title:  c.GetString("title"),
		Target: c.GetString("target"),
		Url:    c.GetString("url"),
		Sort:   sort,
		Pid:    pid,
	}); err == nil {
		response["msg"] = "新增成功！"
		response["code"] = 200
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = response
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Menu by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Menu
// @Failure 403 :id is empty
// @router /menu/:id [get]
func (c *MenuController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, _ := admin.GetMenuById(id)
	c.Data["Data"] = v

	o := orm.NewOrm()
	var menus []*admin.Menu
	qs := o.QueryTable(new(admin.Menu))
	qs.All(&menus)
	c.Data["Menu"] = menus

	c.TplName = "admin/menu-edit.html"
}

// GetAll ...
// @Title Get All
// @Description get Menu
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Menu
// @Failure 403
// @router /menu [get]
func (c *MenuController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, _ := admin.GetAllMenu(query, fields, sortby, order, offset, limit)

	tree := utils.MenuTreeR(l, 0, 0)
	c.Data["Data"] = tree
	c.TplName = "admin/menu.html"
	/*if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
	c.StopRun()*/
}

// Put ...
// @Title Put
// @Description update the Menu
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 200 {object} models.Menu
// @Failure 403 :id is not int
// @router /menu/:id [put]
func (c *MenuController) Put() {
	response := make(map[string]interface{})

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := admin.Menu{Id: id}
	v.Target = c.GetString("target")
	v.Url = c.GetString("url")
	v.Title = c.GetString("title")
	v.Pid,_= c.GetInt("pid")
	v.Sort,_ = c.GetInt("sort")
	if err := admin.UpdateMenuById(&v); err == nil {
		response["msg"] = "修改成功！"
		response["code"] = 200
		c.Data["json"] = response
	} else {
		response["msg"] = "修改失败！"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Menu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /menu/:id [delete]
func (c *MenuController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := admin.DeleteMenu(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
