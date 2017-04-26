package controllers

import (
	"encoding/json"
	"plane/app/models"
	"plane/app/service"

	"github.com/astaxie/beego"
)

// Plane Users API
type UserController struct {
	beego.Controller
}

func (u *UserController) URLMapping() {
	u.Mapping("CreateUser", u.CreateUser)
	u.Mapping("GetAll", u.GetAll)
	u.Mapping("UpdateUser", u.UpdateUser)
	u.Mapping("GetUserById", u.GetUserById)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /createUser [post]
func (u *UserController) CreateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	result, _ := service.UserService.AddUser(&user)
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetAll user
// @Description get all Users
// @Success 200 {object} []models.User
// @Failure 500 系统错误
// @router /user/getAll [get]
func (u *UserController) GetAll() {
	users, _ := service.UserService.GetAllUser()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title GetUser By user Id
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /user/:uid [get]
func (u *UserController) GetUserById() {
	uid, _ := u.GetInt64(":uid")
	if uid != 0 {
		user, err := service.UserService.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /user/update [put]
func (u *UserController) UpdateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uu, err := service.UserService.UpdateUser(&user, "Id")
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = uu
	}
	u.ServeJSON()
}
