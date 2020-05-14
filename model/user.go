package model

import 
(
	// "github.com/jinzhu/gorm"
"time"
)

// type User struct {
// 	gorm.Model
// 	Name      string `gorm:"varchar(20);not null"`
// 	Telephone string `gorm:"varchar(11);not null;unique"`
// 	Password  string `gorm:"size:255;not null"`

// }


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