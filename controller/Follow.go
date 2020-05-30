package controller

import (
	"huana/common"
	"huana/model"
	"huana/response"
	"net/http"
	"time"

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

func AddFollower(ctx *gin.Context) {
	db := common.GetDB()
	var follow = model.Follow{}

	ctx.Bind((&follow))

	userid := follow.Userid
	followerid := follow.Followerid

	newFollower := model.Follow{
		Userid:     userid,
		Followerid: followerid,
		Followdate: time.Now(),
	}

	db.Save(&newFollower)

	ctx.JSON(http.StatusOK, gin.H{"data": newFollower})
}

func DeleteFollower(ctx *gin.Context) {
	db := common.GetDB()

	userid := ctx.Param("userid")
	followerid := ctx.Param("followerid")

	var follow model.Follow

	db.Where("userid =? AND followerid = ?", userid, followerid).Find(&follow).Delete(&follow)
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
