package routers

import (
	"feaner/feaner-api/components"
	"feaner/feaner-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	components.InitTrace()

	beego.Router("/", &controllers.MainController{})
	beego.Router("/version/:key", &controllers.MainController{}, "get:Version")
}
