package controller

import (
	"fmt"
	"huana/common"
	"huana/dto"
	"huana/model"
	"huana/response"
	"io"
	"os"
	"strconv"
	"time"

	//"huana/util"
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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Wrong phone number , it need to fill 11 length")
		return
	}
	if len(password) < 6 || len(password) > 8 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "Password field need min 6 - max 8 length")
		return
	}
	//如果名称为空给一个随机字符串
	if len(name) == 0 {
		//name = util.RandomString(10)
		response.Response(ctx, http.StatusUnprocessableEntity, 405, nil, "Username required")
		return
	}
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 410, nil, "User already exist with this phone number")
		return
	}
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Username:   name,
		Phone:      telephone,
		Password:   string(hasePassowrd),
		Createdate: time.Now(),
	}
	fmt.Println(newUser)
	DB.Save(&newUser)

	//发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "Success")
}

func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	telephone := c.PostForm("Phone")
	password := c.PostForm("Password")
	fmt.Print(len(password))
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 420, nil, "Wrong phone number , it need to fill 11 length")
		return
	}
	if len(password) < 6 || len(password) > 8 {
		response.Response(c, http.StatusUnprocessableEntity, 410, nil, "Password field need min 6 - max 8 length")
		return
	}
	//判断手机号是否存在
	var user model.User
	db.Where("phone=?", telephone).Find(&user)
	fmt.Println(user.Userid)
	if user.Userid == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "User with this phone number doesn't exist")
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

// 以下所有是自己新加上去的哟～

func UserInfo(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"info": dto.ToUserInfoDto(user.(model.User))}})

}

func UpdateUserName(c *gin.Context) {
	DB := common.GetDB()

	var user model.User

	u, _ := c.Get("user")        // 验证之后得来的值
	uid := u.(model.User).Userid // 通过验证得来的user得到user的id
	uname := c.PostForm("Username")

	DB.Where("Userid = ?", uid).Find(&user)

	user.Username = uname
	DB.Save(&user)

	response.Success(c, gin.H{"params": dto.ToUserInfoDto(user)}, "Success")
}

func UpdateUserProfileImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:3000/public/" + filename

	DB := common.GetDB()

	var user model.User

	u, _ := c.Get("user")        // 验证之后得来的值
	uid := u.(model.User).Userid // 通过验证得来的user得到user的id
	//image := c.PostForm("Profileimage")

	DB.Where("Userid = ?", uid).Find(&user)

	user.Profilepic = filepath
	DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"params": dto.ToUserInfoDto(user)})

}

func UpdateUserIntroduction(c *gin.Context) {
	DB := common.GetDB()

	var user model.User

	u, _ := c.Get("user")        // 验证之后得来的值
	uid := u.(model.User).Userid // 通过验证得来的user得到user的id
	intro := c.PostForm("Introduction")

	DB.Where("Userid = ?", uid).Find(&user)

	user.Introduction = intro
	DB.Save(&user)

	response.Success(c, gin.H{"params": dto.ToUserInfoDto(user)}, "Success")

}

func UpdateUserGender(c *gin.Context) {
	DB := common.GetDB()

	var user model.User

	u, _ := c.Get("user")        // 验证之后得来的值
	uid := u.(model.User).Userid // 通过验证得来的user得到user的id
	gender := c.PostForm("Gender")
	sex, _ := strconv.Atoi(gender)

	fmt.Println("sex is : ", sex)

	DB.Where("Userid = ?", uid).Find(&user)

	user.Sex = sex
	DB.Save(&user)

	response.Success(c, gin.H{"params": dto.ToUserInfoDto(user)}, "Success")

}

func UpdateUserBirthday(c *gin.Context) {
	DB := common.GetDB()

	var user model.User

	u, _ := c.Get("user")        // 验证之后得来的值
	uid := u.(model.User).Userid // 通过验证得来的user得到user的id
	birth := c.PostForm("Birthday")

	DB.Where("Userid = ?", uid).Find(&user)

	user.Birthday = birth
	DB.Save(&user)

	response.Success(c, gin.H{"params": dto.ToUserInfoDto(user)}, "Success")

}

func Feedback(c *gin.Context) {

	DB := common.GetDB()
	var fb = model.Feedback{}

	c.Bind(&fb)

	feedbackpost, _ := c.Get("user_fb")
	userid := feedbackpost.(model.User).Userid

	fbcontent := c.PostForm("fbcontent")
	fbscreencapture := c.PostForm("fbscreencapture")
	fbcategory := c.PostForm("fbcategory")
	nsk := "Good Job! Girl (>v<) "

	// if len(fbcontent) == 5 {
	// 	response.Response(c, http.StatusUnprocessableEntity, 422, nil, "反馈内容需要至少五个字！")
	// 	return
	// }

	feedback := model.Feedback{
		Userid:          userid,
		Fbcategory:      fbcategory,
		Fbcontent:       fbcontent,
		Fbscreencapture: fbscreencapture,
	}
	fmt.Println(feedback)
	DB.Save(&feedback)
	//返回结果
	response.Success(c, gin.H{"token": nsk}, "感谢您的反馈！")
}
