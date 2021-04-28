package controller

import (
	"github.com/gin-gonic/gin"
	"videowebsite/service"
)

func UploadToken(c *gin.Context) {
	service := service.UploadTokenService{}
	if err := c.ShouldBind(&service);err == nil {
		resp := service.Post()
		c.JSON(200,resp)
	}else {
		c.JSON(200,ErrorResponse(err))
	}
}