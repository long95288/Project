package routers

import (
	"myproject/go_learn/day0615/learn_bee1/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/Hello", &controllers.HelloController{})
	beego.Router("/", &controllers.MainController{})
}
