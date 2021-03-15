package global

// import (
// 	"strconv"

// 	adminModel "github.com/cn-joyconn/goadmin/models/admin"
// 	// adminServices "github.com/cn-joyconn/goadmin/services/admin"
// 	"github.com/cn-joyconn/goutils/strtool"
// 	"github.com/gin-gonic/gin"
// )
// var adminUserService *adminServices.AdminUserService
// type RegistorJoyContextFunc func(c *JoyContext)

// func RegistorJoyContext(h RegistorJoyContextFunc) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var templateV map[string]interface{} /*创建集合 */
// 		templateV = make(map[string]interface{})
// 		ctx := &JoyContext{
// 			c,
// 			templateV,
// 		}
// 		h(ctx)
// 	}
// }

// type JoyContext struct {
// 	*gin.Context
// 	templateVMap map[string]interface{}
// }
