package main

import "github.com/astaxie/beego"

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	this.Ctx.WriteString("hello1")
}
