package service

import (
	"plane/app/models"
)

type videoService struct{}

func (v *videoService) table() string {
	return tableName("video")
}

// AddVideo 添加视屏
func (v *videoService) AddVideo(video *models.Video) (*models.Video, error) {
	VideoID, err := o.Insert(video)
	video, _ = v.FindOneByID(VideoID)
	return video, err
}

// FindOneByID 根据消息ID获取ID
func (v *videoService) FindOneByID(videoID int64) (*models.Video, error) {
	video := &models.Video{}
	video.Id = videoID
	err := o.Read(video)
	return video, err
}
