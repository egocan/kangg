package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"kangg/models/admin"
	"kangg/utils"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Layout() {

	o := orm.NewOrm()

	// 月份排序
	articleTime := new(admin.Article)
	var articlesTime []*admin.Article
	nqs := o.QueryTable(articleTime)
	nqs = nqs.Filter("status", 1)
	nqs.OrderBy("-Created").RelatedSel().All(&articlesTime,"Created")
	count,_ := nqs.Count()
	var datetime = make(map[string]int64)
	var dateTimeKey []string
	for _,v := range articlesTime  {
		//str = append(str ,v.Created.Format("2006-01"))
		//c.Ctx.WriteString(v.Created.Format("2006-01"))
		k := v.Created.Format("2006-01")
		if datetime[k] == 0{
			dateTimeKey = append(dateTimeKey,k)
		}
		datetime[k] = datetime[k] + 1
	}
	c.Data["DateTime"] = datetime
	c.Data["DateTimeKey"] = dateTimeKey
	c.Data["DateCount"] = count

	// 阅读排序
	articleReadSort := new(admin.Article)
	var articlesReadSort []*admin.Article
	nqrs := o.QueryTable(articleReadSort)
	nqrs = nqrs.Filter("status", 1)
	nqrs = nqrs.OrderBy("-Pv")
	nqrs.Limit(5).All(&articlesReadSort,"Id","Title","Pv")
	c.Data["ArticlesReadSort"] = articlesReadSort
	
	// 最新评论
    review := new(admin.Review)
    var reviewData []*admin.Review
    nqrw := o.QueryTable(review)
    nqrw = nqrw.Filter("status", 1)
    nqrw = nqrw.OrderBy("-Id")
    nqrw.Limit(5).All(&reviewData,"Review","ArticleId")
	reviewCount ,_ := nqrw.Count()
	c.Data["ReviewCount"] = reviewCount
    c.Data["Review"] = reviewData
}

func (c *BaseController) Menu()  {

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	sortby = append(sortby,"sort")
	order = append(order,"asc")

	menu , _  := admin.GetAllMenu(query,fields,sortby,order,offset,limit)
	data := utils.MenuData(menu,0,0)
	/*c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()*/
	c.Data["Menu"] = data

	link , _  := admin.GetAllLink(query,fields,sortby,order,offset,limit)
	c.Data["Link"] = link

}

func (c *BaseController) Prepare(){
	c.Data["bgClass"] = "bgColor"
	c.Data["T"] = time.Now()

	o := orm.NewOrm()
	var setting []*admin.Setting
	o.QueryTable(new(admin.Setting)).All(&setting)

	for _,v := range setting{
		c.Data[v.Name] = v.Value
	}

	pv,_ := o.QueryTable(new(admin.Log)).Count()
	uv,_ := o.QueryTable(new(admin.Log)).GroupBy("ip").Count()

	c.Data["PV"] = pv
	c.Data["UV"] = uv

	c.Layout()
	c.Menu()
	c.Keywords()
	c.AWord()
	c.Banner()
	c.Ad()

}
func(c *BaseController) Ad(){
	o := orm.NewOrm()
	var ad []*admin.Ad
	o.QueryTable(new(admin.Ad)).All(&ad)

	for _ , v :=range ad{
		c.Data[v.Name] = v.Html
		c.Data[v.Name + ".Title"] = v.Title
	}
}
func (c *BaseController) Banner(){
	o := orm.NewOrm()
	banner := new(admin.Banner)

	var banners []*admin.Banner
	qs := o.QueryTable(banner)
	qs = qs.Filter("Article__Title__isnull", false)

	// 获取数据
	_, err := qs.OrderBy("id").RelatedSel().All(&banners)

	if err != nil {
		c.Abort("404")
	}

	c.Data["Banner"] = &banners
}
func (c *BaseController) AWord() {
	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量
	o := orm.NewOrm()
	aword := new(admin.AWord)

	var awords []*admin.AWord
	qs := o.QueryTable(aword)
	qs = qs.Filter("status", 1)

	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&awords)

	if err != nil {
		c.Abort("404")
	}

	// 统计
	count, err := qs.Count()
	if err != nil {
		c.Abort("404")
	}

	c.Data["AWord"] = &awords
	c.Data["WordPages"] = utils.GenPaginator(page, limit, count)
}

func (c *BaseController)Keywords()  {

	o := orm.NewOrm()
	qs := o.QueryTable(new(admin.Article))

	var tag []*admin.Article

	qs = qs.Filter("status", 1)
	qs = qs.Filter("User__Name__isnull", false)
	qs = qs.Filter("Category__Name__isnull", false)
	qs.All(&tag,"tag")



	var tags []string
	for _,v := range tag{
		tags = append(tags,strings.Split(strings.Replace(v.Tag, `，`, `,`, -1),`,`)...)
	}

	var tagsMap = make(map[string]int)

	for _, v:= range tags{
		tagsMap[v] += 1
	}

	for k , _:= range tagsMap{
		tagsMap[k] = tagsMap[k] / 5 + 15
	}

	c.Data["Tag"] = tagsMap

}

func (c *BaseController) Log(page string)  {

	ip := c.Ctx.Input.IP()

	userAgent := c.Ctx.Input.UserAgent()

	url := c.Ctx.Input.URI()
	o := orm.NewOrm()
	var log = admin.Log{
		Ip:    			ip,
		//City:     		city,
		UserAgent:    	userAgent,
		Page:			page,
		Uri:			url,
	}
	o.Insert(&log)

}


