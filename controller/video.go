package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"videowebsite/service"
)

//CreateVideo视频投稿
func CreateVideo(c *gin.Context) {
	user := CurrentUser(c)
	mod := service.CreateVideoService{}
	if err := c.ShouldBind(&mod); err == nil {
		resp := mod.Create(user)
		c.JSON(200, resp)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//VideoDetail  视频详情页
func VideoDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(200, err)
	}
	mod := service.VideoDetailService{}
	resp := mod.Detail(id)
	c.JSON(200, resp)
}

//ListVideos 视频列表页
func ListVideos(c *gin.Context) {
	mods := service.ListVideoService{}
	if err := c.ShouldBind(&mods); err == nil {
		resp := mods.List()
		c.JSON(200, resp)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//UpdateVideo 更新视频
func UpdateVideo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(200, err)
	}
	mod := service.UpdateVideoService{}
	if err = c.ShouldBind(&mod); err == nil {
		resp := mod.Update(id)
		c.JSON(200, resp)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//DeleteVideo  删除视频
func DeleteVideo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(200, err)
	}
	mod := service.DeleteVideoService{}
	resp := mod.Delete(id)
	c.JSON(200, resp)
}



