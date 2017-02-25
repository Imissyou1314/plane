package service

import (
	"plane/app/models"
)

type messageService struct{}

func (this *messageService) table() string {
	return tableName("message")
}

func (this *messageService) GetUserUNReadMessage(userId int) (*models.Message, error) {
	messages := &models.Message{}
	messages.ToUserId = userId
	err := o.Read(messages)
	return messages, err
}

func (this *messageService) ReadMessage(messageId int) error {
	message := &models.Message{}
	message.IsRead = true
	message.Id = messageId
	_, err := o.Update(message, "IsRead")
	return err
}

// 根据消息ID获取ID
func (this *messageService) FindOneById(messageId int) (*models.Message, error) {
	message := &models.Message{}
	message.Id = messageId
	err := o.Read(message);
	return message, err
}
