package controllers

import (
	"feaner/feaner-api/components"
	"github.com/astaxie/beego"
	"strconv"
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

	span := components.Trace.StartSpan("new_span")
	defer span.Finish()
	span.SetOperationName("span_1")

	span.SetBaggageItem("Some_Key", "12345")
	span.SetBaggageItem("key", strconv.Itoa(email))

	span.Finish()

	c.Data["json"] = Info{Website: "feaner.com", Email: email}
	c.ServeJSON()
}
