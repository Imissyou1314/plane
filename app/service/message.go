package service

import (
	"plane/app/models"
)

type messageService struct{}

func (this *messageService) table() string {
	return tableName("message")
}

func (this *messageService) GetUserUNReadMessage(userId int) ([]*models.Message, error) {
	var messages []*models.Message
	_, err := o.QueryTable("t_message").Filter("user_id", userId).All(&messages)
	return messages, err
}

func (this *messageService) ReadMessage(messageId int) (err error) {
	message := &models.Message{}
	message.IsRead = true
	message.Id = messageId
	_, err = o.Update(message, "IsRead")
	return
}

// 根据消息ID获取ID
func (this *messageService) FindOneById(messageId int) (message *models.Message, err error) {
	message.Id = messageId
	err = o.Read(message)
	return
}

func (this *messageService) AddMessage(message *models.Message) (*models.Message, error) {
	messageID, err := o.Insert(message)
	message, _ = this.FindOneById(int(messageID))
	return message, err
}

func (this *messageService) FindMessageByPlaneId(planeId int) (messages []*models.Message, err error) {
	_, err = o.QueryTable(this.table()).Filter("plane_id", planeId).All(messages)
	return
}
