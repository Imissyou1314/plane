package service

type planeLogService struct {
}

func (this *planeLogService) table() string {
	return tableName("planeLog")
}
