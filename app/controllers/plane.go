package controllers

import (
  "plane/app/service"
  "strconv"
  "github.com/astaxie/beego"
)

// Operations about Plane
// @Title getStaticBlock
// @Description get all the staticblock by key
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /staticblock/:key [get]
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
