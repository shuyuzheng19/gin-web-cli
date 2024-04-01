package upload

import (
	"gin-web/handler"
	"gin-web/helper"
	"gin-web/request"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	service FileService
}

func (f FileController) UploadFile(ctx *gin.Context) {
	var uid = handler.GetUser(ctx).ID

	var flag = handler.GetBool(ctx, "isPublic")

	var result = f.service.GlobalUploadFile(ctx, false, &uid, flag)

	helper.ResultSuccessToResponse(ctx, result)
}

func (f FileController) UploadAvatar(ctx *gin.Context) {
	var result = f.service.GlobalUploadFile(ctx, true, nil, false)

	helper.ResultSuccessToResponse(ctx, result)
}

func (f FileController) UploadImage(ctx *gin.Context) {

	var flag = handler.GetBool(ctx, "isPublic")

	var uid = handler.GetUser(ctx).ID

	var result = f.service.GlobalUploadFile(ctx, true, &uid, flag)

	helper.ResultSuccessToResponse(ctx, result)
}

func (f FileController) GetPublicFiles(ctx *gin.Context) {

	var req = request.FileRequest{Page: 1, Sort: "date"}

	ctx.ShouldBindQuery(&req)

	var result = f.service.GetPublicFile(req)

	helper.ResultSuccessToResponse(ctx, result)
}

func (f FileController) GetCurrentUserFiles(ctx *gin.Context) {
	var req = request.FileRequest{Page: 1, Sort: "date"}

	ctx.ShouldBindQuery(&req)

	var uid = handler.GetUser(ctx).ID

	var result = f.service.GetUserFile(uid, req)

	helper.ResultSuccessToResponse(ctx, result)
}

func NewFileController() FileController {
	return FileController{service: NewFileService()}
}
