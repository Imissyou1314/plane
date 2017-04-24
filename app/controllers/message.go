package controllers

import (
	"encoding/json"
	"plane/app/libs"
	"plane/app/models"
	"plane/app/service"
)

type MessageController struct {
	BaseController
}

func (m *MessageController) URLMapping() {
	m.Mapping("CreateMessage", m.CreateMessage)
	m.Mapping("UserMessage", m.UserUNReadMessage)
	m.Mapping("ReadMessage", m.ReadMessage)
	m.Mapping("GetMessageById", m.GetMessageById)
	m.Mapping("GetMessageByPlaneId", m.GetMessageByPlaneId)
}

// @Title AddMessage
// @Description Add the Message
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /add [post]
func (m *MessageController) CreateMessage() {
	var message models.Message
	json.Unmarshal(m.Ctx.Input.RequestBody, &message)
	result, _ := service.MessageService.AddMessage(&message)
	m.Data["json"] = result
	m.ServeJSON()
}

// @Title ReadMessage
// @Description Add the Message
// @Param	Id		body
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /read [post]
func (m *MessageController) ReadMessage() {
	messageID, _ := m.GetInt("messageId", 0)
	err := service.MessageService.ReadMessage(messageID)
	m.SetResult("read the message Success", err)
}

// @Title UserUNReadMessage
// @Description Add the Message
// @Param	Id		body
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /getUserMessage [post]
func (m *MessageController) UserUNReadMessage() {
	userId, err := m.GetInt("userId", 0)
	if m.CheckErr(err) {
		messages, err := service.MessageService.GetUserUNReadMessage(userId)
		m.SetResult(messages, err)
	}
}

// @Title GetMessageById
// @Description get the Message by messageId
// @Param	Id		body
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router / [get]
func (m *MessageController) GetMessageById() {
	messageId, mErr := m.GetInt("messageId", 0)
	libs.Log(messageId)
	if m.CheckErr(mErr) {
		message, err := service.MessageService.FindOneById(messageId)
		libs.Log(message)
		m.SetResult(message, err)
	}
}

// @Title GetMessageByPlaneId
// @Description get the Message by messageId
// @Param	Id		body
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /plane [get]
func (m *MessageController) GetMessageByPlaneId() {
	planeId, err := m.GetInt("planeId", 0)
	if m.CheckErr(err) {
		messages, err := service.MessageService.FindMessageByPlaneId(planeId)
		m.SetResult(messages, err)
	}
}
