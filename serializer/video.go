package serializer

import (
	"videowebsite/model"
)

type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	Avatar    string `json:"avatar"`
	URL       string `json:"url"`
	View      uint64 `json:"view"`
	User      User   `json:"user"`
	CreatedAt int64  `json:"created_at"`
}

func BuildVideo(self model.Video) Video {
	user, _ := model.GetUser(self.UserID)
	return Video{
		ID:        self.ID,
		Title:     self.Title,
		Info:      self.Info,
		URL:       self.URL,
		Avatar:    self.AvatarURL(),
		View:      self.View(),
		User:      BuildUser(user),
		CreatedAt: self.CreatedAt.Unix(),
	}
}

func BuildVideos(self []model.Video) (videos []Video) {
	for _, i := range self {
		video := BuildVideo(i)
		videos = append(videos, video)
	}
	return
}
