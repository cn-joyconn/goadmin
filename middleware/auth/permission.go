package auth

import (
	"path"
	"reflect"
	"runtime"

	"github.com/cn-joyconn/goadmin/utils"
	"github.com/cn-joyconn/goutils/array"
	"github.com/gin-gonic/gin"
)

type JoyPermissionGroup struct {
	GinGroup *gin.RouterGroup
}

type PermissionInfo struct {
	Permission string `json:"permission"`
	Url        string `json:"url"`
}

var PermissionGroup = &JoyPermissionGroup{}
var permissionMap = make(map[string]*[]PermissionInfo, 0)

// var funcMap = make(map[gin.HandlerFunc][]string)
func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		handleName := c.HandlerName()
		permissionInfos, exist := permissionMap[handleName]
		
		
		// isAjax :=false
		if exist {
			permissions :=  make([]string,0)
			for _,permissionInfo:=range *permissionInfos{
				permissions = append(permissions, permissionInfo.Permission)
			}
			permissions = array.RemoveDuplicateStr(permissions)
			validationResultModel := getAuthPass(c, permissions, true)
			pass := handelPass(utils.IsAjax(c), validationResultModel, c)
			if pass {
				c.Next()
			} else {
				c.Abort()
			}
		} else {
			c.Abort()
		}

	}
}
func (group *JoyPermissionGroup) registerHandlerPermission(relativePath string, permission string, handlers ...gin.HandlerFunc) {
	name := ""
	var permissions []PermissionInfo
	for _, h := range handlers {
		name = runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		permissionValue, exsit := permissionMap[name]
		if !exsit {
			permissions = make([]PermissionInfo, 0)
		}else{
			permissions = *permissionValue
		}
		permissions = append(permissions, PermissionInfo{Permission: permission, Url: joinPaths(group.GinGroup.BasePath(), relativePath)})
		permissionMap[name] = &permissions
	}
}

func (group *JoyPermissionGroup) GetAllPermissionName() []PermissionInfo {
	result := make([]PermissionInfo, 0)
	for _, perpermissions := range permissionMap {
		for _, p := range *perpermissions {
			result = append(result, p)
		}
	}
	// result = array.RemoveDuplicateStr(result)
	return result
}
func (group *JoyPermissionGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *JoyPermissionGroup {
	// ginGroup := group.ginGroup.Group(relativePath,handlers...)
	return &JoyPermissionGroup{
		GinGroup: group.GinGroup.Group(relativePath, handlers...),
	}

}

// POST is a shortcut for router.Handle("POST", path, handle).
func (group *JoyPermissionGroup) POST(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	// fmt.Printf(ginGroup.BasePath())
	group.registerHandlerPermission(relativePath, permission, handlers...)
	return group.GinGroup.POST(relativePath, handlers...)
}

// GET is a shortcut for router.Handle("GET", path, handle).
func (group *JoyPermissionGroup) GET(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(relativePath, permission, handlers...)
	return group.GinGroup.GET(relativePath, handlers...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle).
func (group *JoyPermissionGroup) DELETE(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(relativePath, permission, handlers...)
	return group.GinGroup.DELETE(relativePath, handlers...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle).
func (group *JoyPermissionGroup) PATCH(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(relativePath, permission, handlers...)
	return group.GinGroup.PATCH(relativePath, handlers...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle).
func (group *JoyPermissionGroup) PUT(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(relativePath, permission, handlers...)
	return group.GinGroup.PUT(relativePath, handlers...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle).
func (group *JoyPermissionGroup) OPTIONS(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(relativePath, permission, handlers...)
	return group.GinGroup.OPTIONS(relativePath, handlers...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle).
func (group *JoyPermissionGroup) HEAD(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(relativePath, permission, handlers...)
	return group.GinGroup.HEAD(relativePath, handlers...)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (group *JoyPermissionGroup) Any(relativePath string, permission string, handlers ...gin.HandlerFunc) gin.IRoutes {
	group.registerHandlerPermission(relativePath, permission, handlers...)
	return group.GinGroup.Any(relativePath, handlers...)
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
