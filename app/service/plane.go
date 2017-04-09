package service

import (
	"plane/app/models"
)

type planeService struct {
}

func (this *planeService) table() string {
	return tableName("plane")
}

func (this *planeService) FindOneByID(planeId int) (*models.Plane, error) {
	plane := &models.Plane{}
	plane.Id = planeId
	err := o.Read(plane)
	return plane, err
}

func (this *planeService) FindUserPlane(userId int) (plane *models.Plane, err error) {
	plane.UserId = userId
	err = o.Read(plane, "UserId")
	return
}
