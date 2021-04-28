package controller

import (
	"videowebsite/serializer"
	"videowebsite/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	input :=  service.UserRegisterService{}
	if err := c.ShouldBind(&input); err == nil {
		resp := input.Register()
		c.JSON(200, resp)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	input:=  service.UserLoginService{}
	if err := c.ShouldBind(&input); err == nil {
		resp := input.Login(c)
		c.JSON(200, resp)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	resp := serializer.BuildUserResponse(*user)
	c.JSON(200, resp)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}