package service

import (
	"plane/app/models"
)

type messageService struct{}

func (this *messageService) table() string {
	return tableName("message")
}

func (this *messageService) GetUserUNReadMessage(userId int64) ([]*models.Message, error) {
	var messages []*models.Message
	_, err := o.QueryTable("t_message").Filter("to_user_id", userId).All(&messages)
	return messages, err
}

func (this *messageService) ReadMessage(messageId int64) error {
	message := &models.Message{}
	message.IsRead = true
	message.Id = messageId
	_, err := o.Update(message, "IsRead")
	return err
}

// 根据消息ID获取ID
func (this *messageService) FindOneById(messageId int64) (*models.Message, error) {
	message := &models.Message{}
	message.Id = messageId
	err := o.Read(message)
	return message, err
}

func (this *messageService) AddMessage(message *models.Message) (*models.Message, error) {
	messageID, err := o.Insert(message)
	message, _ = this.FindOneById(messageID)
	return message, err
}

func (this *messageService) FindMessageByPlaneId(planeId int64) ([]*models.Message, error) {
	var messages []*models.Message
	_, err := o.QueryTable(this.table()).Filter("plane_id", planeId).All(&messages)
	return messages, err
}
