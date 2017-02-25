package service

type planeService struct {
}

func (this *planeService) table() string {
	return tableName("plane")
}
