package utils

import "github.com/astaxie/beego"

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	this.Ctx.WriteString("hello world")
}

func Beegohello() {
	beego.Router("/", &HelloController{})
	beego.Run()
}
