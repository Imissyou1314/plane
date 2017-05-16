// @APIVersion 1.0.0
// @Title Plane API
// @Description Plane Service API
// @Contact 1255418233@qq.com
// @TermsOfServiceUrl http://127.0.0.1/miss/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"plane/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/miss",

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/log",
			beego.NSInclude(
				&controllers.LogController{},
			),
		),
		beego.NSNamespace("/plane",
			beego.NSInclude(
				&controllers.PlaneController{},
			),
		),
		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{},
			),
		),
		beego.NSNamespace("/video",
			beego.NSInclude(
				&controllers.VideoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
