package initialize

import (
	middleware "github.com/cn-joyconn/goadmin/middleware"
	gologs "github.com/cn-joyconn/gologs"

	// filetool "github.com/cn-joyconn/goutils/filetool"
	gin "github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitServer() *gin.Engine {
	// selfDir := filetool.SelfDir()
	var Router = gin.Default()

	// 日志
	Router.Use(middleware.GinLogger())
	gologs.GetLogger("").Info("use middleware Logger")

	//https
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	// 跨域
	Router.Use(middleware.Cors())
	gologs.GetLogger("").Info("use middleware cors")

	//swagger
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	gologs.GetLogger("").Info("register swagger handler")

	//Error
	Router.NoMethod(middleware.HandleNotFound)
	Router.NoRoute(middleware.HandleNotFound)
	Router.Use(middleware.ErrHandler())
	gologs.GetLogger("").Info("use middleware ErrorHandle")

	//view template
	Router.LoadHTMLGlob("templates/**/*")

	return Router

}
