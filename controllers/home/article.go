package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kangg/models/admin"
	"kangg/utils"
	"html/template"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	BaseController
}
// 类表
func (c *ArticleController) List() {


	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量
	categoryId, _ := c.GetInt("c", 0)           // 页数
	o := orm.NewOrm()
	article := new(admin.Article)

	var articles []*admin.Article
	qs := o.QueryTable(article)
	qs = qs.Filter("status", 1)
	qs = qs.Filter("User__Name__isnull", false)
	qs = qs.Filter("Category__Name__isnull", false)

	if categoryId != 0 {

		category := new(admin.Category)
		var categorys []*admin.Category
		cqs := o.QueryTable(category)
		cqs = cqs.Filter("status", 1)
		cqs.OrderBy("-sort").All(&categorys)

		ids := utils.CategoryTreeR(categorys,categoryId,0)

		var cids []int
		cids = append(cids,categoryId)
		for _,v := range ids{
			cids = append(cids,v.Id)
		}

		/*c.Data["json"] = cids
		c.ServeJSON()
		c.StopRun()*/

		qs = qs.Filter("Category__ID__in", cids)


	}
	c.Data["CategoryID"] = &categoryId
	// 查出当前分类下的所有子分类id

	date := c.GetString("date")
	if date != "" {
		if len(date) == 7 {
			start := ""
			end := ""
			dateNumStr := beego.Substr(date, len("2018-"),2)
			yearNumStr := beego.Substr(date, len("20"),2)
			dateNum,_:=strconv.Atoi(dateNumStr)
			yearNum,_:=strconv.Atoi(yearNumStr)

			start = utils.SubString(date, len("2018-01"))+"-01 00:00:00"
			if dateNum >= 12 {
				endYearStr := strconv.Itoa(yearNum+1)
				end = utils.SubString(date, len("20"))+endYearStr+"-01-01 00:00:00"
			}

			if dateNum < 9 {
				endStr := strconv.Itoa(dateNum+1)
				end = utils.SubString(date, len("2018-0"))+endStr+"-01 00:00:00"
			}
			if dateNum >= 9 && dateNum < 12 {
				endStr := strconv.Itoa(dateNum+1)
				end = utils.SubString(date, len("2018-"))+endStr+"-01 00:00:00"
			}

			/*c.Data["json"] = []string{start,end}
			c.ServeJSON()
			c.StopRun()*/

			qs = qs.Filter("created__gte", start)
			qs = qs.Filter("created__lte", end)
			c.Data["Date"] = utils.SubString(start, len("2018-01"))

		}else {
			date = utils.SubString(date, len("2018-01-01"))
			tm, _ := time.Parse("2006-01-02", date)
			unix := tm.Unix() //1566432000

			startFormat := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
			moreUnix, _ := utils.ToInt64(int64(60 * 60 * 24))
			endFormat := time.Unix(unix+moreUnix, 0).Format("2006-01-02 15:04:05")
			start := utils.SubString(startFormat, len("2018-01-01")) + " 00:00:00"
			end := utils.SubString(endFormat, len("2018-01-01")) + " 00:00:00"

			// 刷选
			qs = qs.Filter("created__gte", start)
			qs = qs.Filter("created__lte", end)
			c.Data["Date"] = utils.SubString(start, len("2018-01-01"))
		}
	}


	tag := c.GetString("tag")
	if tag != "" {
		qs = qs.Filter("tag__icontains", tag)
	}

	cate := c.GetString("cate")
	if cate != ""{
		qs = qs.Filter("Category__Ename",cate)
	}

	// 统计
	count, err := qs.Count()
	if err != nil {
		c.Abort("404")
	}


	// 获取数据
	_, err = qs.OrderBy("-recommend","-id","-pv").RelatedSel().Limit(limit).Offset(offset).All(&articles)
	if err != nil {
		panic(err)
	}
	c.Data["Data"] = &articles
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)


	// Menu
	c.Log("article")

	//if categoryId == 0 {
	//	c.Data["index"] = "我的心理学过程记录"
	//}else{
	//	categoryKey := admin.Category{Id:categoryId}
	//	err = o.Read(&categoryKey)
	//	if err == nil {
	//		c.Data["index"] = categoryKey.Name
	//	}else{
	//		c.Data["index"] = "我的心理学过程记录"
	//	}
	//}
	switch{
	case tag != "":
		c.Data["index"] = tag
	case cate != "":
		categoryKey := admin.Category{Ename:cate}
		err = o.Read(&categoryKey,"Ename")
		if err == nil {
			c.Data["index"] = categoryKey.Name
		}else{
			c.Data["index"] = beego.AppConfig.String("title")
		}
	case categoryId !=0:
		categoryKey := admin.Category{Id:categoryId}
		err = o.Read(&categoryKey)
		if err == nil {
			c.Data["index"] = categoryKey.Name
		}else{
			c.Data["index"] = beego.AppConfig.String("title")
		}
	default:
		c.Data["index"] = beego.AppConfig.String("title")

	}

	c.TplName = "home/" + beego.AppConfig.String("view") + "/list.html"
}

