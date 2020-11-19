package controllers

import (
	"demo/models"

	"github.com/astaxie/beego"
)

type DemoController struct {
	beego.Controller
}

// @router /demo/hello [get]
func (c *DemoController) Get() {
	var (
		title string
	)
	title = "Hello World!!!"
	c.Ctx.WriteString(title)
}

// @router /user/username [get]
func (c *DemoController) GetUsername() {
	var (
		id    int
		err   error
		title string
		user  models.User
	)

	id, err = c.GetInt("id")
	user, err = models.UserInfo(id)
	if err == nil {
		title = user.Name
	} else {
		title = "error"
	}
	c.Ctx.WriteString(title)
}
