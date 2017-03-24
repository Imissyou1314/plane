package controllers

import (
  "plane/app/service"
  "strconv"
  "github.com/astaxie/beego"
)

type PlaneController struct {
  beego.Controller
}

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

func (p *PlaneController) GetPlaneByUserId()  {
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
