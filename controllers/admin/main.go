package admin

import (
	"github.com/astaxie/beego/orm"
	"kangg/utils"
	"kangg/utils/sys"
	"time"
)


type Log struct {
	Num 	int
	Date 	string
	Ip 		[]string
}

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.TplName = "admin/index.html"
}

func (c *MainController) Welcome() {


	o := orm.NewOrm()

	var log []orm.Params

	// time.Now().AddDate(0,0,-7) 七天前
	o.Raw("SELECT `id`,`ip`,`create` FROM `log` where `create` >= ? ORDER BY `create` ASC",time.Now().AddDate(0,0,-30)).Values(&log)

	var pv = make(map[string]Log)
	var uv = make(map[string]Log)
	var dateSlice []string
	var pvSlice []int
	var uvSlice []int

	var keys []string

	for _,v := range log {
		//utils.StringToTime(v["create"])
		// 获取日期
		var key = utils.StringToTime(v["create"]).Format("2006-01-02")

		var flag = true
		for _,k := range keys {
			if k == key {
				flag = false
			}
		}
		if flag {
			keys = append(keys,key)
		}


		// 统计pv
		pvTemp := pv[key]
		pv[key] = Log{
			Num:pvTemp.Num + 1,
		}
		// 统计uv
		uvTemp := uv[key]
		var uvFlag = true
		for _,k := range uvTemp.Ip {
			if k == v["ip"] {
				uvFlag = false
			}
		}
		if uvFlag {
			uv[key] = Log{
				Num: uvTemp.Num + 1,
				Ip:  append(uvTemp.Ip, v["ip"].(string)),
			}
		}

	}

	for _,k := range keys{
		dateSlice = append(dateSlice,k)
		pvSlice = append(pvSlice,pv[k].Num)
		uvSlice = append(uvSlice,uv[k].Num)
	}

	c.Data["Date"] = dateSlice
	c.Data["Pv"] = pvSlice
	c.Data["Uv"] = uvSlice

	df,_:= sys.Df()
	c.Data["Df"] = df

	c.TplName = "admin/welcome.html"
}
