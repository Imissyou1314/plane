package service

import (
	"plane/app/models"
)

type planeService struct {
}

func (this *planeService) table() string {
	return tableName("plane")
}

func (this *planeService) FindOneByID(planeId int64) (*models.Plane, error) {
	plane := &models.Plane{}
	plane.Id = planeId
	err := o.Read(plane)
	return plane, err
}

func (this *planeService) FindUserPlane(userId int64) (*models.Plane, error) {
	var plane = &models.Plane{}
	plane.UserId = userId
	err := o.Read(plane, "UserId")
	return plane, err
}

func (this *planeService) GetAll() ([]*models.Plane, error) {
	var planes []*models.Plane
	_, err := o.QueryTable("t_plane").Filter("is_back", false).All(&planes)
	return planes, err
}

func (this *planeService) AddPlane(plane *models.Plane) (result *models.Plane, err error) {
	id, err := o.Insert(plane)
	if err != nil {
		result, err = this.FindOneByID(id)
		return
	} else {
		result = nil
		return nil, err
	}
}
