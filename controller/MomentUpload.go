package controller

import (
	"fmt"
	"huana/common"
	"huana/dto"
	"huana/model"
	"huana/response"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//insert new row as user post
func NewPost(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	id := user.(model.User).Userid
	fmt.Print(id)

	DB := common.GetDB()
	// var requestUser = model.Userpost{}
	// //json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	// ctx.Bind(&requestUser)
	//获取参数
	//id := requestUser.Userid

	newUserpost := model.Userpost{
		Userid:     id,
		Createdate: time.Now(),
	}
	DB.Save(&newUserpost)
	//发送token
	token, err := common.ReleaseUserPostToken(newUserpost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "Wrong token")
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "Success , user post table get new field!")
}

func UserpostInfo(ctx *gin.Context) {
	userpost, _ := ctx.Get("user_post")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user_post": dto.ToUserPostDto(userpost.(model.Userpost))}})
}

func NewMomentPost(ctx *gin.Context) {

	userpost, _ := ctx.Get("user_post")
	userid := userpost.(model.Userpost).Userid
	userpostid := userpost.(model.Userpost).Userpostid

	DB := common.GetDB()
	var post = model.Momentpost{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&post)
	//获取参数
	//userid := post.Userid
	//userpostid := post.Userpostid
	caption := post.Caption
	link := post.Image
	like := post.Likecount
	//date := post.Createdate

	newMomentpost := model.Momentpost{
		Userid:     userid,
		Userpostid: userpostid,
		Caption:    caption,
		Image:      link,
		Likecount:  like,
		Createdate: time.Now(),
	}
	DB.Save(&newMomentpost)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"moment_post": &newMomentpost}})
	//json.NewEncoder(w).Encode(&newMomentpost)
	//发送token
	// token, err := common.ReleaseUserPostToken(newMo)
	// if err != nil {
	// 	response.Response(ctx, http.StatusInternalServerError, 500, nil, "Wrong token")
	// 	log.Printf("token generate error:%v", err)
	// 	return
	// }

	// //返回结果
	// response.Success(ctx, gin.H{"token": token}, "Success")
}

func AllMomentPost(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	userid := ctx.PostForm("Userid")

	var moments []model.Momentpost
	db.Where("userid =?", userid).Find(&moments)

	ctx.JSON(http.StatusOK, gin.H{"data": moments})
}

func UpdateMomentLike(ctx *gin.Context) {
	db := common.GetDB()

	id := ctx.PostForm("Momentpostid")
	like := ctx.PostForm("Likecount")
	i, err := strconv.Atoi(like)
	if err != nil {
		fmt.Print(err.Error)
	}

	var moment model.Momentpost
	db.Where("momentpostid = ?", id).Find(&moment)

	moment.Likecount = i

	db.Save(&moment)
}

func CheckLike(ctx *gin.Context) {
	db := common.GetDB()

	userid, err := strconv.Atoi(ctx.PostForm("Userid"))
	postid, err := strconv.Atoi(ctx.PostForm("Postid"))
	s := ctx.PostForm("Field")
	if err != nil {
		fmt.Print(err.Error)
	}
	var like model.Likerecord

	db.Where("userid = ? AND postid = ? AND field =?", userid, postid, s).Find(&like)

	ctx.JSON(http.StatusOK, gin.H{"data": like})
}

func RecordLike(ctx *gin.Context) {
	db := common.GetDB()

	userid, err := strconv.Atoi(ctx.PostForm("Userid"))
	postid, err := strconv.Atoi(ctx.PostForm("Postid"))
	s := ctx.PostForm("Field")

	if err != nil {
		fmt.Print(err.Error)
	}

	newLike := model.Likerecord{
		Userid: userid,
		Postid: postid,
		Field:  s,
	}

	db.Save(&newLike)

	ctx.JSON(http.StatusOK, gin.H{"data": newLike})
}

func DeleteLike(ctx *gin.Context) {
	db := common.GetDB()

	userid := ctx.Param("userid")
	postid := ctx.Param("postid")
	field := ctx.Param("field")

	fmt.Print(userid)

	var like model.Likerecord
	db.Where("userid =? AND postid =? AND field =?", userid, postid, field).Find(&like)
	db.Delete(&like)

}
