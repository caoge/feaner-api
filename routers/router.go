package routers

import (
	"feaner/feaner-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/version/:key", &controllers.MainController{}, "get:Version")
}
