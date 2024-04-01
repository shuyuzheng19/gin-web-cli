package v1

import (
	"gin-web/common"
	"gin-web/handler/user"
	"gin-web/middleware"
)

func (r *RouterGroup) SetupUserRouter(name string) *RouterGroup {
	var group = r.Api.Group(name)
	var controller = user.NewUserController()
	{
		group.POST("/login", controller.Login)
		group.POST("/register", controller.Register)
		group.GET("/send_email", controller.SendRegisteredEmail)
		group.POST("/send_pw_code", controller.RetrievePassword)
		group.POST("/reset_password", controller.ResetPassword)
		group.POST("contact_me", controller.ContactMe)
		group.GET("/", controller.GetUsers)
		group.Use(middleware.JwtMiddle(common.USER_ID))
		{
			group.GET("/get", controller.GetUser)
			group.GET("/info", controller.UserPage)
		}

		group.Use(middleware.JwtMiddle(common.ADMIN_ID))
		{
			group.GET("/admin", controller.AdminPage)
		}

		group.Use(middleware.JwtMiddle(common.SUPER_ADMIN_ID))
		{
			group.GET("/super", controller.SuperAdminPage)
		}
	}
	return r
}
