package auth

import (
	"reflect"
	"runtime"

	"github.com/cn-joyconn/goutils/array"
	"github.com/gin-gonic/gin"
)

type JoyAuthorizeGroup struct {
	GinGroup *gin.RouterGroup
}

var AuthorizeGroup = &JoyAuthorizeGroup{}
var AuthorizenMap = make(map[string][]string, 0)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		handleName := c.HandlerName()
		permissions, exist := permissionMap[handleName]
		// isAjax :=false
		if exist {
			validationResultModel := getAuthPass(c, permissions, true)
			//pass := handelPass(utils.IsAjax(c),validationResultModel,c)
			if validationResultModel.Pass == 2 {
				c.Next()
			} else {
				c.Abort()
			}
		} else {
			c.Abort()
		}
	}
}

func (group *JoyAuthorizeGroup) registerHandlerPermission(permission string, handlers ...gin.HandlerFunc) {
	name := ""
	exsit := false
	var permissions []string
	for _, h := range handlers {
		name = runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		permissions, exsit = permissionMap[name]
		if !exsit {
			permissions = make([]string, 0)
		}
		permissions = append(permissions, name)
		permissionMap[name] = permissions
	}
}

func (group *JoyAuthorizeGroup) GetAllPermissionName() []string {
	result := make([]string, 0)
	for _, perpermissions := range permissionMap {
		for _, p := range perpermissions {
			result = append(result, p)
		}
	}
	result = array.RemoveDuplicateStr(result)
	return result
}


func (group *JoyAuthorizeGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *JoyAuthorizeGroup {
	// ginGroup := group.ginGroup.Group(relativePath,handlers...)
	return &JoyAuthorizeGroup{
		GinGroup: group.GinGroup.Group(relativePath,handlers...),
	}

}

// POST is a shortcut for router.Handle("POST", path, handle).
func (group *JoyAuthorizeGroup) POST(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	// fmt.Printf(ginGroup.BasePath())
	group.registerHandlerPermission(permission, handlers...)
	return group.GinGroup.POST(relativePath, handlers...)
}

// GET is a shortcut for router.Handle("GET", path, handle).
func (group *JoyAuthorizeGroup) GET( relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(permission, handlers...)
	return group.GinGroup.GET(relativePath, handlers...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle).
func (group *JoyAuthorizeGroup) DELETE( relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(permission, handlers...)
	return group.GinGroup.DELETE(relativePath, handlers...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle).
func (group *JoyAuthorizeGroup) PATCH( relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(permission, handlers...)
	return group.GinGroup.PATCH(relativePath, handlers...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle).
func (group *JoyAuthorizeGroup) PUT( relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(permission, handlers...)
	return group.GinGroup.PUT(relativePath, handlers...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle).
func (group *JoyAuthorizeGroup) OPTIONS( relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(permission, handlers...)
	return group.GinGroup.OPTIONS(relativePath, handlers...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle).
func (group *JoyAuthorizeGroup) HEAD( relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(permission, handlers...)
	return group.GinGroup.HEAD(relativePath, handlers...)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (group *JoyAuthorizeGroup) Any( relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(permission, handlers...)
	return group.GinGroup.Any(relativePath, handlers...)
}
