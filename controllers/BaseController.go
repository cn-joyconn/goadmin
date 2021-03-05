package controllers

import (
	"net/http"

	global "github.com/cn-joyconn/goadmin/models/global"
	"github.com/gin-gonic/gin"
	// "strconv"
)

type BaseController struct {
	gin.Context
}

// 设置模板
// 第一个参数模板，第二个参数为data
func (bc *BaseController) ResponseHtml(c *gin.Context, name string, data gin.H) {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["staticResourcesPathJs"] = global.AppConf.JSPath
	data["staticResourcesPathCss"] = global.AppConf.CSSPath
	data["staticResourcesPathFile"] = global.AppConf.FilePath
	data["staticResourcesPathImage"] = global.AppConf.ImagePath
	userObj, exit := c.Get(global.Context_UserInfo)
	if exit {
		data["userInfo"] = userObj
	}
	c.HTML(http.StatusOK, name, data)
}

//JSON输出
func (bc *BaseController) ApiJson(c *gin.Context,code int16, msg string, data interface{}) {
	if data == nil {
		data = ""
	}
	// c.ServeJSON()
	// c.StopRun()
	c.JSON(http.StatusOK, &global.Response{
		Code: int(code),
		Msg:  msg,
		Data: data,
		Url:  "",
		Wait: 0,
	})
}

//返回成功的API成功
func (bc *BaseController) ApiSuccess(c *gin.Context,msg string, data interface{}) {
	bc.ApiJson(c,0, msg, data)
}

//返回失败的API请求
func (bc *BaseController) ApiError(c *gin.Context,msg string, data interface{}) {
	bc.ApiJson(c,-1, msg, data)
}

//返回失败且带code的API请求
func (bc *BaseController) ApiErrorCode(c *gin.Context,msg string, data interface{}, code int16) {
	bc.ApiJson(c,code, msg, data)
}

// //请求出错
// func (bc *BaseController) RequestError(c *gin.Context,msg string, route string) {
// 	c.Redirect(route, 302)
// }

// //请求成功
// func (bc *BaseController) RequestSuccess(c *gin.Context,msg string, route string) {
// 	c.Redirect(route, 302)
// }
