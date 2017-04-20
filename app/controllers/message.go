package controllers

import (
	"encoding/json"
	"plane/app/models"
	"plane/app/service"

	"github.com/astaxie/beego"
)

type MessageController struct {
	beego.Controller
}

func (m *MessageController) CreateMessage() {
	var message models.Message
	json.Unmarshal(m.Ctx.Input.RequestBody, &message)
	result, _ := service.MessageService.AddMessage(&message)
	m.Data["json"] = result
	m.ServeJSON()
}
