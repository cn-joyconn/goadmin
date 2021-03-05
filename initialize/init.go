package initialize

import (
	"strconv"

	"github.com/cn-joyconn/goadmin/models/global"
	"github.com/gin-gonic/gin"
)
type  ExtInit func( *gin.Engine) bool
// Init 初始化
func Init(f ExtInit){
	InitConfig()
	InitDB(true)
	Router:=InitServer()
	RegistorRouters(Router)
	if f(Router) {
		Router.Run(":"+strconv.Itoa(global.AppConf.WebPort))
	}
	
}