package initialize

import (
	global "github.com/cn-joyconn/goadmin/models/global"
	filetool "github.com/cn-joyconn/goutils/filetool"
)

// InitConfig 加载配置
// @title    InitConfig
// @description   加载配置
// @auth      eric.zsp         时间（2021/03/04  16:04 ）
func InitConfig() {
	selfDir := filetool.SelfDir()
	appConfigPath := selfDir + "/conf/app.yml"
	dbConfigPath := selfDir + "/conf/db.yml"
	adminConfigPath := selfDir + "/conf/admin.yml"
	global.InitAppConf(appConfigPath)
	global.InitDBConf(dbConfigPath)
	global.LoadAdmin(adminConfigPath)	
}


