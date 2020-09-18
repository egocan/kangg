package wechat

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/silenceper/wechat"
	wcUser "github.com/silenceper/wechat/user"
	"kangg/models/admin"
	"kangg/utils"
	"time"
)

type UserController struct {
	BaseController
}

func (c *UserController)GetUser(){
	response := make(map[string]interface{})
	var list []*wcUser.Info
	wc := wechat.NewWechat(config)
	user := wc.GetUser()
	ds, err := user.ListUserOpenIDs()

	if err != nil{
		response["msg"] = "获取失败！"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	for _,v := range ds.Data.OpenIDs{
		info,err := user.GetUserInfo(v)
		if err != nil{
			response["msg"] = "更新用户信息失败！"
			response["code"] = 500
			response["err"] = err.Error()
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
		list = append(list,info)
	}


	var insert []admin.WechatUser
	o := orm.NewOrm()
	err = o.Begin()

	var i = 0
	for _,v:= range list{
		i++
		insert = append(insert,admin.WechatUser{
			Openid:        v.OpenID,
			Nickname:      v.Nickname,
			Sex:           v.Sex,
			City:          v.City,
			Country:       v.Country,
			Province:      v.Province,
			Language:      v.Language,
			Headimgurl:    v.Headimgurl,
			SubscribeTime: v.SubscribeTime,
			Created:       time.Time{},
		})
		// 先清除数据
		o.QueryTable("wechat_user").Filter("openid",v.OpenID).Delete()
	}

	multi, err := o.InsertMulti(i, insert)

	if err != nil{
		err = o.Rollback()
		response["msg"] = "用户信息插入失败！"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}else{
		err = o.Commit()
	}

	response["msg"] = "获取成功！"
	response["code"] = 200
	response["data"] = multi
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()

}

func (c *UserController)List()  {

	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量

	o := orm.NewOrm()
	var users []*admin.WechatUser
	qs := o.QueryTable(new(admin.WechatUser))
	qs.Limit(limit).Offset(offset).All(&users)
	count,_ := qs.Count()
	c.Data["User"] = users
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["Sex"] = admin.Sex
	c.TplName = "admin/wechat-user-list.html"
}