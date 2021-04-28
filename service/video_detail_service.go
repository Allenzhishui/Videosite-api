package service

import (
	"videowebsite/model"
	"videowebsite/serializer"
)

type VideoDetailService struct {

}

func (service *VideoDetailService) Detail(id int64) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	//处理视频的一系列问题
	video.AddView()
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}



