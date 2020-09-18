package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type DataBase interface {
	GetStr() 			string
	GetDriverName() 	string
	GetAliasName() 		string
	GetDriver() 		orm.DriverType
}

type Instance func() DataBase

var adapters = make(map[string]Instance)

func NewDataBase(adapterName string) (adapter DataBase, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("dataBase: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	if err != nil {
		adapter = nil
	}
	return
}

func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("dataBase: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("dataBase: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}
