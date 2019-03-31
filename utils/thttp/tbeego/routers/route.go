package routers

import "github.com/astaxie/beego"
import "homework/utils/thttp/tbeego/controllers"

func init() {
	beego.Router("/t", &controllers.HelloController{})
}
