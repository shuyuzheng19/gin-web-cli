package utils

import (
	"crypto/md5"
	"encoding/hex"
	"gin-web/common"
	"gin-web/helper"
	"math/rand"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	var hasher = md5.New()
	hasher.Write([]byte(password))
	var hashedPassword = hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}

// 将密码进行hash化
func BcryptPassword(password string) string {
	var bcryptPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		helper.ErrorToResponseAndError(common.ERROR, err)
	}
	return string(bcryptPassword)
}

// 验证密码
func ValidatorPassword(password, hashPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}

// 随机生成6位验证码
func RandomNumberCode() string {
	return strconv.Itoa(rand.Intn(900000) + 100000)
}
