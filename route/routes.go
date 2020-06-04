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
	r.POST("/api/articlepost", middleware.AuthUserPost(), controller.NewArticlePost)
	r.POST("/api/follower", controller.GetFollower)
	r.POST("/api/follower/add", controller.AddFollower)
	r.DELETE("/api/follower/delete/:userid/:followerid", controller.DeleteFollower)
	r.POST("/api/following", controller.GetFollowing)
	r.POST("/api/user/info", controller.UserInfo)
	r.POST("/api/user/all", controller.AllUser)
	r.POST("/api/unfollow/user", controller.UnfollowUser)
	r.POST("/api/user", controller.GetAllUser)
	r.GET("/api/category", controller.GetCategory)
	r.POST("/api/videopost", middleware.AuthUserPost(), controller.PostVideo)
	r.POST("/api/video/all", controller.GetAllVideo)
	r.POST("/api/all/momentpost", controller.AllMomentPost)
	r.PUT("/api/update/moment/like", controller.UpdateMomentLike)
	r.POST("/api/userpost/count", controller.UserPost)
	r.POST("/api/all/articlepost", controller.AllArticlePost)
	r.PUT("/api/article/like/update", controller.UpdateArticleLike)
	r.POST("/api/like/record", controller.RecordLike)
	r.DELETE("/api/like/record/delete/:userid/:postid/:field", controller.DeleteLike)
	r.POST("/api/like/record/check", controller.CheckLike)
	r.POST("/api/comment/record", controller.RecordComment)
	r.DELETE("/api/comment/record/delete/:commentid", controller.DeleteComment)
	r.POST("/api/comment/all", controller.GetAllComment)
	r.PUT("/api/update/comment/like", controller.UpdateCommentLike)
	r.DELETE("/api/comment/like/record/delete/:postid/:field", controller.DeleteCommentLikeRecord)
	r.POST("/api/comment/count", controller.GetCommentCount)
	r.POST("/api/reply/record", controller.RecordReply)
	r.POST("/api/reply/all", controller.GetAllReply)
	r.POST("/api/reply/count", controller.GetReplyCount)
	r.DELETE("/api/reply/delete/:replyid", controller.DeleteReply)
	r.DELETE("/api/reply/like/delete/:postid/:field", controller.DeleteReplyLikeRecord)
	r.PUT("/api/reply/like/update", controller.UpdateReplyLike)
	// 以下所有路由是自己新加上去的哟～
	r.GET("/api/auth/user/info", middleware.AuthMiddleware(), controller.UserProfileInfo)
	r.POST("/api/auth/feedback", middleware.AuthUserFeedback(), controller.Feedback)
	r.POST("/api/auth/user/update/name", middleware.AuthMiddleware(), controller.UpdateUserName)
	r.POST("/api/auth/user/update/profileimage", middleware.AuthMiddleware(), controller.UpdateUserProfileImage)
	r.POST("/api/auth/user/update/introduction", middleware.AuthMiddleware(), controller.UpdateUserIntroduction)
	r.POST("/api/auth/user/update/gender", middleware.AuthMiddleware(), controller.UpdateUserGender)
	r.POST("/api/auth/user/update/birthday", middleware.AuthMiddleware(), controller.UpdateUserBirthday)

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
