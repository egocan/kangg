package db

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
)

type PostgresDataBase struct {
}

func NewPostgresDataBase() DataBase {
	database := PostgresDataBase{}
	return &database
}

func (database *PostgresDataBase) GetStr() string{
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	dbUser := conf.String("db::dbUser")
	dbPass := conf.String("db::dbPass")
	dbHost := conf.String("db::dbHost")
	dbPort := conf.String("db::dbPort")
	dbName := conf.String("db::dbName")

	return "user="+dbUser+" host="+dbHost+" port="+dbPort+" password="+dbPass+" dbname="+dbName+" sslmode=disable"

}

func (database *PostgresDataBase) GetDriverName() string{
	return "postgres"
}

func (database *PostgresDataBase) GetAliasName() string{
	return "default"
}

func (database *PostgresDataBase) GetDriver() orm.DriverType{
	return orm.DRPostgres
}


func init(){
	Register("postgres",NewPostgresDataBase)
}




