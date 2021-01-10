package global

import (
	"errors"
	"net/http"
	"strings"

	// handle "github.com/cn-joyconn/goadmin/handle"
	gin "github.com/gin-gonic/gin"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

//响应参数结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Url  string      `json:"url"`
	Wait int         `json:"wait"`
}

//返回结果辅助函数
func Result(code int, msg string, data interface{}, url string, wait int, header map[string]string, ctx *gin.Context ) {
	if strings.ToLower(ctx.Request.Method) == "post" {
		result := Response{
			Code: code,
			Msg:  msg,
			Data: data,
			Url:  url,
			Wait: wait,
		}

		if len(header) > 0 {
			for k, v := range header {
				ctx.Header(k,v)
			}
		}

		ctx.JSON(200, result)

		//Controller中this.StopRun()用法
		panic(errors.New("user stop run"))
	}

	if url == "" {
		url = ctx.Request.Referer()
		if url == "" {
			url = "/admin/index/index"
		}
	}

	ctx.Redirect(http.StatusFound, url)
}

//Success 成功、普通返回
func Success(ctx *gin.Context) {
	Result(SUCCESS, "操作成功", "", URL_BACK, 0, map[string]string{}, ctx)
}

//SuccessWithMessage 成功、返回自定义信息
func SuccessWithMessage(msg string, ctx *gin.Context) {
	Result(SUCCESS, msg, "", URL_BACK, 0, map[string]string{}, ctx)
}

//SuccessWithMessageAndURL 成功、返回自定义信息和url
func SuccessWithMessageAndURL(msg string, url string, ctx *gin.Context) {
	Result(SUCCESS, msg, "", url, 0, map[string]string{}, ctx)
}

//SuccessWithDetailed 成功、返回所有自定义信息
func SuccessWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *gin.Context) {
	Result(SUCCESS, msg, data, url, wait, header, ctx)
}

//Error 失败、普通返回
func Error(ctx *gin.Context) {
	Result(ERROR, "操作失败", "", URL_CURRENT, 0, map[string]string{}, ctx)
}

//ErrorWithMessage 失败、返回自定义信息
func ErrorWithMessage(msg string, ctx *gin.Context) {
	Result(ERROR, msg, "", URL_CURRENT, 0, map[string]string{}, ctx)
}

//ErrorWithMessageAndURL 失败、返回自定义信息和url
func ErrorWithMessageAndURL(msg string, url string, ctx *gin.Context) {
	Result(ERROR, msg, "", url, 0, map[string]string{}, ctx)
}

//ErrorWithDetailed 失败、返回所有自定义信息
func ErrorWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *gin.Context) {
	Result(ERROR, msg, data, url, wait, header, ctx)
}
