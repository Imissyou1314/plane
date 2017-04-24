package controllers

import (
	"plane/app/models"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["StartTime"] = time.Now()
	log.Info(this.Ctx.Input)
}

func (this *BaseController) CheckErr(err error) bool {
	if err != nil {
		this.Data["json"] = models.SetWithErr(err)
		this.ServeJSON()
		return false
	} else {
		return true
	}
}

// Set result Data
func (this *BaseController) SetResult(data interface{}, err error) {
	result := models.SetResultModel(data, err)
	this.Data["json"] = result
	this.ServeJSON()
}
