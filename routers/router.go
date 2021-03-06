package initialize

import (
	// "net/http"

	controllers "github.com/cn-joyconn/goadmin/controllers"
	// filetool "github.com/cn-joyconn/goutils/filetool"
	// gologs "github.com/cn-joyconn/gologs"
	// middleware "github.com/cn-joyconn/goadmin/middleware"
	"github.com/gin-gonic/gin"
)
//通用
func InitCommonRouter(publicGroup *gin.RouterGroup, authGroup *gin.RouterGroup, permissioneGroup *gin.RouterGroup) {
	controller := &controllers.CommonController{}
	commonRouter := publicGroup.Group("common")
	{
		commonRouter.GET("authimage", controller.AuthImage)
	}

}
//登录
func InitAccountRouter(publicGroup *gin.RouterGroup, authGroup *gin.RouterGroup, permissioneGroup *gin.RouterGroup) {
	controller := &controllers.AccountController{}
	accountRouter := publicGroup.Group("account")
	{
		accountRouter.GET("login", controller.LoginPage)
		accountRouter.POST("dologin", controller.LoginApi)
	}
	myInfoRouter := authGroup.Group("myinfo")
	{
		myInfoRouter.GET("getme", controller.LoginPage)
	}
	userRouter := publicGroup.Group("user")
	{
		userRouter.GET("list", controller.LoginPage)
		userRouter.POST("update", controller.LoginApi) 
	}

}
