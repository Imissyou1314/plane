package controllers

import (
	"plane/app/service"
	"strconv"

	"github.com/astaxie/beego"
)

// Plane API
type PlaneController struct {
	beego.Controller
}

func (p *PlaneController) URLMapping() {
	p.Mapping("GetPlaneById", p.GetPlaneById)
	p.Mapping("GetPlaneByUserId", p.GetPlaneByUserId)
}

// @Title 根据飞机Id 获取飞机
// @Description 根据飞机Id 获取指定飞机
// @param Id query int true "PlaneID"
// @Success 200 {object} models.Plane
// @Failure 500 系统发生错误
// @router /plane [get]
func (p *PlaneController) GetPlaneById() {
	Id := p.GetString(":Id")
	PlaneID, err := strconv.Atoi(Id)
	plane, _ := service.PlaneService.FindOneByID(PlaneID)
	if err != nil {
		p.Data["json"] = err.Error()
	} else {
		p.Data["json"] = plane
	}
	p.ServeJSON()
}

// @Title 根据用户ID获取飞机
// @Description 根据用户ID获取用户飞机
// @param userId query int true "userId"
// @Success 200 {object} models.Plane
// @Failure 500 未知错误
// @router /user/plane [get]
func (p *PlaneController) GetPlaneByUserId() {
	userId := p.GetString(":userId")
	UserID, err := strconv.Atoi(userId)
	plane, _ := service.PlaneService.FindUserPlane(UserID)
	if err != nil {
		p.Data["json"] = err.Error()
	} else {
		p.Data["json"] = plane
	}
	p.ServeJSON()
}

func (p *PlaneController) GetAllPlane() {
	planes, err := service.PlaneService.GetAll()
	if err != nil {
		p.Data["json"] = err.Error()
	} else {
		p.Data["json"] = planes
	}
	p.ServeJSON()
}
