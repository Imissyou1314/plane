package controllers

import (
	"plane/app/models"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)
}

func (base *BaseController) Prepare() {
	base.Data["StartTime"] = time.Now()
	log.Info(base.Ctx.Input)
}

func (base *BaseController) CheckErr(err error) bool {
	if err != nil {
		base.Data["json"] = models.SetWithErr(err)
		base.ServeJSON()
		return false
	}
	return true
}

// Set result Data
func (base *BaseController) SetResult(data interface{}, err error) {
	result := models.SetResultModel(data, err)
	base.Data["json"] = result
	base.ServeJSON()
}
