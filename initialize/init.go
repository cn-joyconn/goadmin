package initialize

import (
	"strconv"

	"github.com/cn-joyconn/goadmin/models/global"
	"github.com/cn-joyconn/goadmin/utils/joyCaptcha"
	"github.com/cn-joyconn/goutils/snowflake"
	"github.com/gin-gonic/gin"
)

type ExtInit func(*gin.Engine) bool

// Init 初始化
func Init(f ExtInit) {
	InitConfig()
	initCacheCfg()
	initSnowflake()
	initTokenHelper()
	InitDB(true)
	joyCaptcha.InitCaptcha()
	
	Router := InitServer()
	RegistorRouters(Router)
	if f(Router) {
		Router.Run(":" + strconv.Itoa(global.AppConf.WebPort))
	}

}

func initSnowflake(){
	snowflakeWorker, _ := snowflake.NewWorker(global.AppConf.SnowflakeWorkID)
	global.SnowflakeWorker = snowflakeWorker
}

func initCacheCfg() {
	var ok bool
	global.AdminCatalog, ok = global.AppConf.Cache["adminCatalog"]
	if !ok {
		global.AdminCatalog = "joyconn"
	}
	global.AdminCacheName, ok = global.AppConf.Cache["adminCacheName"]
	if !ok {
		global.AdminCacheName = "admin"
	}
}
