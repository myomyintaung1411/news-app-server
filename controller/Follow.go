package controller

import (
	"huana/common"
	"huana/model"
	"huana/response"

	"github.com/gin-gonic/gin"
)

//insert new row as user post
func GetFollower(ctx *gin.Context) {

	DB := common.GetDB()
	var follow = model.Follow{}
	// //json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&follow)
	//获取参数
	id := follow.Userid

	var follows []model.Follow
	//DB.Where("userid = ?", id).Find(&follow).Count(&count)
	//DB.Table("follows").Where("userid = ?", id).Count(&count)
	//DB.Table("follows").Where("userid =", id).Find(&follows)
	//DB.Find(&users, "userid = ?", id)
	DB.Find(&follows, "userid = ?", id)

	//返回结果
	response.Success(ctx, gin.H{"follower": follows}, "Success , get follower")
}

func GetFollowing(ctx *gin.Context) {

	DB := common.GetDB()
	var follow = model.Follow{}
	// //json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&follow)
	//获取参数
	id := follow.Followerid

	var follows []model.Follow
	//DB.Where("userid = ?", id).Find(&follow).Count(&count)
	//DB.Table("follows").Where("userid = ?", id).Count(&count)
	//DB.Table("follows").Where("userid =", id).Find(&follows)
	//DB.Find(&users, "userid = ?", id)
	DB.Find(&follows, "followerid = ?", id)

	//返回结果
	response.Success(ctx, gin.H{"following": follows}, "Success , get following")
}
