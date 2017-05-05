package service

import (
	"errors"
	"plane/app/models"

	log "github.com/Sirupsen/logrus"
)

type logService struct {
}

func (l *logService) table() string {
	return tableName("log")
}

func (l *logService) FindByID(logID int64) (*models.Log, error) {
	if logID != 0 {
		log := &models.Log{}
		log.Id = logID
		err := o.Read(log)
		return log, err
	}
	return nil, errors.New("传入ID不能为空")
}

func (l *logService) FindByUserID(userID int64) (*models.Log, error) {
	if userID != 0 {
		log := &models.Log{}
		log.UserId = userID
		err := o.Read(log)
		return log, err
	}
	return nil, errors.New("传入用户ID不能为空")
}

func (l *logService) FindByUserName(userName string) (*models.Log, error) {
	if len(userName) != 0 {
		log := &models.Log{}
		log.UserName = userName
		err := o.Read(log)
		return log, err
	}
	return nil, errors.New("传入的用户名为空")
}

func (l *logService) GetAllLogs() ([]models.Log, error) {
	var logs []models.Log
	num, err := o.Raw("select * from t_log").QueryRows(&logs)
	if err != nil {
		log.Error(err)
	} else {
		log.Info(num)
	}
	return logs, err
}

// func (this *logService) getLog(*models.Log log) (*models.Log, error) {
// 		err = o.Read(log)
// 		return log, err
// }
