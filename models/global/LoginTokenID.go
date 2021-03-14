package global

import (
	"strconv"

	"github.com/cn-joyconn/goutils/strtool"
	"github.com/gin-gonic/gin"
)





func GetContextUserId(c *gin.Context) int{
	userid := GetContextUserIdStr(c)
	if strtool.IsBlank(userid){
		return 0
	}
	uid,err:=strconv.Atoi(userid)
	if err!=nil{
		return 0
	}
	return uid
}
func GetContextUserIdStr(c *gin.Context) string{
	userid := c.GetString(Context_UserId)
	if !strtool.IsBlank(userid){
		userid = TokenHelper.GetMyAuthenticationID(c)
		if !strtool.IsBlank(userid){
			c.Set(Context_UserId,userid)
		}
	}
	return userid
	
}
// func GetContextUserObj(c *gin.Context) adminModel.{
// 	userid := GetContextUserIdStr(c)
// 	if strtool.IsBlank(userid){
// 		return 0
// 	}
// 	uid,err:=strconv.Atoi(userid)
// 	if err!=nil{
// 		return 0
// 	}
// 	return uid
// }