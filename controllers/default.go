package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["isEmail"] = 0
	pages := []struct {
		Num int
	}{{10}, {20}, {30}}
	c.Data["Pages"] = pages
	c.TplName = "index.tpl"
}

func (c *MainController) GetHello() {
	var (
		title string
	)
	title = "Hello World"
	c.Ctx.WriteString(title)
}
