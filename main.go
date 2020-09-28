package main

import (
	"HelloBeego06/db_mysql"
	_ "HelloBeego06/routers"
	"github.com/astaxie/beego"
)

func main() {

	/*获取配置选项
	appName :=config.String("appname")
	fmt.Println("项目应用名称:",appName)
	port,err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目配置信息解析错误，请查验后重试")
	}
	fmt.Println("应用监听端口:",port)

	 */
	db_mysql.Connect()
	beego.Run(":7076" )


}

