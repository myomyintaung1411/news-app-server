package controller

import (
	"huana/common"
	"huana/model"
	"huana/response"

	"github.com/gin-gonic/gin"
)

//insert new row as user post
func GetCategory(ctx *gin.Context) {

	DB := common.GetDB()
	var categorys []model.Category
	//获取参数

	DB.Find(&categorys)

	//返回结果
	response.Success(ctx, gin.H{"category": categorys}, "Success , get categories")
}
