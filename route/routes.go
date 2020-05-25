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
	r.GET("/api/category", controller.GetCategory)
	r.POST("/upload", controller.UploadFile)
	// 以下所有路由是自己新加上去的哟～
	r.GET("/api/auth/user/info", middleware.AuthMiddleware(), controller.UserInfo)
	r.POST("/api/auth/feedback", middleware.AuthUserFeedback(), controller.Feedback)
	r.POST("/api/auth/user/update/name", middleware.AuthMiddleware(), controller.UpdateUserName)
	r.POST("/api/auth/user/update/profileimage", middleware.AuthMiddleware(), controller.UpdateUserProfileImage)
	r.POST("/api/auth/user/update/introduction", middleware.AuthMiddleware(), controller.UpdateUserIntroduction)
	r.POST("/api/auth/user/update/gender", middleware.AuthMiddleware(), controller.UpdateUserGender)
	r.POST("/api/auth/user/update/birthday", middleware.AuthMiddleware(), controller.UpdateUserBirthday)

	return r
}
