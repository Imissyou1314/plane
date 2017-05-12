package controllers

import (
	"encoding/json"
	"plane/app/models"
	"plane/app/service"
)

// Plane Users API
type UserController struct {
	BaseController
}

func (u *UserController) URLMapping() {
	u.Mapping("CreateUser", u.CreateUser)
	u.Mapping("GetAll", u.GetAll)
	u.Mapping("UpdateUser", u.UpdateUser)
	u.Mapping("GetUserByID", u.GetUserByID)
	u.Mapping("delete", u.DeleteUser)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /create [post]
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
// @router /getAll [get]
func (u *UserController) GetAll() {
	users, err := service.UserService.GetAllUser()
	u.SetResult(users, err)
}

// @Title GetUser By user Id
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) GetUserByID() {
	uid, _ := u.GetInt64(":uid")
	user, err := service.UserService.GetUser(uid)
	u.SetResult(user, err)
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /update [post]
func (u *UserController) UpdateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uu, err := service.UserService.UpdateUser(&user, "Id")
	u.SetResult(uu, err)
}

// @Title DeleteUser
// @Description Delete the user
// @Param	uid		path 	string	true		"The uid you want to c"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /delete [put]
func (u *UserController) DeleteUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, user)
	user.Status = 1
	uu, err := service.UserService.UpdateUser(&user, "status")
	u.SetResult(uu, err)
}
