package main

import (
	_ "HelloBeego06/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//定义config变量，接收并赋值为全局配置变量
	config := beego.AppConfig
	//获取配置选项
	appName :=config.String("appname")
	fmt.Println("项目应用名称:",appName)
	port,err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目配置信息解析错误，请查验后重试")
	}
	fmt.Println("应用监听端口:",port)
	beego.Run(":7076" )

	driver := config.String("db_driver")//数据库驱动
	dbUser := config.String("db_user")//数据库用户名
	dbPassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_Name")
	db, err :=sql.Open(driver,dbUser+":"+dbPassword+"@tcp("+dbIP+")/"+dbName)
	if err != nil {
		panic("数据库连接打开失败，请重试")
	}
	fmt.Println(db)
}

