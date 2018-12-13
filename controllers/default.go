package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Info struct {
	Website string
	Email   string
}

func (c *MainController) Get() {
	mystruct := Info{Website: "beego.me", Email: "astaxie@gmail.com"}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}
