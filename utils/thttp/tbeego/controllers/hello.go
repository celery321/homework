package controllers

type HelloController struct {
	BaseController
	BaseResponse
}

func (this *HelloController) Get() {
	code := 200
	msg := "1111"
	data := []string{"111", "2222"}

	content := this.BaseResponse.AddJsonFormat(code, msg, data)
	this.BaseController.ReturnFormat(content)

}
