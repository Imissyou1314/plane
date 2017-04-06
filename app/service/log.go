package service

import (
	"errors"
	"plane/app/models"
)

type logService struct {
}

func (this *logService) table() string {
	return tableName("log")
}

func (this *logService) FindById(logId int) (*models.Log, error) {
	if logId != 0 {
		log := &models.Log{}
		log.Id = logId
		err := o.Read(log)
		return log, err
	} else {
		return nil, errors.New("传入ID不能为空")
	}
}

func (this *logService) FindByUserId(userId int) (*models.Log, error) {
	if userId != 0 {
		log := &models.Log{}
		log.UserId = userId
		err := o.Read(log)
		return log, err
	} else {
		return nil, errors.New("传入用户ID不能为空")
	}
}

func (this *logService) FindByUserName(userName string) (*models.Log, error) {
	if len(userName) != 0{
		log := &models.Log{}
		log.UserName = userName
		err := o.Read(log)
		return log, err
	} else {
		return nil, errors.New("传入的用户名为空")
	}
}

// func (this *logService) getLog(*models.Log log) (*models.Log, error) {
// 		err = o.Read(log)
// 		return log, err
// }
