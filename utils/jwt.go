package utils

import (
	"gin-web/common"
	"gin-web/configs"
	"gin-web/helper"
	"gin-web/response"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateAccessToken 创建Token
func CreateAccessToken(id int, username string) response.TokenResponse {

	var config = configs.CONFIG.Jwt

	var create = time.Now()

	var expire = time.Now().Add(time.Duration(config.Expire) * (time.Hour * 24))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expire.Unix(),
		IssuedAt:  create.Unix(),
		Id:        strconv.Itoa(id),
		Issuer:    username,
		Subject:   username,
	})

	accessToken, err := claims.SignedString([]byte(config.Encrypt))

	if err != nil {
		helper.ErrorToResponse(common.CreateTokenFail)
	}

	return response.TokenResponse{
		Token:  accessToken,
		Expire: FormatDate(expire),
		Create: FormatDate(create),
	}

}

// ParseTokenToUserId 解析Token
func ParseTokenToUserId(token string) int {

	claims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.CONFIG.Jwt.Encrypt), nil
	})

	if err != nil || !claims.Valid {
		return -1
	}

	var uid, _ = strconv.Atoi(claims.Claims.(*jwt.StandardClaims).Id)

	return uid
}
