package service

import (
	"plane/app/models"
)

type userService struct {
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

func (this *userService) GetUser(userId int) (models.User, error) {
	var user models.User
	user.Id = userId
	err := o.Read(user)
	return user, err
}

func (this *userService) UpdateUser(userId int, user *models.User) (int64, error) {
	user.Id = userId
	Id, err := o.Insert(user)
	return Id, err
}
