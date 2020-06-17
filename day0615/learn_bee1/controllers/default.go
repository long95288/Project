package controllers

import (
	"encoding/json"

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

type Hello struct {
	Name string
	Age  int
}

func (c *HelloController) Get() {

	data, _ := json.Marshal(Hello{Name: "名称", Age: 132})
	c.Ctx.WriteString(string(data))
}
