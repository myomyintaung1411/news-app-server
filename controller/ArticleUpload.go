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

func NewArticlePost(ctx *gin.Context) {

	userpost, _ := ctx.Get("user_post")
	userid := userpost.(model.Userpost).Userid
	userpostid := userpost.(model.Userpost).Userpostid

	DB := common.GetDB()
	var post = model.Articlepost{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&post)
	//获取参数
	//userid := post.Userid
	//userpostid := post.Userpostid
	categoryid := post.Categoryid
	caption := post.Caption
	content := post.Content
	cover := post.Cover
	view := post.Viewcount
	like := post.Likecount
	//date := post.Createdate

	newArticlepost := model.Articlepost{
		Userpostid: userpostid,
		Userid:     userid,
		Categoryid: categoryid,
		Caption:    caption,
		Content:    content,
		Cover:      cover,
		Viewcount:  view,
		Likecount:  like,
		Createdate: time.Now(),
	}
	DB.Save(&newArticlepost)
	ctx.JSON(http.StatusOK, gin.H{"article_post": &newArticlepost})
}

func AllArticlePost(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	categoryid := ctx.PostForm("Categoryid")

	var articles []model.Articlepost
	db.Where("categoryid = ?", categoryid).Find(&articles)

	ctx.JSON(http.StatusOK, gin.H{"data": articles})
}

func UpdateArticleLike(ctx *gin.Context) {
	db := common.GetDB()

	id := ctx.PostForm("Articlepostid")
	like := ctx.PostForm("Likecount")
	i, err := strconv.Atoi(like)
	if err != nil {
		fmt.Print(err.Error)
	}

	var article model.Articlepost
	db.Where("articlepostid = ?", id).Find(&article)

	article.Likecount = i

	db.Save(&article)
}