// 详情
func (c *ArticleController) Detail() {

	id := c.Ctx.Input.Param(":id")
	viewType := c.GetString("type")
	// 基础数据
	o := orm.NewOrm()
	article := new(admin.Article)
	var articles []*admin.Article
	qs := o.QueryTable(article)
	err := qs.Filter("id", id).RelatedSel().One(&articles)
	if err != nil {
		c.Abort("404")
	}

	/*c.Data["json"]= &articles
	c.ServeJSON()
	c.StopRun()*/

	c.Data["Data"] = &articles[0]


	if  beego.AppConfig.String("view") == "default" {
		var listData = make(map[string][]*admin.Article)
		var list []*admin.Article
		_, err = o.QueryTable(article).Filter("status", 1).Filter("User__Name__isnull", false).Filter("Category__Name__isnull", false).OrderBy("id").RelatedSel().All(&list, "id", "title")

		for _,v := range list{
			listData[v.Category.Name] = append(listData[v.Category.Name],v)
		}
		c.Data["List"] = &listData
		articleId, _ := strconv.Atoi(id)
		c.Data["ArticleId"] = articleId
		/*c.Data["json"]= &listData
		c.ServeJSON()
		c.StopRun()*/

	}

	c.Log("detail")
	c.Data["index"] = &articles[0].Title

	if viewType == "single"{
		c.TplName = "home/" + beego.AppConfig.String("view") + "/doc.html"
	}else {
		c.TplName = "home/" + beego.AppConfig.String("view") + "/detail.html"
	}
}

