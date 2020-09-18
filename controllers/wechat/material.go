package wechat

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/silenceper/wechat"
	wcMaterial "github.com/silenceper/wechat/material"
	"kangg/models/admin"
	"strconv"
)

type MaterialController struct {
	BaseController
}

func (c *MaterialController)AddNews()  {
	response := make(map[string]interface{})

	wc := wechat.NewWechat(config)
	string, error := wc.GetAccessToken()

	response["err"] = error.Error()
	response["ids"] = string
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()

	material := wc.GetMaterial()
	ids,e := material.ImageUpload("E:/project/go/src/kangg/static/uploads/20190830155527847.jpg")




	response["msg"] = "获取失败！"
	response["code"] = 500
	response["err"] = e.Error()
	response["ids"] = ids

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()




	id, _ := c.GetInt("id", 0)
	var articles []*admin.Article
	o := orm.NewOrm()
	err := o.QueryTable(new(admin.Article)).RelatedSel().Filter("id",id).One(&articles)
	data := articles[0]

	if err != nil{
		response["msg"] = "获取失败！"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}
	//wc := wechat.NewWechat(config)
	//material := wc.GetMaterial()

	var article []*wcMaterial.Article
	var url = beego.AppConfig.String("url")
	article = append(article,&wcMaterial.Article{
		Title:data.Title,
		ThumbMediaID:"12",
		Author:data.User.Name,
		Digest:data.Remark,
		ShowCoverPic:1,
		Content:data.Desc,
		ContentSourceURL: url + "/detail/" + strconv.Itoa(id)+ ".html",
		URL:"",
		DownURL:"",
	})
	mid,err := material.AddNews(article)

	if err != nil{
		response["msg"] = "获取失败！"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}
	response["msg"] = "获取失败！"
	response["code"] = 200
	response["id"] = mid
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
