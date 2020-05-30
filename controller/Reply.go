package controller

import (
	"fmt"
	"huana/common"
	"huana/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RecordReply(ctx *gin.Context) {
	db := common.GetDB()

	var reply = model.Reply{}
	ctx.Bind(&reply)
	//获取参数
	commentid := reply.Commentid
	userid := reply.Userid
	username := reply.Username
	text := reply.Text
	like := reply.Likecount

	newReply := model.Reply{
		Commentid:  commentid,
		Userid:     userid,
		Username:   username,
		Text:       text,
		Likecount:  like,
		Createdate: time.Now(),
	}

	db.Save(&newReply)

	ctx.JSON(http.StatusOK, gin.H{"data": newReply})
}

func GetReplyCount(c *gin.Context) {
	db := common.GetDB()

	id := c.PostForm("Commentid")
	var count int
	var replies []model.Reply

	db.Model(&replies).Where("commentid = ?", id).Count(&count)

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetAllReply(c *gin.Context) {
	db := common.GetDB()

	postid := c.PostForm("Commentid")

	var replies []model.Reply
	db.Where("commentid = ? ", postid).Find(&replies)

	c.JSON(http.StatusOK, gin.H{"data": replies})
}

func DeleteReply(ctx *gin.Context) {
	db := common.GetDB()

	id := ctx.Param("replyid")

	var reply model.Reply
	db.Where("replyid = ?", id).Find(&reply)
	db.Delete(&reply)

	//ctx.JSON(http.StatusOK, gin.H{"data": newLike})

}

func DeleteReplyLikeRecord(c *gin.Context) {
	db := common.GetDB()

	id := c.Param("postid")
	field := c.Param("field")

	var like model.Likerecord
	db.Where("postid =? AND field = ?", id, field).Find(&like).Delete(&like)
	//db.Delete(&like)
}

func UpdateReplyLike(ctx *gin.Context) {
	db := common.GetDB()

	id := ctx.PostForm("Replyid")
	like := ctx.PostForm("Likecount")
	i, err := strconv.Atoi(like)
	if err != nil {
		fmt.Print(err.Error)
	}

	var reply model.Reply
	db.Where("replyid = ?", id).Find(&reply)

	reply.Likecount = i

	db.Save(&reply)
}
