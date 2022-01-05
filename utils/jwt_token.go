package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserToken struct {
	jwt.StandardClaims
	UserName string `json:"user_name"`
}

func GenToken(UserName string, expireDuration time.Duration, secret_key []byte) (string, error) {
	user := UserToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			Issuer:    "micro_gin_vue",
		},
		UserName: UserName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(secret_key)
}

// 前端用户token过期时间
var FrontUserExpireDuration = time.Hour * 1
var FrontUserSecretKey = []byte("front_user_token")

// 后端用户过期时间

var AdminUserExpireDuration = time.Hour * 2
var AdminUserSecretKey = []byte("admin_user_token")

func AuthToken(tokenString string, secretKey []byte) (*UserToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(t *jwt.Token) (key interface{}, err error) {
		return secretKey, err
	})
	if err != nil {
		return nil, err
	}
	claims, is_ok := token.Claims.(*UserToken)
	if is_ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token valid err")
}
