package router

import (
	"github.com/gin-gonic/gin"
	"os"
	"videowebsite/controller"
	"videowebsite/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())  //跨域
	r.Use(middleware.CurrentUser())  //验证登陆

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", controller.Ping)
		// 用户注册
		v1.POST("user/register", controller.UserRegister)
		// 用户登录
		v1.POST("user/login", controller.UserLogin)

		//需要登录保护
		auth := v1.Group("/")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", controller.UserMe)
			auth.DELETE("user/logout", controller.UserLogout)

			auth.PUT("video/:id", controller.UpdateVideo)
			auth.DELETE("video/:id", controller.DeleteVideo)
			auth.POST("video", controller.CreateVideo)
		}
		v1.GET("video/:id", controller.VideoDetail)
		v1.GET("videos", controller.ListVideos)
		//排行榜
		v1.GET("rank/daily", controller.DailyRank)

		v1.POST("upload/token", controller.UploadToken)
	}
	// 需要登录保护的

	return r
}
