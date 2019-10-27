package routers

import (
	"PCService/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/host",
		beego.NSRouter("ip", &controllers.IPController{}),
	)
	beego.AddNamespace(ns)
}
