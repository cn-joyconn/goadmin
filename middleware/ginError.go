package middleware

import (
	"net/http"

	"github.com/cn-joyconn/goadmin/controllers"
	"github.com/cn-joyconn/goadmin/models/global"
	"github.com/cn-joyconn/goadmin/utils"
	"github.com/gin-gonic/gin"
)

// 错误处理的结构体
type JoyError struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	JoySuccess     = NewJoyError(http.StatusOK, global.SUCCESS, "success")
	JoyServerError = NewJoyError(http.StatusInternalServerError, global.ServiceError, "系统异常，请稍后重试!")
	JoyNotFound    = NewJoyError(http.StatusNotFound, global.NotFound, http.StatusText(http.StatusNotFound))
)

func JoyOtherError(message string) *JoyError {
	return NewJoyError(http.StatusForbidden, global.ERROR, message)
}

func (e *JoyError) Error() string {
	return e.Msg
}

func NewJoyError(statusCode, Code int, msg string) *JoyError {
	return &JoyError{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	err := JoyNotFound
	if utils.IsAjax(c) {
		c.JSON(err.StatusCode, err)
		c.Abort()
	} else {
		baseController := &controllers.BaseController{}
		baseController.ResponseHtmlByStatusCode(c, "layout/error_404.html", 404, gin.H{
			"pageTitle": "404错误",
		})
		// c.Request.Response.StatusCode = 404
		c.Abort()
	}

	return
}
func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *JoyError
				if e, ok := err.(*JoyError); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = JoyOtherError(e.Error())
				} else {
					Err = JoyServerError
				}
				// 记录一个错误的日志

				if utils.IsAjax(c) {
					c.JSON(Err.StatusCode, Err)
					c.Abort()
				} else if utils.IsPage(c) {
					baseController := &controllers.BaseController{}
					// c.Request.Response.StatusCode = 500
					baseController.ResponseHtmlByStatusCode(c, "layout/error_500.html", 500, gin.H{
						"pageTitle": "500错误",
					})
					c.Abort()
				} else {
					c.JSON(Err.StatusCode, "")
					c.Abort()
				}
				return
			}
		}()
		c.Next()
	}
}
