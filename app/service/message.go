package service

import (
	"plane/app/models"
)

type messageService struct{}

func (m *messageService) table() string {
	return tableName("message")
}

func (m *messageService) GetUserUNReadMessage(userID int64) ([]*models.Message, error) {
	var messages []*models.Message
	_, err := o.QueryTable("t_message").Filter("to_user_id", userID).All(&messages)
	return messages, err
}

func (m *messageService) ReadMessage(messageID int64) error {
	message := &models.Message{}
	message.IsRead = true
	message.Id = messageID
	_, err := o.Update(message, "IsRead")
	return err
}

// 根据消息ID获取ID
func (m *messageService) FindOneByID(messageID int64) (*models.Message, error) {
	message := &models.Message{}
	message.Id = messageID
	err := o.Read(message)
	return message, err
}

func (m *messageService) AddMessage(message *models.Message) (*models.Message, error) {
	messageID, err := o.Insert(message)
	message, _ = m.FindOneByID(messageID)
	return message, err
}

func (m *messageService) FindMessageByPlaneID(planeID int64) ([]*models.Message, error) {
	var messages []*models.Message
	_, err := o.QueryTable(m.table()).Filter("plane_id", planeID).All(&messages)
	return messages, err
}
