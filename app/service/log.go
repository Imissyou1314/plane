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
