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

func (this *planeService) GetAll() (planes []models.Plane, err error) {
	err = o.Read(planes)
	return
}

func (this *planeService) AddPlane(plane *models.Plane) (result *models.Plane, err error) {
	id, err := o.Insert(plane)
	if err != nil {
		result, err = this.FindOneByID(int(id))
		return
	} else {
		result = nil
		return nil, err
	}
}
