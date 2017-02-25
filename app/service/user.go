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

func (this *userService) table() string {
	return tableName("user")
}

func (this *userService) GetAllUser() ([]models.User, error) {
	var users []models.User
	o.Read(users)
	if len(users) != 0 {
		return users, nil
	} else {
		err := newError("查询没有结果")
		return nil, err
	}
}

func (this *userService) GetUser(userId int) (*models.User, error) {
	user := &models.User{}
	user.Id = userId
	err := o.Read(user)
	return user, err
}

// 根据用户名来获取用户
func (this *userService) GetUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	user.UserName = userName
	err := o.Read(user, "UserName")
	return user, err
}

// 统计用户总数
func (this *userService) GetTotal() (int64, error) {
	return o.QueryTable(this.table()).Count()
}

// 添加新用户
func (this *userService) AddUser(userName, email, password string, sex int) (*models.User, error) {
	if exists, _ := this.GetUserByName(userName); exists.Id > 0 {
		return nil, errors.New("用户名已经存在")
	}

	user := &models.User{}
	user.UserName = userName
	user.Sex = sex
	user.Email = email
	user.Salt = string(utils.RandomCreateBytes(10))
	user.Password = libs.Md5([]byte(password + user.Salt))
	_, err := o.Insert(user)
	return user, err
}

func (this *userService) UpdateUser(user *models.User, fileds ...string) (*models.User, error) {
	fmt.Print(fileds)
	userId, err := o.Update(user)
	if err != nil {
		return nil, errors.New("更新失败")
	} else {
		return this.GetUser(int(userId))
	}
}

// 修改登录密码
func (this *userService) ModifyPassword(userId int, password string) error {
	user, err := this.GetUser(userId)
	if err != nil {
		return err
	}

	user.Salt = string(utils.RandomCreateBytes(10))
	user.Password = libs.Md5([]byte(password + user.Salt))
	_, err = o.Update(user, "Salt", "Password")
	return err
}

// 根据ID删除用户
func (this *userService) DeteleUser(userId int) error {
	if userId == 1 {
		return errors.New("不允许删除用户ID为1的用户")
	}
	user := &models.User{
		Id: userId,
	}
	_, err := o.Delete(user)
	return err
}

// 根据用户名删除用户
func (this *userService) DeteleUserByUserName(userName string) error {
	if userName == "" {
		return errors.New("用户名不能为空")
	}
	user := &models.User{
		UserName: userName,
	}
	_, err := o.Delete(user)
	return err
}

// 登录操作
func (this *userService) Login(userName, password string) (*models.User, error) {
	if len(userName) == 0 || len(password) == 0 {
		return nil, errors.New("用户名与密码不能为空")
	}
	user, err := this.GetUserByName(userName)
	if user == nil {
		return nil, errors.New("账号错误")
	}
	if user.Password != libs.Md5([]byte(password+user.Salt)) {
		return nil, errors.New("密码错误")
	}
	return user, err
}

func (this *userService) Logout(userId int) (*models.User, error) {
	user := &models.User{
		Id:     userId,
		Status: 0,
	}
	return this.UpdateUser(user)
}
