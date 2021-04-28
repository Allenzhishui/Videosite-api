package service

import (
	"videowebsite/model"
	"videowebsite/serializer"
)

type CreateVideoService struct {
	Title  string ` form:"title" json:"title" binding:"required,min=2,max=30"`
	Info   string ` form:"info" json:"info" binding:"max=300"`
	URL    string ` form:"url" json:"url" `
	Avatar string ` form:"avatar" json:"avatar"`
}

func (service *CreateVideoService) Create(user *model.User) serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
		URL:   service.URL,
		Avatar : service.Avatar,
		UserID : user.ID,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
