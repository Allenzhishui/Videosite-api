package service

import (
	"videowebsite/model"
	"videowebsite/serializer"
)

type UpdateVideoService struct {
	Title string ` form:"title" json:"title"  binding:"required,min=2,max=30"`
	Info  string ` form:"info"  json:"info" binding:"max=300"`
}

func (service * UpdateVideoService) Update(id int64) serializer.Response {
	var video model.Video
	err := model.DB.First(&video,id).Error
	if err != nil {
		return serializer.Response {
			Code :404,
			Msg : "视频不存在",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:502,
			Msg: "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data:serializer.BuildVideo(video),
	}

}