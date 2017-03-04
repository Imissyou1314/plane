package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["plane/app/controllers:LogController"] = append(beego.GlobalControllerRouter["plane/app/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetLogs",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
