package common

import (
	"github.com/dgrijalva/jwt-go"
	// "time"
	"huana/model"
)

var jwtkey = []byte("a_secret_create")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	// expireTime := time.Now().Add(1 * 60 * time.Second)
	claims := &Claims{
		UserId: user.Userid,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: expireTime.Unix(), //过期时间
			// IssuedAt:  time.Now().Unix(),
			Issuer:    "huanaguoji.newsapp",
			Subject:   "user token",
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
