package main

import (
	_ "HelloBeego06/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run(":8090" )
}

