package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["plane/app/controllers:LogController"] = append(beego.GlobalControllerRouter["plane/app/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetLog",
			Router: `/GetLog/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:LogController"] = append(beego.GlobalControllerRouter["plane/app/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetAllLogs",
			Router: `/getAllLogs`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:MessageController"] = append(beego.GlobalControllerRouter["plane/app/controllers:MessageController"],
		beego.ControllerComments{
			Method: "CreateMessage",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:MessageController"] = append(beego.GlobalControllerRouter["plane/app/controllers:MessageController"],
		beego.ControllerComments{
			Method: "ReadMessage",
			Router: `/read/:messageId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:MessageController"] = append(beego.GlobalControllerRouter["plane/app/controllers:MessageController"],
		beego.ControllerComments{
			Method: "UserUNReadMessage",
			Router: `/getUserMessage`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:MessageController"] = append(beego.GlobalControllerRouter["plane/app/controllers:MessageController"],
		beego.ControllerComments{
			Method: "GetMessageByID",
			Router: `/:messageId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:MessageController"] = append(beego.GlobalControllerRouter["plane/app/controllers:MessageController"],
		beego.ControllerComments{
			Method: "GetMessageByPlaneID",
			Router: `/plane/:planeId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "CreatePlane",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "GetPlaneByID",
			Router: `/:planeId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "GetPlaneByUserID",
			Router: `/user/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "GetAllPlane",
			Router: `/getAll`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getAll`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserByID",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateUser",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/delete`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