// 统计访问量
func (c *ArticleController) Pv()  {

	ids := c.Ctx.Input.Param(":id")
	id,_:=strconv.Atoi(ids)
	/*c.Data["json"] = c.Input()
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	o := orm.NewOrm()

	article := admin.Article{Id: id}
	if o.Read(&article) == nil {
		article.Pv = article.Pv + 1

		valid := validation.Validation{}
		valid.Required(article.Id, "Id")
		if valid.HasErrors() {
			// 如果有错误信息，证明验证没通过
			// 打印错误信息
			for _, err := range valid.Errors {
				//log.Println(err.Key, err.Message)
				response["msg"] = "Error."
				response["code"] = 500
				response["err"] = err.Key + " " + err.Message
				c.Data["json"] = response
				c.ServeJSON()
				c.StopRun()
			}
		}

		if _, err := o.Update(&article); err == nil {
			response["msg"] = "Success."
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "Error."
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "Error."
		response["code"] = 500
		response["err"] = "ID 不能为空！"
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

// 评论
func (c *ArticleController) Review() {

	aid, _ := c.GetInt("aid")
	name := c.GetString("name")
	review := c.GetString("review")
	mail := c.GetString("mail","")

	o := orm.NewOrm()
	reviewsMd := admin.Review{
		Name:    	template.HTMLEscapeString(name),
		Review:     template.HTMLEscapeString(review),
		Mail:    	template.HTMLEscapeString(mail),
		ArticleId:	aid,
		Status:   	2,
	}

	response := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(reviewsMd.Name, "Name")
	valid.Required(reviewsMd.Review, "Review")
	valid.Required(reviewsMd.ArticleId, "ArticleId")

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

	// 更新评论数量
	article:= admin.Article{Id: aid}
	o.Read(&article)
	article.Review = article.Review + 1
	o.Update(&article)


	if id, err := o.Insert(&reviewsMd); err == nil {
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

func (c *ArticleController) ReviewList()  {

	id := c.Ctx.Input.Param(":id")
	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量
	response := make(map[string]interface{})

	o := orm.NewOrm()
	review := new(admin.Review)

	var reviews []*admin.Review
	qs := o.QueryTable(review)
	qs = qs.Filter("status", 1)
	qs = qs.Filter("article_id", id)

	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&reviews)

	if err != nil {
		response["msg"] = "Error."
		response["code"] = 500
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
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

	response["Data"] = &reviews
	response["Paginator"] = utils.GenPaginator(page, limit, count)

	response["msg"] = "Success."
	response["code"] = 200
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()

}

func (c *ArticleController) Like()  {

	response := make(map[string]interface{})
	ip := c.Ctx.Input.IP()
	id,_ := c.GetInt("id")

	o := orm.NewOrm()
	qs := o.QueryTable(new(admin.Log))

	qs = qs.Filter("ip",ip)
	qs = qs.Filter("create__gte",beego.Date(time.Now(),"Y-m-d 00:00:00"))
	qs = qs.Filter("create__lte",beego.Date(time.Now(),"Y-m-d H:i:s"))
	qs = qs.Filter("page","like" + strconv.Itoa(id))

	count, e := qs.Count()
	if e != nil {
		response["msg"] = "Error."
		response["code"] = 500
		response["err"] = e.Error()
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}
	if count >= 1 {
		response["msg"] = "Error."
		response["code"] = 500
		response["err"] = "亲，点赞过了，明天再来哦！"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	c.Log("like" + strconv.Itoa(id))

	article := admin.Article{Id: id}
	if o.Read(&article) == nil {
		article.Like = article.Like + 1

		valid := validation.Validation{}
		valid.Required(article.Id, "Id")
		if valid.HasErrors() {
			// 如果有错误信息，证明验证没通过
			// 打印错误信息
			for _, err := range valid.Errors {
				//log.Println(err.Key, err.Message)
				response["msg"] = "Error."
				response["code"] = 500
				response["err"] = err.Key + " " + err.Message
				c.Data["json"] = response
				c.ServeJSON()
				c.StopRun()
			}
		}

		if _, err := o.Update(&article); err == nil {
			response["msg"] = "Success."
			response["code"] = 200
			response["like"] = article.Like
		} else {
			response["msg"] = "Error."
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "Error."
		response["code"] = 500
		response["err"] = "ID 不能为空！"
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *ArticleController) Search(){
	var searchEscapePattern = []string{
		`\\`, `(`, `)`, `|`, `-`, `!`, `@`, `~`, `'`, `&`, `/`, `^`, `$`, `=`,
		`\\\\`, `\(`, `\)`, `\|`, `\-`, `\!`, `\@`, `\~`, `\'`, `\&`, `\/`, `\^`, `\$`, `\=`,
	}
	q := c.GetString("query")
	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量

	q = strings.TrimSpace(q)
	q = strings.NewReplacer(searchEscapePattern...).Replace(q)


	if len(q) != 0 {
		o := orm.NewOrm()
		cond := orm.NewCondition()
		article := new(admin.Article)

		var articles []*admin.Article
		qs := o.QueryTable(article)
		qs = qs.Filter("status", 1)
		qs = qs.Filter("User__Name__isnull", false)
		qs = qs.Filter("Category__Name__isnull", false)

		cdt := cond.Or("title__icontains",q).Or("desc__icontains",q).Or("html__icontains",q)
		qs = qs.SetCond(cdt).Distinct()


		// 统计
		count, err := qs.Count()
		if err != nil {
			c.Abort("404")
		}

		// 获取数据
		_, err = qs.OrderBy("-id").RelatedSel().Limit(limit).Offset(offset).All(&articles)
		if err != nil {
			panic(err)
		}


		c.Data["Data"] = &articles
		c.Data["Query"] = q
		c.Data["Paginator"] = utils.GenPaginator(page, limit, count)

		c.TplName = "home/" + beego.AppConfig.String("view") + "/search.html"

	} else {
		c.List()
	}


}