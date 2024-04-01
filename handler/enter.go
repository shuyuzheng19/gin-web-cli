package handler

import (
	"gin-web/common"
	"gin-web/models"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetValidateErr(obj any, rawErr error) common.E {
	validationErrs, ok := rawErr.(validator.ValidationErrors)
	if !ok {
		return common.AutoFail(common.BadRequest)
	}
	field, ok := reflect.TypeOf(obj).FieldByName(validationErrs[0].Field())
	if ok {
		if e := field.Tag.Get("msg"); e != "" {
			return common.BadRequestFail(e)
		}
	}
	return common.AutoFail(common.BadRequest)
}

func GetPage(ctx *gin.Context) int {
	var pageStr = ctx.Query("page")

	if pageStr == "" {
		return 1
	}

	if page, err := strconv.Atoi(pageStr); err != nil {
		return 1
	} else {
		return page
	}
}

func GetUser(ctx *gin.Context) *models.User {
	var user, exists = ctx.Get("user")

	if !exists {
		return nil
	}

	return user.(*models.User)
}

func GetBool(ctx *gin.Context, queryName string) bool {
	var isPublic = ctx.DefaultQuery(queryName, "false")

	if flag, err := strconv.ParseBool(isPublic); err != nil {
		return false
	} else {
		return flag
	}
}
