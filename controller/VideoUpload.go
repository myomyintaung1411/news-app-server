package controller

import (
	"huana/common"
	"huana/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostVideo(ctx *gin.Context) {
	userpost, _ := ctx.Get("user_post")
	DB := common.GetDB()
	//model videopost

	var requestVideo = model.Videopost{}

	ctx.Bind(&requestVideo)

	//èŽ·å–å‚æ•°
	userpostid := userpost.(model.Userpost).Userpostid
	userid := userpost.(model.Userpost).Userid
	videourl := requestVideo.Videourl
	caption := requestVideo.Caption
	categoryid := requestVideo.Categoryid
	description := requestVideo.Description
	viewcount := requestVideo.Viewcount
	likecount := requestVideo.Likecount

	newVideo := model.Videopost{
		Videourl:    videourl,
		Caption:     caption,
		Categoryid:  categoryid,
		Userpostid:  userpostid,
		Userid:      userid,
		Description: description,
		Viewcount:   viewcount,
		Likecount:   likecount,
		Createdate:  time.Now(),
	}

	DB.Save(&newVideo)
	ctx.JSON(http.StatusOK, newVideo)

}

func GetAllVideo(ctx *gin.Context) {
	// Categoryid := ctx.Query("Categoryid")
	// println("leeeeeeeeeeeeeeeeeeeeeeee", Categoryid)
	DB := common.GetDB()
	id := ctx.PostForm("Categoryid")
	var requestVideo []model.Videopost
	DB.Where("categoryid =?", id).Find(&requestVideo)

	ctx.JSON(http.StatusOK, gin.H{"data": requestVideo})

}
