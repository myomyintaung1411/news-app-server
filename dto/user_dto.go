package dto

import "huana/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Username,
		Telephone: user.Phone,
	}
}
