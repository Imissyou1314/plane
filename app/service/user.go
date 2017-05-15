package service

import (
	"errors"
	"fmt"
	"plane/app/libs"
	"plane/app/models"

	"github.com/astaxie/beego/utils"
)

type userService struct {
}

func (u *userService) table() string {
	return tableName("user")
}

func (u *userService) GetAllUser() ([]*models.User, error) {
	var users []*models.User
	_, err := o.QueryTable("t_user").All(&users)
	return users, err
}

func (u *userService) GetUser(userID int64) (*models.User, error) {
	user := &models.User{}
	user.Id = userID
	err := o.Read(user)
	return user, err
}

// 根据用户名来获取用户
func (u *userService) GetUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	user.UserName = userName
	err := o.Read(user, "UserName")
	return user, err
}

// 统计用户总数
func (u *userService) GetTotal() (int64, error) {
	return o.QueryTable(u.table()).Count()
}

// 添加新用户
func (u *userService) AddUser(user *models.User) (*models.User, error) {
	if exists, _ := u.GetUserByName(user.UserName); exists.Id > 0 {
		return nil, errors.New("用户名已经存在")
	}

	user.Salt = string(utils.RandomCreateBytes(10))
	user.Password = libs.Md5([]byte(user.Password + user.Salt))
	_, err := o.Insert(user)
	return user, err
}

func (u *userService) UpdateUser(user *models.User, fileds ...string) (*models.User, error) {
	fmt.Print(fileds)
	userID, err := o.Update(user)
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return u.GetUser(userID)
}

// 修改登录密码
func (u *userService) ModifyPassword(userID int64, password string) error {
	user, err := u.GetUser(userID)
	if err != nil {
		return err
	}

	user.Salt = string(utils.RandomCreateBytes(10))
	user.Password = libs.Md5([]byte(password + user.Salt))
	_, err = o.Update(user, "Salt", "Password")
	return err
}

// 根据ID删除用户
func (u *userService) DeteleUser(userID int64) error {
	if userID == 1 {
		return errors.New("不允许删除用户ID为1的用户")
	}
	user := &models.User{
		Id: userID,
	}
	_, err := o.Delete(user)
	return err
}

// 根据用户名删除用户
func (u *userService) DeteleUserByUserName(userName string) error {
	if userName == "" {
		return errors.New("用户名不能为空")
	}
	user := &models.User{
		UserName: userName,
	}
	_, err := o.Delete(user)
	return err
}

// Login 登录操作
func (u *userService) Login(userName, password string) (*models.User, error) {
	if len(userName) == 0 || len(password) == 0 {
		return nil, errors.New("用户名与密码不能为空")
	}
	user, err := u.GetUserByName(userName)
	if user == nil {
		return nil, errors.New("账号错误")
	}
	if user.Password != libs.Md5([]byte(password+user.Salt)) {
		return nil, errors.New("密码错误")
	}
	return user, err
}

// Logout 推出登陆
func (u *userService) Logout(userID int64) (*models.User, error) {
	user := &models.User{
		Id:     userID,
		Status: 0,
	}
	return u.UpdateUser(user)
}

func (u *userService) UpdateUserHeadImage(userId int64, imageUrl string) (*models.User, error) {
	user, err := u.GetUser(userId)
	if err != nil {
		return nil, err
	}
	user.ImageUrl = imageUrl
	newUser, err := u.UpdateUser(user, "image_url")
	return newUser, err
}
