package routers

import (
	"demo/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.MainController{},"get:GetHello")
    beego.Include(&controllers.DemoController{})


    var FilterDemo = func(ctx * context.Context) {
    	var (
    		title string
		)
		title ="禁止访问"
		ctx.WriteString(title)
	}
    beego.InsertFilter("/demo/*",beego.BeforeRouter,FilterDemo)

}
