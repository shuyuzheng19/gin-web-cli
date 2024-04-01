package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"mime/multipart"
	"regexp"
	"strings"
)

func ObjectToJson[T any](t T) string {
	var buff, _ = json.Marshal(&t)
	return string(buff)
}

func JsonToObject[T any](str string) *T {
	var result T

	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil
	}

	return &result
}

func IsImageURL(url string) bool {
	imageRegex := `^https?://.*\.(png|jpe?g|gif|svg|ico)$`

	regex := regexp.MustCompile(imageRegex)

	return regex.MatchString(url)
}

func IsImageFile(suffix string) bool {
	// 将文件名转换为小写字母，并获取文件扩展名
	suffix = strings.ToLower(suffix)

	// 检查文件扩展名是否为图片格式
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	for _, allowedExt := range allowedExtensions {
		if suffix == allowedExt {
			return true
		}
	}

	return false
}

func GetFileMd5(file multipart.File) string {
	hash := md5.New()

	io.Copy(hash, file)

	md5 := hex.EncodeToString(hash.Sum(nil))

	file.Close()

	return md5
}
