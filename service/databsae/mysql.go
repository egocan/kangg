package db

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
)

type MysqlDataBase struct {
}

func NewMysqlDataBase() DataBase {
	database := MysqlDataBase{}
	return &database
}

func (database *MysqlDataBase) GetStr() string{
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	dbUser := conf.String("db::dbUser")
	dbPass := conf.String("db::dbPass")
	dbHost := conf.String("db::dbHost")
	dbPort := conf.String("db::dbPort")
	dbName := conf.String("db::dbName")

	return dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName+ "?charset=utf8&loc=Asia%2FShanghai"

}

func (database *MysqlDataBase) GetDriverName() string{
	return "mysql"
}

func (database *MysqlDataBase) GetAliasName() string{
	return "default"
}

func (database *MysqlDataBase) GetDriver() orm.DriverType{
	return orm.DRMySQL
}


func init(){
	Register("mysql",NewMysqlDataBase)
}




