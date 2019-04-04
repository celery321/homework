package controllers

import (
	"github.com/astaxie/beego/orm"
	"homework/utils/thttp/tbeego/log"
)
import "homework/utils/thttp/tbeego/models"

type HelloController struct {
	BaseController
	BaseResponse
}

func (this *HelloController) Get() {
	code := 200
	msg := "1111"
	data := []string{"111", "2222"}

	l := log.Log()
	o1 := orm.NewOrm()
	o1.Using("test1")
	user1 := models.User1{Name: "111111"}
	id, err := o1.Insert(&user1)

	l.Info("id: %d, err： %v", id, err)

	o2 := orm.NewOrm()
	o2.Using("test2")
	user2 := models.User2{Name: "22222"}
	id2, err2 := o2.Insert(&user2)
	l.Info("id: %d, err： %v", id2, err2)

	content := this.BaseResponse.AddJsonFormat(code, msg, data)
	this.BaseController.ReturnFormat(content)
}
