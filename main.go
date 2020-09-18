package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	_ "kangg/routers"
	db "kangg/service/databsae"
	"kangg/utils"
)

func init() {
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	database,_ := db.NewDataBase(conf.String("db::dbType"))
	orm.RegisterDriver(database.GetDriverName(), database.GetDriver())
	orm.RegisterDataBase(database.GetAliasName(), database.GetDriverName(), database.GetStr())

	beego.AddFuncMap("IndexForOne", utils.IndexForOne)
	beego.AddFuncMap("IndexAddOne",utils.IndexAddOne)
	beego.AddFuncMap("IndexDecrOne",utils.IndexDecrOne)
	beego.AddFuncMap("StringReplace",utils.StringReplace)
	beego.AddFuncMap("TimeStampToTime",utils.TimeStampToTime)

}

func main() {


	beego.Run()
}
