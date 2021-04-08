package global

import (
	loginToken "github.com/cn-joyconn/goadmin/utils/loginToken"
	// snowflake "github.com/cn-joyconn/goutils/snowflake"
)

//不做任何操作
const URL_CURRENT = "url://current"

//刷新页面
const URL_RELOAD = "url://reload"

//返回上一个页面
const URL_BACK = "url://back"



//上下文中存储用户信息的键
const Context_UserInfo = "Context_UserInfo"
//上下文中存储用户ID的键
const Context_UserId = "Context_UserId"

//AdminCatalog admin缓存类别
var AdminCatalog string

//AdminCacheName admin缓存名称
var AdminCacheName string

// var SnowflakeWorker *snowflake.Worker

var TokenHelper *loginToken.TokenHelper 
