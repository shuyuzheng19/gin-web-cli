package upload

import (
	"gin-web/common"
	"gin-web/configs"
	"gin-web/helper"
	"gin-web/models"
	"gin-web/request"
	"gin-web/response"
	"gin-web/utils"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileServiceImpl struct {
	config     configs.UploadConfig
	repository FileRepository
}

var mb int64 = 1024 * 1024

func getFiles(ctx *gin.Context) []*multipart.FileHeader {
	var files, err = ctx.MultipartForm()

	if err != nil {
		configs.LOGGER.Warn("请求体找不到文件")
		helper.ErrorToResponseAndError(common.NoFile, err)
		return nil
	}

	return files.File["files"]
}

func (u FileServiceImpl) GlobalUploadFile(ctx *gin.Context, isImage bool, uid *int, isPub bool) []response.SimpleFileResponse {
	var frs = make([]response.SimpleFileResponse, 0)

	var files = getFiles(ctx)

	var errResponse response.SimpleFileResponse

	var infos = make([]models.FileInfo, 0)

	for _, file := range files {

		var size = file.Size

		var fileName = file.Filename

		var suffix = filepath.Ext(fileName)

		var create = time.Now()

		if isImage {
			if utils.IsImageFile(suffix) {
				errResponse = response.SimpleFileResponse{
					Status:  "fail",
					Message: "这不是一个图片文件",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				}
				frs = append(frs, errResponse)
				configs.LOGGER.Info("上传文件错误 这不是一个图片文件", zap.String("fileName", fileName))
				continue
			} else if size > u.config.MaxImageSize*mb {
				errResponse = response.SimpleFileResponse{
					Status:  "fail",
					Message: "图片文件大小超出",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				}
				frs = append(frs, errResponse)
				configs.LOGGER.Info("上传文件错误 图片文件大小超出", zap.String("fileName", fileName))
				continue
			}
		} else {
			if size > u.config.MaxFileSize*mb {
				errResponse = response.SimpleFileResponse{
					Status:  "fail",
					Message: "文件大小超出",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				}
				frs = append(frs, errResponse)
				configs.LOGGER.Info("上传文件错误 文件大小超出", zap.String("fileName", fileName))
				continue
			}
		}

		var f, err = file.Open()

		if err != nil {
			continue
		}

		var md5 = utils.GetFileMd5(f)

		var newName = md5 + suffix

		var saveFilePath = u.config.Path + "/" + newName

		var url string

		var existsMd5 = false

		url = u.config.Uri + "/" + newName

		if uid != nil {
			if dbUrl := u.repository.FindByMd5(md5); dbUrl != "" {
				existsMd5 = true
				url = dbUrl
			} else {
				var uploadError = ctx.SaveUploadedFile(file, saveFilePath)

				if uploadError != nil {
					continue
				}
			}
		} else {
			var uploadError = ctx.SaveUploadedFile(file, saveFilePath)

			if uploadError != nil {
				continue
			}
		}

		var fileInfo = models.FileInfo{
			OldName: fileName,
			NewName: newName,
			Suffix:  suffix,
			Size:    size,
			UserID:  uid,
			FileMd5: md5,
			IsPub:   isPub,
		}

		if !existsMd5 {
			fileInfo.FileMd5Info = models.FileMd5Info{
				Md5:          md5,
				Url:          url,
				AbsolutePath: saveFilePath,
			}
		}

		infos = append(infos, fileInfo)

		configs.LOGGER.Info("文件上传成功", zap.String("fileName", fileInfo.NewName),
			zap.String("url", fileInfo.FileMd5Info.Url),
			zap.String("md5", fileInfo.FileMd5Info.Md5),
			zap.Int64("size", fileInfo.Size))

		frs = append(frs, response.SimpleFileResponse{
			Status:  "ok",
			Message: "上传成功",
			Name:    fileName,
			Create:  utils.FormatDate(create),
			Url:     url,
		})
	}

	if uid != nil && len(infos) > 0 {
		u.repository.BatchSave(infos)
	}

	return frs
}

func (fs FileServiceImpl) GetPublicFile(req request.FileRequest) common.PageInfo {
	var files, count = fs.repository.FindFileInfos(-1, req)

	return common.PageInfo{
		Page:  req.Page,
		Total: count,
		Size:  common.FilePageCount,
		Data:  files,
	}
}

func (fs FileServiceImpl) GetUserFile(uid int, req request.FileRequest) common.PageInfo {
	var files, count = fs.repository.FindFileInfos(uid, req)

	return common.PageInfo{
		Page:  req.Page,
		Total: count,
		Size:  common.FilePageCount,
		Data:  files,
	}
}

func NewFileService() FileService {
	return FileServiceImpl{config: configs.CONFIG.Upload, repository: NewFileRepository(configs.DB)}
}
