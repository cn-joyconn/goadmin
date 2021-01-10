package main

import (
	"testing"

	// handle "github.com/cn-joyconn/goadmin/handle"
	modles "github.com/cn-joyconn/goadmin/models"
	gin "github.com/gin-gonic/gin"
)

func TestGoAdmin(t *testing.T) {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// r.Use(handle.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	initFilter()
	//启动beego
	// web.Run()
	router := gin.Default()
	router.Run()
}

func TestInitDB(t *testing.T) {
	modles.InitDB()
}
func initFilter() {

	// //过滤器：加日志
	// web.InsertFilter("/admin/*",web.BeforeRouter, sysinit.FilterAddLog)

	// //后台权限过滤
	// web.InsertFilter("/admin/*",web.BeforeRouter, sysinit.FilterAdminPermission)

	// //自定义错误页面
	// web.ErrorController(&controllers.ErrorController{})

}
