package dto

import (
	"huana/model"
	//"time"
)

type UserDto struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserpostDto struct {
	Userpostid int    `json:"id"`
	Userid     int    `json:"user_id"`
	Createdate string `json:"create_date"`
}

type UserInfoDto struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	AvatorImage  string `json:"image"`
	Introduction string `json:"introduction"`
	Gender       int    `json:"gender"`
	Birthday     string `json:"birthday"`
	Phone        string `json:"phone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		ID:       user.Userid,
		Name:     user.Username,
		Password: user.Password,
		Phone:    user.Phone,
	}
}

func ToUserPostDto(userpost model.Userpost) UserpostDto {
	return UserpostDto{
		Userpostid: userpost.Userpostid,
		Userid:     userpost.Userid,
		Createdate: userpost.Createdate.String(),
	}
}

// 以下所有是自己新加上去的哟～

func ToUserInfoDto(user model.User) UserInfoDto {
	return UserInfoDto{
		Id:           user.Userid,
		Name:         user.Username,
		AvatorImage:  user.Profilepic,
		Introduction: user.Introduction,
		Gender:       user.Sex,
		Birthday:     user.Birthday,
		Phone:        user.Phone,
	}
}
