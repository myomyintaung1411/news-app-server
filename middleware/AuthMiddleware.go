package middleware

import (
	"huana/common"
	"huana/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//vcalidate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Wrong Token"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 405, "msg": "Wrong Token"})
			ctx.Abort()
			return
		}

		//验证通过后获取Claiim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户不存在
		if user.Userid == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 410, "msg": "User doesn't exist"})
			ctx.Abort()
			return
		}
		//用户存在，将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}

//check auth for userpost
func AuthUserPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//vcalidate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Wrong Token"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseUserPostToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 405, "msg": "Wrong Token"})
			ctx.Abort()
			return
		}

		//验证通过后获取Claiim中的userId
		userpostid := claims.Userpostid
		DB := common.GetDB()
		var userpost model.Userpost
		DB.First(&userpost, userpostid)

		//用户不存在
		if userpost.Userpostid == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 410, "msg": "User post doesn't exist"})
			ctx.Abort()
			return
		}
		//用户存在，将user信息写入上下文
		ctx.Set("user_post", userpost)
		ctx.Next()
	}

}

//check auth for feedback
func AuthUserFeedback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//vcalidate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Wrong Token"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 405, "msg": "Wrong Token"})
			ctx.Abort()
			return
		}

		//验证通过后获取Claiim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户不存在
		if user.Userid == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 410, "msg": "User doesn't exist"})
			ctx.Abort()
			return
		}

		//用户存在，将user信息写入上下文
		ctx.Set("user_fb", user)
		ctx.Next()
	}
}
