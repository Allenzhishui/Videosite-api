package controller

import (
	"github.com/gin-gonic/gin"
	"videowebsite/service"
)

func DailyRank(c *gin.Context) {
	service := service.DailyRankService{}
	if err := c.ShouldBind(&service);err == nil {
		resp := service.Get()
		c.JSON(200,resp)
	}else {
		c.JSON(200,ErrorResponse(err))
	}
}
