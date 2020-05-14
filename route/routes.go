package route

import (
"github.com/gin-gonic/gin"
"huana/controller"
// "net/http"
"huana/middleware"

)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
		r.POST("/api/auth/upload", controller.Upload)
		r.POST("/api/auth/multipleupload", controller.MultipleUpload)

	

	
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