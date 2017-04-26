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
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:MessageController"] = append(beego.GlobalControllerRouter["plane/app/controllers:MessageController"],
		beego.ControllerComments{
			Method: "UserUNReadMessage",
			Router: `/getUserMessage`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:MessageController"] = append(beego.GlobalControllerRouter["plane/app/controllers:MessageController"],
		beego.ControllerComments{
			Method: "GetMessageById",
			Router: `/:messageId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:MessageController"] = append(beego.GlobalControllerRouter["plane/app/controllers:MessageController"],
		beego.ControllerComments{
			Method: "GetMessageByPlaneId",
			Router: `/plane/:planeId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "CreatePlane",
			Router: `/createPlane`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "GetPlaneById",
			Router: `/:planeId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "GetPlaneByUserId",
			Router: `/user/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:PlaneController"] = append(beego.GlobalControllerRouter["plane/app/controllers:PlaneController"],
		beego.ControllerComments{
			Method: "GetAllPlane",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["plane/app/controllers:UserController"] = append(beego.GlobalControllerRouter["plane/app/controllers:UserController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/createUser`,
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
