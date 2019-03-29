package main

import (
	"github.com/astaxie/beego"
	_ "github.com/beego/bee"
	)


func main() {
	beego.Router("/", &controllers.HelloController{})
	beego.Run()
}
