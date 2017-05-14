package controllers

import (
	"encoding/json"
	"plane/app/libs"
	"plane/app/models"
	"plane/app/service"
)

// Operations about logs
type MessageController struct {
	BaseController
}

// Operations about message
func (m *MessageController) URLMapping() {
	m.Mapping("CreateMessage", m.CreateMessage)
	m.Mapping("UserMessage", m.UserUNReadMessage)
	m.Mapping("ReadMessage", m.ReadMessage)
	m.Mapping("GetMessageByID", m.GetMessageByID)
	m.Mapping("GetMessageByPlaneID", m.GetMessageByPlaneID)
}

// getName 设置获取模型名字
func (m *MessageController) getModelName() string {
	return "message"
}

// @Title CreateMessage
// @Description Add the Message
// @Param	json	body 	models.Message	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /add [post]

func (m *MessageController) CreateMessage() {
	var message models.Message
	json.Unmarshal(m.Ctx.Input.RequestBody, &message)
	result, err := service.MessageService.AddMessage(&message)
	m.SetResult(m.getModelName(), result, err)
}

// @Title ReadMessage
// @Description read the Message
// @Param	messageId	 path int true
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /read/:messageId [get]
func (m *MessageController) ReadMessage() {
	messageID, _ := m.GetInt64("messageId", 0)
	err := service.MessageService.ReadMessage(messageID)
	m.SetResultWithInfo("", "read the message Success", nil, err)
}

// @Title UserUNReadMessage
// @Description Add the Message
// @Param	Id		body int true
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /getUserMessage [post]
func (m *MessageController) UserUNReadMessage() {
	userID, err := m.GetInt64("userId", 0)
	if m.CheckErr(err) {
		messages, err := service.MessageService.GetUserUNReadMessage(userID)
		m.SetResult(m.getModelName(), messages, err)
	}
}

// @Title GetMessageById
// @Description get the Message by messageId
// @Param	Id	path int true
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /:messageId [get]
func (m *MessageController) GetMessageByID() {
	messageID, mErr := m.GetInt64(":messageId", 0)
	libs.Log(messageID)
	if m.CheckErr(mErr) {
		message, err := service.MessageService.FindOneByID(messageID)
		libs.Log(message)
		m.SetResult(m.getModelName(), message, err)
	}
}

// @Title GetMessageByPlaneId
// @Description get the Message by messageId
// @Param	Id path int true  "for message planeId"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /plane/:planeId [get]
func (m *MessageController) GetMessageByPlaneID() {
	planeID, err := m.GetInt64(":planeId", 0)
	if m.CheckErr(err) {
		messages, err := service.MessageService.FindMessageByPlaneID(planeID)
		m.SetResult(m.getModelName(), messages, err)
	}
}
