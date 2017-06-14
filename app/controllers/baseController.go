package controllers

import (
	"plane/app/libs"
	"plane/app/models"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
)

// BaseController 基路由类
type BaseController struct {
	beego.Controller
}

var log = logrus.New()
var savePath = "static/upload/"
var videoPath = "static/video"

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)
}

// Prepare 前置控制器
func (base *BaseController) Prepare() {
	base.Data["StartTime"] = time.Now()
	log.Info(base.Ctx.Input)
}

// CheckErr 集成检查错误
func (base *BaseController) CheckErr(err error) bool {
	if err != nil {
		base.bindData(models.SetWithErr(err))
		return false
	}
	return true
}

// SetResult Set result Data 设置单一返回数据
func (base *BaseController) SetResult(key string, data interface{}, err error) {
	base.bindData(models.SetResultModel(key, "", data, err))
}

// SetResultWithInfo 设置带info的返回数据
func (base *BaseController) SetResultWithInfo(key, info string, data interface{},
	err error) {
	base.bindData(models.SetResultModel(key, info, data, err))
}

// SetResultWithMuilt 设置多值返回
func (base *BaseController) SetResultWithMuilt(key []string, data []interface{},
	err error) {
	mapData := models.GetMapData()
	for i := range key {
		mapData[key[i]] = data[i]
	}
	base.bindData(models.SetResultWithMapData(mapData, err))
}

// bindData 设置返回数据
func (base *BaseController) bindData(resultData *models.Result) {
	base.Data["json"] = resultData
	base.ServeJSON()
}

// SaveFileToLoaction 保存文件到本地
func (base *BaseController) SaveFileToLoaction(formfile, toFile,
	fileTypeName string) (fileNmae string, err error) {
	fileNmae = libs.Md5([]byte(toFile)) + fileTypeName
	err = base.SaveToFile(formfile, savePath+fileNmae)
	return
}
