package user

import (
	"gin-web/handler"
	"gin-web/helper"
	"gin-web/middleware"
	"gin-web/request"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service UserService
}

func (u UserController) Login(ctx *gin.Context) {
	var userRequest request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		helper.ResultBadRequestFail(ctx, handler.GetValidateErr(userRequest, err))
		return
	}

	var token = u.service.Login(userRequest)

	helper.ResultSuccessToResponse(ctx, token)
}

func (u UserController) RetrievePassword(ctx *gin.Context) {
	var request request.RetrieveRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.ResultBadRequestFail(ctx, handler.GetValidateErr(request, err))
		return
	}

	u.service.RetrievePassword(request)

	helper.ResultSuccessToResponse(ctx, "验证码已发送至你的邮箱")
}

func (u UserController) SendRegisteredEmail(ctx *gin.Context) {

	var request request.SendEmailCode

	if err := ctx.ShouldBindQuery(&request); err != nil {
		helper.ResultBadRequestFail(ctx, handler.GetValidateErr(request, err))
		return
	}

	u.service.SendEmailCode(request.Email)

	helper.ResultSuccessToResponse(ctx, "验证码已发送至你的邮箱")
}

func (u UserController) ContactMe(ctx *gin.Context) {

	var request request.ContactRequest

	if err := ctx.ShouldBindQuery(&request); err != nil {
		helper.ResultBadRequestFail(ctx, handler.GetValidateErr(request, err))
		return
	}

	u.service.ContactMe(request)
}

func (u UserController) ResetPassword(ctx *gin.Context) {
	var request request.PasswordRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.ResultBadRequestFail(ctx, handler.GetValidateErr(request, err))
		return
	}

	u.service.UpdatePassword(request)

	helper.ResultSuccessToResponse(ctx, "修改成功")
}

func (u UserController) GetUser(ctx *gin.Context) {
	var user = handler.GetUser(ctx)
	helper.ResultSuccessToResponse(ctx, user.ToUserResponse())
}

func (u UserController) UserPage(ctx *gin.Context) {
	helper.ResultSuccessToResponse(ctx, "this is user page")
}

func (u UserController) AdminPage(ctx *gin.Context) {
	helper.ResultSuccessToResponse(ctx, "this is admin page")
}

func (u UserController) SuperAdminPage(ctx *gin.Context) {
	helper.ResultSuccessToResponse(ctx, "this is super_admin page")
}

func (u UserController) Register(ctx *gin.Context) {
	var userRequest request.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		helper.ResultBadRequestFail(ctx, handler.GetValidateErr(userRequest, err))
		return
	}

	u.service.Register(userRequest)

	helper.ResultSuccessToResponse(ctx, nil)
}

func (u UserController) GetUsers(ctx *gin.Context) {
	var filter request.UserFilter

	ctx.ShouldBindQuery(&filter)

	if filter.Page <= 0 {
		filter.Page = 1
	}

	var pageInfo = u.service.GetUsers(filter)

	helper.ResultSuccessToResponse(ctx, pageInfo)
}

func NewUserController() UserController {
	var service = NewUserService()
	middleware.GetJwtUser = service.GetUser
	middleware.GetToken = service.GetToken
	return UserController{service: service}
}
