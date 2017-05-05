package service

type planeLogService struct {
}

func (l *planeLogService) table() string {
	return tableName("planeLog")
}
