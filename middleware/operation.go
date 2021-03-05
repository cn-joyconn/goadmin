package middleware

// import (
// 	"bytes"

// 	"github.com/cn-joyconn/goadmin/models/global"
// 	"github.com/cn-joyconn/goadmin/models/admin"
// 	gologs "github.com/cn-joyconn/gologs"

// 	// "gin-vue-admin/global"
// 	// "gin-vue-admin/model"
// 	// "gin-vue-admin/model/request"
// 	// "gin-vue-admin/service"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"go.uber.org/zap"
// )

// func OperationRecord() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var body []byte
// 		var userId uint32
// 		if c.Request.Method != http.MethodGet {
// 			var err error
// 			body, err = ioutil.ReadAll(c.Request.Body)
// 			if err != nil {
// 				gologs.GetLogger("").Error("read body from request error:", zap.Any("err", err))
// 			} else {
// 				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
// 			}
// 		}
// 		if uid, ok := c.Get(global.Context_UserId); ok {
// 			userId = uid.(uint32)
// 		}
// 		record := admin.AdminLog{
// 			Ip:     c.ClientIP(),
// 			Title: c.Request.Method,
// 			Url:   c.Request.URL.Path,
// 			Agent:  c.Request.UserAgent(),
// 			Body:   string(body),
// 			AdminUserId: userId,
// 		}
// 		// 存在某些未知错误 TODO
// 		//values := c.Request.Header.Values("content-type")
// 		//if len(values) >0 && strings.Contains(values[0], "boundary") {
// 		//	record.Body = "file"
// 		//}
// 		writer := responseBodyWriter{
// 			ResponseWriter: c.Writer,
// 			body:           &bytes.Buffer{},
// 		}
// 		c.Writer = writer
// 		now := time.Now()

// 		c.Next()

// 		latency := time.Now().Sub(now)
// 		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
// 		record.Status = c.Writer.Status()
// 		record.Latency = latency
// 		record.Resp = writer.body.String()

// 		if err := service.CreateSysOperationRecord(record); err != nil {
// 			gologs.GetLogger("").Error("create operation record error:", zap.Any("err", err))
// 		}
// 	}
// }

// type responseBodyWriter struct {
// 	gin.ResponseWriter
// 	body *bytes.Buffer
// }

// func (r responseBodyWriter) Write(b []byte) (int, error) {
// 	r.body.Write(b)
// 	return r.ResponseWriter.Write(b)
// }
