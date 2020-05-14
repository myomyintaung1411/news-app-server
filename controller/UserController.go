package controller

import (
	"fmt"
	"huana/common"
	"huana/dto"
	"huana/model"
	"huana/response"
	"huana/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)
	//获取参数
	name := requestUser.Username
	telephone := requestUser.Phone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称为空给一个随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Username: name,
		Phone:    telephone,
		Password: string(hasePassowrd),
	}
	DB.Save(&newUser)
	//发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	telephone := c.PostForm("Phone")
	password := c.PostForm("Password")
	fmt.Print(len(password))
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	db.Where("phone=?", telephone).Find(&user)
	fmt.Println(user.Userid)
	if user.Userid == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "wrong user data")
		return
	}
	//判断密码是否正确
	fmt.Println("Password >>>>> ", user.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "Wrong password")
		//c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "Wrong token")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		//log.Printf("token generate error:%v", err)
	}

	//返回结果
	response.Success(c, gin.H{"token": token}, "Success")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("phone = ? ", telephone).First(&user)
	if user.Userid != 0 {
		return true
	}
	return false
}
