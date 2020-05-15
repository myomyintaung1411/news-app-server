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
	Momentpostid uint `gorm:"primary_key"`
	Userid       int
	Userpostid   int
	Caption      string
	Image        string
	Likecount    int
	Createdate   time.Time
}

//article post table
type Articlepost struct {
	Articlepostid uint `gorm:"primary_key"`
	Userpostid    int
	Userid        int
	Categoryid    int
	Caption       string
	Content       string
	Cover         string
	Viewcount     int
	Likecount     int
	Createdate    time.Time
}

//video post table
type Videopost struct {
	Videopostid uint `gorm:"primary_key"`
	Videourl    string
	Caption     string
	Categoryid  int
	Userpostid  int
	Userid      int
	Viewcount   int
	Likecount   int
	Createdate  time.Time
}

//history table
type History struct {
	Historyid     uint `gorm:"primay_key"`
	Userid        int
	Createdate    time.Time
	Videopostid   int
	Articlepostid int
	Momentpostid  int
}

//favourite table
type Favourite struct {
	Favid         uint `gorm:"primay_key"`
	Userid        int
	Createdate    time.Time
	Videopostid   int
	Articlepostid int
	Momentpostid  int
}

//comment table
type Comment struct {
	Commentid  uint `gorm:"primary_key"`
	Userid     int
	Postid     int
	Text       int
	Createdate time.Time
}

//advetisement table
type Advertisement struct {
	Advid        uint `gorm:"primary_key"`
	Caption      string
	Advimage     string
	Createdate   time.Time
	Validday     int
	Activestatus int
}

//follow table
type Follow struct {
	Followid   uint `gorm:"primary_key"`
	Userid     int
	Followerid int
	Followdate time.Time
}
