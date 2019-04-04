package controllers

import "homework/utils/thttp/tbeego/models/repository/db"

type HelloController struct {
	BaseController
	BaseResponse
}

func (this *HelloController) Get() {
	code := 200
	msg := "1111"
	data := []string{"111", "2222"}

	d := db.GetSession()

	content := this.BaseResponse.AddJsonFormat(code, msg, data)
	this.BaseController.ReturnFormat(content)
}
