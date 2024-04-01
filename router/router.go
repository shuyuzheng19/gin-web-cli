package router

import (
	"fmt"
	"gin-web/configs"
	"gin-web/helper"
	"gin-web/middleware"
	v1 "gin-web/router/v1"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	config   configs.ServerConfig //服务器配置
	service  *gin.Engine          //gin的服务
	apiGroup *gin.RouterGroup     //gin的路由组
}

// 为路由组添加中间件
func (r *Router) AddMiddlewares(middlewares ...gin.HandlerFunc) *Router {
	r.apiGroup.Use(middlewares...)
	return r
}

// 启动web服务器
func (r *Router) RunServer() {
	var corsConfig = configs.CONFIG.Cors

	//是否开启跨域
	if corsConfig.Enable {
		r.service.Use(middleware.Cors(corsConfig))
	}

	fmt.Println(configs.CONFIG.Upload.Path)

	r.service.Static("/static/", configs.CONFIG.Upload.Path)

	r.service.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "pong"})
	})

	var server = &http.Server{
		ReadTimeout:  time.Duration(r.config.ReadTimeOut) * time.Second,  //读取超时时间
		WriteTimeout: time.Duration(r.config.WriteTimeOut) * time.Second, //写入超时时间
		Addr:         r.config.Addr,                                      //监听的HOST&&PORT
		Handler:      r.service,                                          //gin的服务
	}

	var err = server.ListenAndServe() //启动服务器

	helper.PanicError(err)

	configs.LOGGER.Info("服务启动成功")
}

func (r *Router) SetupRouters() *Router {
	var group = &v1.RouterGroup{Api: r.apiGroup}
	group.SetupUserRouter("users").SetupFileRouter("file")
	return r
}

func NewRouter(config configs.ServerConfig) *Router {
	var service *gin.Engine

	//判断是否为发布模式，如果为true，则是没有任何中间件和日志
	if config.Release {
		gin.SetMode(gin.ReleaseMode)
		service = gin.New()
	} else {
		service = gin.Default()
	}

	service.Use(middleware.ErrorMiddle)

	var apiGroup = service.Group(config.ApiPrefix)

	var author = configs.CONFIG.Author

	configs.LOGGER.Info("欢迎使用Yushu Gin脚手架",
		zap.String("author", author.Name),
		zap.String("home", author.Home),
		zap.String("github", author.Github),
		zap.String("version", author.Version),
		zap.Time("start_time", time.Now()),
	)

	return &Router{
		config:   config,
		service:  service,
		apiGroup: apiGroup,
	}
}
