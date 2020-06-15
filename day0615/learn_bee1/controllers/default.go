package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type HelloController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *HelloController) Get() {
	c.Data["title"] = "Hello"
	c.Data["number"] = 10
	c.TplName = "index.tpl"
}
