package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["plane/app/controllers:LogController"] = append(beego.GlobalControllerRouter["plane/app/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetLog",
			Router: `/getlog/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:LogController"] = append(beego.GlobalControllerRouter["plane/app/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetAllLogs",
			Router: `/getAllLogs`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "GetPlaneById",
			Router: `/plane`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "GetPlaneByUserId",
			Router: `/user/plane`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/user/createUser`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/user/getAll`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserById",
			Router: `/user/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateUser",
			Router: `/user/update`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
