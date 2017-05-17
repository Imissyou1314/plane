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
}

func (v *VideoController) getModelName() string {
	return "video"
}

// @Title UploadVideo 添加带图片消息
// @Description get the Message by messageId
// @Param	Id path int true  "for message planeId"
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
