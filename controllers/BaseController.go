package controllers

import (
	"net/http"
	"strconv"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	global "github.com/cn-joyconn/goadmin/models/global"
	adminServices "github.com/cn-joyconn/goadmin/services/admin"
	"github.com/cn-joyconn/goutils/strtool"
	"github.com/gin-gonic/gin"
	// gologs "github.com/cn-joyconn/gologs"
	// "strconv"
)

var adminUserService *adminServices.AdminUserService = new(adminServices.AdminUserService)

type BaseController struct {
	//gin.Context
}

// 设置模板
// 第一个参数模板，第二个参数为data
func (bc *BaseController) ResponseHtmlByStatusCode(c *gin.Context, name string, statusCode int, data gin.H) {
	// gologs.GetLogger("").Info(c.FullPath())
	// gologs.GetLogger("").Info(c.Request.RequestURI)
	if data == nil {
		data = make(map[string]interface{})
	}
	data["staticResourcesPathJs"] = global.AppConf.JSPath
	data["staticResourcesPathCss"] = global.AppConf.CSSPath
	data["staticResourcesPathFile"] = global.AppConf.FilePath
	data["staticResourcesPathImage"] = global.AppConf.ImagePath
	data["pageContentPath"] = global.AppConf.ContextPath
	data["pageTitleSuffix"] = global.AppConf.Name
	

	c.HTML(statusCode, name, data)
}

// 第一个参数模板，第二个参数为data
func (bc *BaseController) ResponseHtml(c *gin.Context, name string, data gin.H) {

	bc.ResponseHtmlByStatusCode(c, name, http.StatusOK, data)
}

//JSON输出
func (bc *BaseController) ApiJson(c *gin.Context, code int, msg string, data interface{}, allcount int64) {
	if data == nil {
		data = ""
	}
	// c.ServeJSON()
	// c.StopRun()
	c.JSON(http.StatusOK, &global.Response{
		Code:     code,
		Msg:      msg,
		Data:     data,
		Url:      "",
		Wait:     0,
		AllCount: allcount,
	})
}

//返回成功的API成功
func (bc *BaseController) ApiSuccess(c *gin.Context, msg string, data interface{}) {
	bc.ApiJson(c, global.SUCCESS, msg, data, 0)
}

//返回成功的API成功
func (bc *BaseController) ApiDataList(c *gin.Context, msg string, data interface{}, allcount int64) {
	bc.ApiJson(c, global.SUCCESS, msg, data, allcount)
}

//返回失败的API请求
func (bc *BaseController) ApiError(c *gin.Context, msg string, data interface{}) {
	bc.ApiJson(c, global.ERROR, msg, data, 0)
}

//返回失败且带code的API请求
func (bc *BaseController) ApiErrorCode(c *gin.Context, msg string, data interface{}, code int) {
	bc.ApiJson(c, code, msg, data, 0)
}

// //请求出错
// func (bc *BaseController) RequestError(c *gin.Context,msg string, route string) {
// 	c.Redirect(route, 302)
// }

// //请求成功
// func (bc *BaseController) RequestSuccess(c *gin.Context,msg string, route string) {
// 	c.Redirect(route, 302)
// }

func (bc *BaseController) GetContextUserId(c *gin.Context) int {
	userid := bc.GetContextUserIdStr(c)
	if strtool.IsBlank(userid) {
		return 0
	}
	uid, err := strconv.Atoi(userid)
	if err != nil {
		return 0
	}
	return uid
}
func (bc *BaseController) GetContextUserIdStr(c *gin.Context) string {
	userid := c.GetString(global.Context_UserId)
	if strtool.IsBlank(userid) {
		userid = global.TokenHelper.GetMyAuthenticationID(c)
		if !strtool.IsBlank(userid) {
			c.Set(global.Context_UserId, userid)
		}
	}
	return userid

}
func (bc *BaseController) GetContextUserObj(c *gin.Context) *adminModel.AdminUserBasic {
	userObj, exist := c.Get(global.Context_UserInfo)
	if !exist {
		userid := bc.GetContextUserIdStr(c)
		if !strtool.IsBlank(userid) {
			var userids = make([]string, 1)
			userids[0] = userid
			models := (&adminServices.AdminUserService{}).GetUserInfoByUserIDS(userids)

			if models != nil && len(*models) > 0 {
				// userObj = models[0]
				c.Set(global.Context_UserInfo, (*models)[0])
				return &((*models)[0])
			}
		}
	} else {
		return userObj.(*adminModel.AdminUserBasic)
	}
	return nil
}
