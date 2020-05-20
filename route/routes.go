package route

import (
	"huana/controller"

	"github.com/gin-gonic/gin"

	// "net/http"
	"huana/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/auth/upload", controller.Upload)
	r.POST("/api/auth/multipleupload", controller.MultipleUpload)
	r.POST("/api/userpost", middleware.AuthMiddleware(), controller.NewPost)
	r.POST("/api/userpost/info", middleware.AuthUserPost(), controller.UserpostInfo)
	r.POST("/api/momentpost", middleware.AuthUserPost(), controller.NewMomentPost)
	r.POST("/api/follower", controller.GetFollower)
	r.POST("/api/following", controller.GetFollowing)
	r.POST("/api/user/info", controller.UserInfo)
	r.GET("/api/category", controller.GetCategory)
	r.POST("/api/videopost", middleware.AuthUserPost(), controller.PostVideo)
	r.GET("/api/all/momentpost", controller.AllMomentPost)

	return r
}

// func CollectRoute() *gin.Engine {
// 	r := gin.Default()

// 	v1 := r.Group("/api/auth")
// 	{
// v1.Use(middleware.CORSMiddleware())
// v1.POST("register", controller.Register)
// v1.POST("register", controller.Login)
// v1.GET("info", middleware.AuthMiddleware(), controller.Info)

// 	}
// return r
// }
