package controllers

import (
	"encoding/json"
	"plane/app/models"
	"plane/app/service"
)

// VideoController 视屏控制器
type VideoController struct {
	BaseController
}

// URLMapping about Video
func (v *VideoController) URLMapping() {
	v.Mapping("uploadVideo", v.UploadVideo)
	v.Mapping("GetAllVideo", v.GetAllVideo)
	v.Mapping("GetVideoByID", v.GetVideoByID)
}

func (v *VideoController) getModelName() string {
	return "video"
}

// @Title UploadVideo 添加视频
// @Description 上传视频到服务器
// @Param	video body models.Video true  "视频"
// @Success 200 {object} models.User
// @Failure 403 :dont have this user
// @router /addVideo [post]
func (v *VideoController) UploadVideo() {
	var video models.Video
	json.Unmarshal(v.Ctx.Input.RequestBody, &video)
	file, h, err := v.GetFile("video")
	defer file.Close()
	if v.CheckErr(err) {
		fileName, fileErr := v.SaveFileToLoaction("video", h.Filename+"", ".mp4")
		if v.CheckErr(fileErr) {
			video.VideoUrl = fileName
			result, err := service.VideoService.AddVideo(&video)
			v.SetResult(v.getModelName(), result, err)
		}
	}
}

// @Title GetAllVideo 获取所有的视频
// @Description 获取所有的视频
// @Success 200 {object} []models.Video
// @Failure 500 :dont have this videos
// @router /getAll [get]
func (v *VideoController) GetAllVideo() {
	data, err := service.VideoService.GetAllVideo()
	v.SetResult(v.getModelName(), data, err)
}

// @Title GetVideoByID 获取已知Id的视频
// @Description 获取已知Id的视频
// @param videoId path int true 已知视频的Id
// @Success 200 {object} models.Video
// @Failure 500 :dont have this videos
// @router /getById/{:videoId} [get]
func (v *VideoController) GetVideoByID() {
	videoID, err := v.GetInt64("videoId")
	if v.CheckErr(err) {
		data, err := service.VideoService.FindOneByID(videoID)
		v.SetResult(v.getModelName(), data, err)
	}
}
