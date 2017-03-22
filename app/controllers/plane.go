package controllers

import (
  "plane/app/service"
  "plane/app/models"

  "github.com/astaxie/beego"
)

type PlaneController struct {
  beego.Controller
}

func (p *PlaneController) getPlaneById() {
  PlaneID := p.GetString(":Id")
  plane, _ := service.PlaneService.getPlaneById(PlaneID)
  p.Data["json"] = plane
  p.ServeJSON()
}

func (p *PlaneController) getPlaneByUserId()  {
    UserID := p.GetString(":userId")
    plane, _ := service.PlaneService.getPlaneByUserId(UserID)
    p.Data["json"] = plane
    p.ServeJSON()
}
