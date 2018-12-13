package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Info struct {
	Website string
	Email   int
}

func (c *MainController) Get() {
	mystruct := Info{Website: "beego.me", Email: 11}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}

func (c *MainController) Version() {
	email, _ := c.GetInt(":key")
	c.Data["json"] = Info{Website: "feaner.com", Email: email}
	c.ServeJSON()
}
