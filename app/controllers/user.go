package controllers

import (
	"encoding/json"
	"plane/app/models"
	"plane/app/service"
)

// UserController Plane Users API
type UserController struct {
	BaseController
}

// getModelName 获取模型名字
func (u *UserController) getModelName() string {
	return "User"
}

// Operations about User
func (u *UserController) URLMapping() {
	u.Mapping("CreateUser", u.CreateUser)
	u.Mapping("GetAll", u.GetAll)
	u.Mapping("UpdateUser", u.UpdateUser)
	u.Mapping("GetUserByID", u.GetUserByID)
	u.Mapping("DeleteUser", u.DeleteUser)
	u.Mapping("UpdateUserImage", u.UpdateUserImage)
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
	result, err := service.UserService.AddUser(&user)
	u.SetResult(u.getModelName(), result, err)
}

// @Title GetAll user
// @Description get all Users
// @Success 200 {object} []models.User
// @Failure 500 系统错误
// @router /getAll [get]
func (u *UserController) GetAll() {
	users, err := service.UserService.GetAllUser()
	u.SetResult(u.getModelName(), users, err)
}

// @Title GetUser By user Id
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for model.UserId"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) GetUserByID() {
	uid, _ := u.GetInt64(":uid")
	user, err := service.UserService.GetUser(uid)
	u.SetResult(u.getModelName(), user, err)
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
	u.SetResult(u.getModelName(), uu, err)
}

// @Title DeleteUser
// @Description Delete the user
// @Param	uid		path 	string	true		"The uid you want to c"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /delete [post]
func (u *UserController) DeleteUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Status = 1
	uu, err := service.UserService.UpdateUser(&user, "status")
	u.SetResult(u.getModelName(), uu, err)
}

// @Title UpdateUserImage 更新用户头像
// @Description Delete the user
// @Param	uid		path 	string	true		"The uid you want to c"
// @Param	body		body 	file	true		"body for Image file content"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /updateImage [post]
func (u *UserController) UpdateUserImage() {
	userID, _ := u.GetInt64("uid", 0)
	file, h, err := u.GetFile("userImage")
	defer file.Close()
	if u.CheckErr(err) {
		fileName, fileErr := u.SaveFileToLoaction("userImage", h.Filename, ".png")
		if u.CheckErr(fileErr) {
			user, upErr := service.UserService.UpdateUserHeadImage(userID, fileName)
			u.SetResult(u.getModelName(), user, upErr)
		}
	}
}
