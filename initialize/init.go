package initialize

import (
	// "log"
	"strconv"

	"github.com/cn-joyconn/goadmin/models/global"
	"github.com/cn-joyconn/goadmin/utils/joyCaptcha"

	// "github.com/cn-joyconn/goutils/snowflake"
	"github.com/gin-gonic/gin"
	snowflake "github.com/cn-joyconn/goutils/snowflake"
)

type ExtInit func(*gin.Engine) bool

// Init 初始化
func Init(f ExtInit) {
	InitConfig()
	initCacheCfg()
	// initSnowflake()
	initIdGenerator()
	initTokenHelper()
	InitDB(true)
	joyCaptcha.InitCaptcha()

	Router := InitServer()
	RegistorRouters(Router)

	if f(Router) {
		Router.Run(":" + strconv.Itoa(global.AppConf.WebPort))
	}

}

// func initSnowflake(){
// 	snowflakeWorker, _ := snowflake.NewWorker(global.AppConf.SnowflakeWorkID)
// 	global.SnowflakeWorker = snowflakeWorker
// }
func initIdGenerator() {
	var options = snowflake.NewSnowOptions(global.AppConf.SnowflakeWorkID)
	options.WorkerIdBitLength = 12
	options.SeqBitLength = 10
	snowflake.InitGenerator(options)

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
