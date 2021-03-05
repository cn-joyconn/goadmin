package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "os"
	// "fmt"
	"time"

	log "github.com/cn-joyconn/gologs"
)

//GinLogger  gin定制logger
func GinLogger() gin.HandlerFunc {  
	logger :=log.GetLogger("gin")
	
 
	return func (c *gin.Context) {  
	   // 开始时间  
	   start := time.Now()  
	   // 处理请求  
	   c.Next()  
	   // 结束时间  
	   end := time.Now()  
	   //执行时间  
	   latency := end.Sub(start)  
 
	   path := c.Request.URL.Path 
 
	   clientIP := c.ClientIP()  
	   method := c.Request.Method  
	   statusCode := c.Writer.Status()  
	   logger.Info(fmt.Sprintf( "| %3d | %13v | %15s | %s  %s |",  
		  statusCode,  
		  latency,  
		  clientIP,  
		  method, path,  
	   ))
	}  
 }