package mysql

import (
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
)

var demourl = ""

func init() {
	//demourl = beego.AppConfig.String("serverName")
	//logs.Info("运行推送访问域名: " + demourl)
	//InitSql()
}

func GetDemoUrl() string {
	return demourl
}
