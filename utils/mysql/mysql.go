package mysql

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	if nil != err {
		port = 3306
	}
	logs.Info("注册数据库")
	if beego.AppConfig.String("runmode") == "dev" {
		logs.Info("应用当前运行模式：")
		orm.Debug = true

	}
	e1 := orm.RegisterDriver("mysql", orm.DRMySQL)
	if e1 != nil {
		logs.Info("注册数据库驱动失败")
	}
	e2 := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
	if e2 != nil {
		logs.Info("注册数据库失败:" + e2.Error())
	}
}
