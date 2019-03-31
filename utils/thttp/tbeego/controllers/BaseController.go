package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ReturnFormat(content []byte) {
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Body(content)
}

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (this *BaseResponse) AddJsonFormat(code int, msg string, data interface{}) []byte {
	this.Code = code
	this.Msg = msg
	this.Data = data
	content, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return content

}
