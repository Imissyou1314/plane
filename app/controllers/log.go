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

func (l *LogController) URLMapping() {
	l.Mapping("GetLog", l.GetLog)
	l.Mapping("GetAllLogs", l.GetAllLogs)
}

// @Title GetLogs
// @Description get log By logId
// @param id string true
// @Success 200 {Array} []models.Log
// @Failure 404 Log not found
// @router /GetLog/:id [get]
func (l *LogController) GetLog() {
	logID, err := l.GetInt64("id", 0)
	logs.Info(logID)
	Log, err := service.LogService.FindByID(logID)
	if err != nil {
		l.Data["json"] = err
	} else {
		l.Data["json"] = Log
	}
	l.ServeJSON()
}

// @Title get All logs
// @Description get All Logs
// @Success 200 {array} models.Logs.LogsList
// @Failure 500 Logs not fild
// @router /getAllLogs [get]
func (l *LogController) GetAllLogs() {
	logs, err := service.LogService.GetAllLogs()
	if err != nil {
		l.Data["json"] = err
	} else {
		l.Data["json"] = logs
	}
	l.ServeJSON()
}
