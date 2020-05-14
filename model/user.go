package model

import (
	//"github.com/jinzhu/gorm"
	"time"
)

//User table
type User struct {
	Userid       uint `gorm:"primary_key"`
	Username     string
	Phone        string
	Password     string
	Createdate   time.Time
	Profilepic   string
	Imei         string
	Qq           string
	Sex          int
	Email        string
	Address      string
	Birthday     time.Time
	Introduction string
}

//category table
type Category struct {
	Categoryid    uint `gorm:"primary_key"`
	Categoryname  string
	Categoryorder int
}

//user post table
type Userpost struct {
	Userpostid uint `gorm:"primary_key"`
	Userid     int
	Createdate time.Time
}

//moment post tabel
type Momentpost struct {
	Momentid   uint `gorm:"primary_key"`
	Userid     int
	Userpostid int
	Categoryid int
	Caption    string
	Image      string
	Likecount  int
	Createdate time.Time
}

//video post table
type Videopost struct{
	Videoid uint `gorm:"primary_key"`
	Videourl string
	Caption string
	Categoryid int
	Userpostid int
	Userid int
	Viewcount int
	Likecount int
}
