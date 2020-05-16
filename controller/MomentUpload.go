package controller

import (
	"huana/common"
	"huana/dto"
	"huana/model"
	"huana/response"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//insert new row as user post
func NewPost(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	id := user.(model.User).Userid


	DB := common.GetDB()
	var requestUser = model.Userpost{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)
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
	response.Success(ctx, gin.H{"token": token}, "Success")
}

func UserpostInfo(ctx *gin.Context) {
	userpost, _ := ctx.Get("user_post")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user_post": dto.ToUserPostDto(userpost.(model.Userpost))}})
}

func NewMomentPost(ctx *gin.Context) {
	DB := common.GetDB()
	var post = model.Momentpost{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&post)
	//获取参数
	userid := post.Userid
	userpostid := post.Userpostid
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
