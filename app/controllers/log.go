package controllers

import (
	"plane/app/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about logs
type LogController struct {
	beego.Controller
}

// @Title GetLogs
// @Description get the logs by logsId
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.log	true		"body for Log content"
// @Success 200 {object} models.Log
// @Failure 403 :id is not int
// @router /:id [put]
func (this *LogController) GetLogs() {
	logId, err := this.GetInt("id", 0)
	logs.Info(logId)
	Log, err := service.LogService.FindById(logId)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = Log
	}
	this.ServeJSON()
}
