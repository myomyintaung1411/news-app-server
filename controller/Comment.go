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

func RecordComment(ctx *gin.Context) {
	db := common.GetDB()

	var comment = model.Comment{}
	ctx.Bind(&comment)
	//获取参数
	userid := comment.Userid
	postid := comment.Postid
	field := comment.Field
	name := comment.Username
	text := comment.Text
	like := comment.Likecount

	newComment := model.Comment{
		Userid:     userid,
		Postid:     postid,
		Field:      field,
		Username:   name,
		Text:       text,
		Likecount:  like,
		Createdate: time.Now(),
	}

	db.Save(&newComment)

	ctx.JSON(http.StatusOK, gin.H{"data": newComment})
}

func DeleteComment(ctx *gin.Context) {
	db := common.GetDB()

	id := ctx.Param("commentid")

	var comment model.Comment
	db.Where("commentid = ?", id).Find(&comment)
	db.Delete(&comment)

	//ctx.JSON(http.StatusOK, gin.H{"data": newLike})

}

func DeleteCommentLikeRecord(c *gin.Context) {
	db := common.GetDB()

	id := c.Param("postid")
	field := c.Param("field")

	var like model.Likerecord
	db.Where("postid =? AND field = ?", id, field).Find(&like).Delete(&like)
	//db.Delete(&like)
}

func GetCommentCount(c *gin.Context) {
	db := common.GetDB()

	id := c.PostForm("Postid")
	field := c.PostForm("Field")
	var count int
	var comments []model.Comment

	db.Model(&comments).Where("postid = ? AND field =?", id, field).Count(&count)

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetAllComment(c *gin.Context) {
	db := common.GetDB()

	postid := c.PostForm("Postid")
	field := c.PostForm("Field")

	var comments []model.Comment
	db.Where("postid = ? AND field =?", postid, field).Find(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func UpdateCommentLike(ctx *gin.Context) {
	db := common.GetDB()

	id := ctx.PostForm("Commentid")
	like := ctx.PostForm("Likecount")
	i, err := strconv.Atoi(like)
	if err != nil {
		fmt.Print(err.Error)
	}

	var comment model.Comment
	db.Where("commentid = ?", id).Find(&comment)

	comment.Likecount = i

	db.Save(&comment)
}
