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
	m.Mapping("GetMessageByID", m.GetMessageByID)
	m.Mapping("GetMessageByPlaneID", m.GetMessageByPlaneID)
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
// @router /read/:messageId [get]
func (m *MessageController) ReadMessage() {
	messageID, _ := m.GetInt64("messageId", 0)
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
	userID, err := m.GetInt64("userId", 0)
	if m.CheckErr(err) {
		messages, err := service.MessageService.GetUserUNReadMessage(userID)
		m.SetResult(messages, err)
	}
}

// @Title GetMessageById
// @Description get the Message by messageId
// @Param	Id		body
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /:messageId [get]
func (m *MessageController) GetMessageByID() {
	messageID, mErr := m.GetInt64(":messageId", 0)
	libs.Log(messageID)
	if m.CheckErr(mErr) {
		message, err := service.MessageService.FindOneByID(messageID)
		libs.Log(message)
		m.SetResult(message, err)
	}
}

// @Title GetMessageByPlaneId
// @Description get the Message by messageId
// @Param	Id		body
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /plane/:planeId [get]
func (m *MessageController) GetMessageByPlaneID() {
	planeID, err := m.GetInt64(":planeId", 0)
	if m.CheckErr(err) {
		messages, err := service.MessageService.FindMessageByPlaneID(planeID)
		m.SetResult(messages, err)
	}
}
