package v1

import (
	"gin-web/common"
	"gin-web/handler/upload"
	"gin-web/middleware"
)

func (r *RouterGroup) SetupFileRouter(name string) *RouterGroup {
	var group = r.Api.Group(name)
	var controller = upload.NewFileController()
	{
		group.POST("upload_avatar", controller.UploadFile)
		group.GET("public_file", controller.GetPublicFiles)
		group.Use(middleware.JwtMiddle(common.ADMIN_ID))
		{
			group.GET("current_file", controller.GetCurrentUserFiles)
			group.POST("upload", controller.UploadFile)
			group.POST("upload_image", controller.UploadFile)
		}
	}
	return r
}
