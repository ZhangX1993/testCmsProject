package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"myapp/cmsProject/model"
)

//数据库引擎
func NewMysqlEngine() *xorm.Engine{
	engine,err:=xorm.NewEngine("mysql","root:wuzeyusb@/test2?charset=utf8")

	err=engine.CreateTables(new(model.Admin))
	err=engine.Sync2(new(model.Admin))

	if err!=nil{
		panic(err.Error())
	}

	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}