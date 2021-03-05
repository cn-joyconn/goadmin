package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 错误处理的结构体
type JoyError struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}


var (
	JoySuccess     = NewJoyError(http.StatusOK, 0, "success")
	JoyServerError = NewJoyError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	JoyNotFound    = NewJoyError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
)

func JoyOtherError(message string) *JoyError {
	return NewJoyError(http.StatusForbidden, 100403, message)
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
	c.JSON(err.StatusCode,err)
	return
}
func ErrHandler() gin.HandlerFunc  {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *JoyError
				if e,ok := err.(*JoyError); ok {
					Err = e
				}else if e, ok := err.(error); ok {
					Err = JoyOtherError(e.Error())
				}else{
					Err = JoyServerError
				}
				// 记录一个错误的日志
				c.JSON(Err.StatusCode,Err)
				return
			}
		}()
		c.Next()
	}
}