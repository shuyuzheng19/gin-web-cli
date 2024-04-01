package upload

import (
	"gin-web/common"
	"gin-web/models"
	"gin-web/request"
	"gin-web/response"

	"github.com/gin-gonic/gin"
)

type FileService interface {
	GlobalUploadFile(ctx *gin.Context, isImage bool, uid *int, isPub bool) []response.SimpleFileResponse
	GetPublicFile(req request.FileRequest) common.PageInfo
	GetUserFile(uid int, req request.FileRequest) common.PageInfo
}

type FileRepository interface {
	FindByMd5(md5 string) string
	BatchSave(files []models.FileInfo) error
	FindFileInfos(uid int, req request.FileRequest) (_ []response.FileResponse, count int64)
}
