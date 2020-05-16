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
