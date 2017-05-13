package controllers

import (
	"encoding/json"
	"plane/app/models"
	"plane/app/service"
)

// Plane API get extend the BaseController
type PlaneController struct {
	BaseController
}

// Operations about plane
func (p *PlaneController) URLMapping() {
	p.Mapping("GetPlaneByID", p.GetPlaneByID)
	p.Mapping("GetPlaneByUserID", p.GetPlaneByUserID)
	p.Mapping("CreatePlane", p.CreatePlane)
}

// @Title CreatePlane 创建飞机
// @Description 添加飞机
// @param Plane json Plane true
// @Success 200 {object} models.Plane
// @Failure 500 系统发生错误
// @router /create [post]
func (p *PlaneController) CreatePlane() {
	var plane models.Plane
	json.Unmarshal(p.Ctx.Input.RequestBody, &plane)
	result, err := service.PlaneService.AddPlane(&plane)
	p.SetResult(result, err)
}

// @Title 根据飞机Id 获取飞机
// @Description 根据飞机Id 获取指定飞机
// @param Id query int true "PlaneID"
// @Success 200 {object} models.Plane
// @Failure 500 系统发生错误
// @router /:planeId [get]
func (p *PlaneController) GetPlaneByID() {
	planeID, _ := p.GetInt64(":planeId")
	plane, err := service.PlaneService.FindOneByID(planeID)
	p.SetResult(plane, err)
}

// @Title 根据用户ID获取飞机
// @Description 根据用户ID获取用户飞机
// @param userId query int true "userId"
// @Success 200 {object} models.Plane
// @Failure 500 未知错误
// @router /user/:userId [get]
func (p *PlaneController) GetPlaneByUserID() {
	userID, _ := p.GetInt64(":userId")
	plane, err := service.PlaneService.FindUserPlane(userID)
	p.SetResult(plane, err)
}

// @Title 获取所有的飞机
// @Description 获取所有的飞机
// @param userId query int true "userId"
// @Success 200 {object} models.Plane
// @Failure 500 未知错误
// @router /getAll [get]
func (p *PlaneController) GetAllPlane() {
	planes, err := service.PlaneService.GetAll()
	p.SetResult(planes, err)
}
