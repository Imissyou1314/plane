package service

import (
	"plane/app/models"
)

type planeService struct {
}

func (p *planeService) table() string {
	return tableName("plane")
}

func (p *planeService) FindOneByID(planeID int64) (*models.Plane, error) {
	plane := &models.Plane{}
	plane.Id = planeID
	err := o.Read(plane)
	return plane, err
}

func (p *planeService) FindUserPlane(userID int64) (*models.Plane, error) {
	var plane = &models.Plane{}
	plane.UserId = userID
	err := o.Read(plane, "UserId")
	return plane, err
}

func (p *planeService) GetAll() ([]*models.Plane, error) {
	var planes []*models.Plane
	_, err := o.QueryTable("t_plane").Filter("is_back", false).All(&planes)
	return planes, err
}

func (p *planeService) AddPlane(plane *models.Plane) (result *models.Plane, err error) {
	id, err := o.Insert(plane)
	if err != nil {
		result, err = p.FindOneByID(id)
	}
	return
}
