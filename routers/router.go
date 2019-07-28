package routers

import (
	"github.com/astaxie/beego"
	"go-JWT/controllers"
	"go-JWT/filter"

)

func init() {
	// jwt 路由过滤
	beego.InsertFilter("/api/v1/:*FT_:*", beego.BeforeRouter,filter.AuthFilter)
	beego.Router("/",&controllers.MainController{},"get:Gen")

	ns :=
		beego.NewNamespace("/api",
			beego.NSNamespace("/v1",
				// 登录
				beego.NSRouter("/login", &controllers.MainController{}, "post:Login"),
				// 平台用户
				beego.NSNamespace("/user",
					// 修改密码
					beego.NSRouter("/FT_xiugai",&controllers.MainController{},"get:XiuGai"),
					beego.NSRouter("/xiugai",&controllers.MainController{},"get:XiuGai"),
			),
		))
	beego.AddNamespace(ns)
}
