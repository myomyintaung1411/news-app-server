package common

import (
	"huana/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("a_secret_create")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

type UserpostClaims struct {
	Userpostid int
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.Userid,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: expireTime.Unix(), //过期时间
			//IssuedAt:  time.Now().Unix(),
			Issuer:  "xietong.me",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}

//token for userpost table
func ReleaseUserPostToken(userpost model.Userpost) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &UserpostClaims{
		Userpostid: userpost.Userpostid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "xietong.me",
			Subject:   "userpost token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//parse token for userpost
func ParseUserPostToken(tokenString string) (*jwt.Token, *UserpostClaims, error) {
	UserpostClaims := &UserpostClaims{}
	token, err := jwt.ParseWithClaims(tokenString, UserpostClaims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, UserpostClaims, err
}
